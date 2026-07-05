// Patched by patch-cli-auth-commands.sh — top-level logout shortcut.

package cli

import (
	"github.com/spf13/cobra"
)

func initLogoutCmd(parent *cobra.Command) error {
	cmd := &cobra.Command{
		Use:   "logout",
		Short: "Clear all stored authentication credentials",
		Long: `Clear all stored authentication credentials from both the OS keychain and config file.

This removes all credentials previously set via login or configure.`,
		RunE: runAuthLogoutCmd,
	}
	parent.AddCommand(cmd)
	return nil
}
