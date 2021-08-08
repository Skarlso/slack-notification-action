package cmd

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"

	"github.com/Skarlso/slack-notification-action/pkg"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Slack notifier command",
		Run:   runRootCmd,
	}
	rootArgs struct {
		token           string
		channel         string
		message         string
		timestamp       string
		threadTimestamp string
	}
)

func init() {
	flag := rootCmd.Flags()
	// Server Configs
	flag.StringVar(&rootArgs.token, "token", "", "--token slack-token")
	flag.StringVar(&rootArgs.channel, "channel", "", "--channel The ID of the channel to post messages to.")
	flag.StringVar(&rootArgs.message, "message", "", "--message The message to send.")
	flag.StringVar(&rootArgs.timestamp, "timestamp", "", "--timestamp The timestamp of the previously sent message.")
	flag.StringVar(&rootArgs.threadTimestamp, "thread-ts", "", "--thread-ts The timestamp of the previously sent message. If set, message will be sent in thread of that message.")
}

// runRootCmd runs the main notifier command.
func runRootCmd(cmd *cobra.Command, args []string) {
	client := slack.New(rootArgs.token)
	n := pkg.NewNotifier(rootArgs.message, rootArgs.channel, rootArgs.timestamp, rootArgs.threadTimestamp, client)
	if err := n.Notify(); err != nil {
		fmt.Printf("failed to send notification: %s\n", err)
		os.Exit(1)
	}
}

// Execute runs the main krok command.
func Execute() error {
	return rootCmd.Execute()
}
