package commands

import (
	"os"

	"github.com/spf13/cobra"
	"yadroTestAssignment/client/internal/contracts"
)

func NewRootCmd(svc contracts.DNSService) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "dns-cli",
		Short: "CLI client for managing DNS servers",
		Long: `dns-cli is a command-line tool for managing DNS servers
on a remote machine via the DNS management API.

The server address is configured via the DNS_SERVER_HOST environment variable
(default: http://localhost:8000).

Examples:
  dns-cli add 8.8.8.8
  dns-cli add 1.1.1.1
  dns-cli list
  dns-cli delete 8.8.8.8`,
		SilenceUsage: true,
	}

	cmd.AddCommand(
		NewAddCmd(svc),
		NewDeleteCmd(svc),
		NewListCmd(svc),
	)

	return cmd
}

func Execute(svc contracts.DNSService) {
	if err := NewRootCmd(svc).Execute(); err != nil {
		os.Exit(1)
	}
}
