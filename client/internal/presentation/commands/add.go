package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"yadroTestAssignment/client/internal/contracts"
)

func NewAddCmd(svc contracts.DNSService) *cobra.Command {
	return &cobra.Command{
		Use:   "add <ip>",
		Short: "Add a DNS server",
		Long:  "Add a new DNS server by its IP address.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ip := args[0]
			if err := svc.Add(ip); err != nil {
				return fmt.Errorf("failed to add DNS server %s: %w", ip, err)
			}
			fmt.Printf("DNS server %s added successfully.\n", ip)
			return nil
		},
	}
}
