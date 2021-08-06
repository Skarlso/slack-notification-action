package pkg_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Skarlso/slack-notification-action/pkg"
	"github.com/Skarlso/slack-notification-action/pkg/fakes"
)

func TestNotifySlack(t *testing.T) {
	fakeClient := &fakes.FakeSlackClient{}
	fakeClient.PostMessageReturns("test-channel", "123456", nil)
	n := pkg.Notifier{
		Message: "test message",
		Channel: "test-channel",
		Client:  fakeClient,
	}
	err := n.Notify()
	assert.NoError(t, err)
	assert.Equal(t, fakeClient.PostMessageCallCount(), 1)
}

func TestNotifySlackFailsWhenSlackErrors(t *testing.T) {
	fakeClient := &fakes.FakeSlackClient{}
	fakeClient.PostMessageReturns("test-channel", "123456", errors.New("nope"))
	n := pkg.Notifier{
		Message: "test message",
		Channel: "test-channel",
		Client:  fakeClient,
	}
	err := n.Notify()
	assert.EqualError(t, err, "failed to post api message: nope")
}
