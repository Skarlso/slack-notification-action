package pkg

import (
	"fmt"

	"github.com/slack-go/slack"
)

// SlackClient defines a client for slack.
//go:generate counterfeiter -o fakes/fake_slack_client.go . SlackClient
type SlackClient interface {
	PostMessage(channelID string, options ...slack.MsgOption) (string, string, error)
}

// Notifier notifies.
type Notifier struct {
	Message string
	Channel string
	Client  SlackClient
}

// NewNotifier creates a new Slack notifier.
func NewNotifier(message string, channel string, client SlackClient) *Notifier {
	return &Notifier{
		Message: message,
		Channel: channel,
		Client:  client,
	}
}

// Notify notifies.
func (n *Notifier) Notify() error {
	fmt.Printf("sending message to channel %s\n", n.Channel)
	channelID, ts, err := n.Client.PostMessage(n.Channel, slack.MsgOptionText(n.Message, false))
	if err != nil {
		return fmt.Errorf("failed to post api message: %w", err)
	}
	fmt.Printf("message posted to channel %s at %s\n", channelID, ts)
	return nil
}
