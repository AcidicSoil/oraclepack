package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/app"
)

var validateCmd = &cobra.Command{
	Use:   "validate [pack.md]",
	Short: "Validate an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := app.Config{PackPath: args[0]}
		a := app.New(cfg)
		if err := a.LoadPack(); err != nil {
			return err
		}
		fmt.Println("Pack is valid.")
		return nil
	},
}

var listCmd = &cobra.Command{
	Use:   "list [pack.md]",
	Short: "List steps in an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := app.Config{PackPath: args[0]}
		a := app.New(cfg)
		if err := a.LoadPack(); err != nil {
			return err
		}
		for _, s := range a.Pack.Steps {
			fmt.Printf("%s: %s\n", s.ID, s.OriginalLine)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(listCmd)
}
