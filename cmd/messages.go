package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var messagesLimit int

var messagesCmd = &cobra.Command{
	Use:   "messages <chat-id>",
	Short: "Get messages from a chat",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		chatID := args[0]

		client, err := loadMessenger()
		if err != nil {
			return err
		}

		messages, err := client.GetMessages(chatID, messagesLimit)
		if err != nil {
			return fmt.Errorf("get messages: %w", err)
		}

		if jsonOutput {
			return printJSON(map[string]any{
				"chat_id":  chatID,
				"count":    len(messages),
				"messages": messages,
			})
		}

		fmt.Printf("Messages in chat %s:\n\n", chatID)
		for i, msg := range messages {
			fmt.Printf("%d. [%s] %s\n", i+1, msg.CreatedDateTime, msg.From.User.DisplayName)
			fmt.Printf("   %s\n", msg.Body.Content)
			fmt.Println()
		}
		return nil
	},
}

func init() {
	messagesCmd.Flags().IntVar(&messagesLimit, "limit", 50, "Maximum number of messages to retrieve")
	rootCmd.AddCommand(messagesCmd)
}
