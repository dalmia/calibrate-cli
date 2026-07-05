// Patched by patch-cli-auth-commands.sh — top-level login shortcut.

package cli

import (
	"github.com/spf13/cobra"
)

func initLoginCmd(parent *cobra.Command) error {
	cmd := &cobra.Command{
		Use:   "login",
		Short: "Interactively configure authentication credentials",
		Long: `Interactively configure authentication credentials for calibrate.
Secret credentials are stored in the OS keychain when available,
with a config file fallback.

All fields are optional — press Enter to skip any field you don't need.
Use the configure command for both authentication and global parameters.`,
		RunE: runAuthLoginCmd,
	}
	parent.AddCommand(cmd)
	return nil
}
