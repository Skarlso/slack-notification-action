package cmd

import (
	"fmt"
	"os"

	"github.com/Skarlso/slack-notification-action/pkg"

	"github.com/slack-go/slack"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "root",
		Short: "Slack notifier command",
		Run:   runRootCmd,
	}
	rootArgs struct {
		token   string
		channel string
		message string
	}
)

func init() {
	flag := rootCmd.Flags()
	// Server Configs
	flag.StringVar(&rootArgs.token, "token", "", "--token slack-token")
	flag.StringVar(&rootArgs.channel, "channel", "", "--channel ")
	flag.StringVar(&rootArgs.message, "message", "", "--message Event occurred.")
}

// runRootCmd runs the main notifier command.
func runRootCmd(cmd *cobra.Command, args []string) {
	client := slack.New(rootArgs.token)
	n := pkg.NewNotifier(rootArgs.message, rootArgs.channel, client)
	if err := n.Notify(); err != nil {
		fmt.Printf("failed to send notification: %s", err)
		os.Exit(1)
	}
}

// Execute runs the main krok command.
func Execute() error {
	return rootCmd.Execute()
}
