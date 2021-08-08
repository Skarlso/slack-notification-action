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
	Message         string
	Channel         string
	Timestamp       string
	ThreadTimestamp string
	Client          SlackClient
}

// NewNotifier creates a new Slack notifier.
func NewNotifier(message string, channel string, timestamp string, threadTimestamp string, client SlackClient) *Notifier {
	return &Notifier{
		Message:         message,
		Channel:         channel,
		Timestamp:       timestamp,
		ThreadTimestamp: threadTimestamp,
		Client:          client,
	}
}

// Notify notifies.
func (n *Notifier) Notify() error {
	fmt.Printf("sending message to channel %s\n", n.Channel)
	opts := make([]slack.MsgOption, 0)
	opts = append(opts, slack.MsgOptionText(n.Message, false))
	if n.Timestamp != "" {
		opts = append(opts, slack.MsgOptionUpdate(n.Timestamp))
	}
	if n.ThreadTimestamp != "" {
		opts = append(opts, slack.MsgOptionTS(n.ThreadTimestamp))
	}
	channelID, ts, err := n.Client.PostMessage(n.Channel, opts...)
	if err != nil {
		return fmt.Errorf("failed to post api message: %w", err)
	}
	fmt.Printf("message posted to channel %s at %s\n", channelID, ts)
	fmt.Printf("::set-output name=channel::%s\n", channelID)
	fmt.Printf("::set-output name=timestamp::%s\n", ts)
	return nil
}
