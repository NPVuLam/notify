package ntfy

import "github.com/nikoksr/notify/v2"

var _ notify.SendConfig = (*SendConfig)(nil)

// SendConfig represents the configuration needed for sending a message.
//
// This struct complies with the notify.SendConfig interface and allows you to alter
// the behavior of the send function. This can be achieved by either passing send options
// to the send function or by manipulating the fields of this struct in your custom
// message renderer.
//
// All fields of this struct are exported to offer maximum flexibility to users.
// However, users must be aware that they are responsible for managing thread-safety
// and other similar concerns when manipulating these fields directly.
type SendConfig struct {
	Subject     string
	Message     string
	Attachments []notify.Attachment
	Metadata    map[string]any
	DryRun      bool

	// Ntfy specific fields

	ParseMode   Mode
	Priority    Priority
	Tags        []string
	Delay       string
	ClickAction string
}

// SetAttachments adds attachments to the message. This method is needed as part of the notify.SendConfig interface.
func (c *SendConfig) SetAttachments(attachments ...notify.Attachment) {
	c.Attachments = attachments
}

// SetMetadata sets the metadata of the message. This method is needed as part of the notify.SendConfig interface.
func (c *SendConfig) SetMetadata(metadata map[string]any) {
	c.Metadata = metadata
}

// SetDryRun sets the dry run flag of the message. This method is needed as part of the notify.SendConfig interface.
func (c *SendConfig) SetDryRun(dryRun bool) {
	c.DryRun = dryRun
}

// Send options

// SendWithPriority is a send option that sets the priority of the message.
func SendWithPriority(priority Priority) notify.SendOption {
	return func(config notify.SendConfig) {
		if typedConf, ok := config.(*SendConfig); ok {
			typedConf.Priority = priority
		}
	}
}

// SendWithParseMode is a send option that sets the parse mode of the message.
func SendWithParseMode(parseMode Mode) notify.SendOption {
	return func(config notify.SendConfig) {
		if typedConf, ok := config.(*SendConfig); ok {
			typedConf.ParseMode = parseMode
		}
	}
}

// SendWithTags is a send option that sets the tags of the message.
func SendWithTags(tags ...string) notify.SendOption {
	return func(config notify.SendConfig) {
		if typedConf, ok := config.(*SendConfig); ok {
			typedConf.Tags = tags
		}
	}
}

// SendWithDelay is a send option that sets the delay of the message.
func SendWithDelay(delay string) notify.SendOption {
	return func(config notify.SendConfig) {
		if typedConf, ok := config.(*SendConfig); ok {
			typedConf.Delay = delay
		}
	}
}

// SendWithClickAction is a send option that sets the click action of the message.
func SendWithClickAction(clickAction string) notify.SendOption {
	return func(config notify.SendConfig) {
		if typedConf, ok := config.(*SendConfig); ok {
			typedConf.ClickAction = clickAction
		}
	}
}
