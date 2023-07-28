package slack

import (
	"strings"
	"sync"

	"github.com/nikoksr/onelog"
	nopadapter "github.com/nikoksr/onelog/adapter/nop"
	"github.com/slack-go/slack"

	"github.com/nikoksr/notify/v2"
)

// Compile time check to make sure the service implements the notify.Service interface.
var _ notify.Service = (*Service)(nil)

// defaultMessageRenderer is a helper function to format messages for Slack.
func defaultMessageRenderer(conf *SendConfig) string {
	var builder strings.Builder

	builder.WriteString(conf.Subject)
	builder.WriteString("\n\n")
	builder.WriteString(conf.Message)

	return builder.String()
}

// Service is a structure that contains data needed for interaction with Slack's APIs.
type Service struct {
	client *slack.Client

	name          string
	mu            sync.RWMutex
	logger        onelog.Logger
	renderMessage func(conf *SendConfig) string
	dryRun        bool

	// Slack specific
	channelIDs    []string
	escapeMessage bool
}

func (s *Service) applyOptions(opts ...Option) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, opt := range opts {
		opt(s)
	}
}

// New creates a new instance of the Slack service with a default configuration. It receives as input the required Slack
// token and optional configurations. If no configuration is provided, the default values are used.
//
// Note: This function never returns an error. It has a return value for consistency with other services.
func New(token string, opts ...Option) (*Service, error) {
	client := slack.New(token)

	s := &Service{
		client:        client,
		name:          "slack",
		logger:        nopadapter.NewAdapter(),
		renderMessage: defaultMessageRenderer,
		dryRun:        false,
	}

	s.applyOptions(opts...)

	return s, nil
}

// Name returns the name of the service, which identifies the type of service in use (in this case, Slack).
func (s *Service) Name() string {
	return s.name
}

// AddRecipients appends given channel IDs onto an internal list that Send uses to distribute the notifications.
func (s *Service) AddRecipients(channelIDs ...string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.channelIDs = append(s.channelIDs, channelIDs...)
	s.logger.Info().Int("count", len(channelIDs)).Int("total", len(s.channelIDs)).Msg("Recipients added")
}
