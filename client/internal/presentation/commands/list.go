package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"yadroTestAssignment/client/internal/contracts"
)

func NewListCmd(svc contracts.DNSService) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all DNS servers",
		Long:  "Retrieve and display all currently configured DNS servers.",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			lines, err := svc.GetAll()
			if err != nil {
				return fmt.Errorf("failed to get DNS servers: %w", err)
			}

			if len(lines) == 0 {
				fmt.Println("No DNS servers configured.")
				return nil
			}

			fmt.Println("Configured DNS servers:")
			for i, line := range lines {
				fmt.Printf("  %d. %s\n", i+1, line)
			}
			return nil
		},
	}
}
