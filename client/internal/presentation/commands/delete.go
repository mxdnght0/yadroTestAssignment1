package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"yadroTestAssignment/client/internal/contracts"
)

func NewDeleteCmd(svc contracts.DNSService) *cobra.Command {
	return &cobra.Command{
		Use:   "delete <ip>",
		Short: "Delete a DNS server",
		Long:  "Remove a DNS server by its IP address.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ip := args[0]
			if err := svc.Delete(ip); err != nil {
				return fmt.Errorf("failed to delete DNS server %s: %w", ip, err)
			}
			fmt.Printf("DNS server %s deleted successfully.\n", ip)
			return nil
		},
	}
}
