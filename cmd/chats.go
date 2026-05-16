package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var chatsCmd = &cobra.Command{
	Use:   "chats",
	Short: "List all chats",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := loadMessenger()
		if err != nil {
			return err
		}

		chats, err := client.ListChats()
		if err != nil {
			return fmt.Errorf("list chats: %w", err)
		}

		if jsonOutput {
			return printJSON(map[string]any{
				"count": len(chats),
				"chats": chats,
			})
		}

		fmt.Printf("Found %d chats:\n\n", len(chats))
		for i, chat := range chats {
			topic := chat.Topic
			if topic == "" {
				topic = "(no topic)"
			}
			fmt.Printf("%d. %s\n", i+1, topic)
			fmt.Printf("   ID: %s\n", chat.ID)
			fmt.Printf("   Type: %s\n", chat.ChatType)
			if chat.CreatedDateTime != "" {
				fmt.Printf("   Updated: %s\n", chat.CreatedDateTime)
			}
			fmt.Println()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(chatsCmd)
}
