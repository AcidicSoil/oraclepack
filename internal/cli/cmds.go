package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/app"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/validate"
)

var validateCmd = &cobra.Command{
	Use:   "validate [pack.md]",
	Short: "Validate an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		data, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}
		p, err := pack.Parse(data)
		if err != nil {
			return err
		}
		if err := p.Validate(); err != nil {
			return err
		}
		cv := validate.CompositeValidator{}
		results := cv.ValidatePack(p)
		out := cmd.OutOrStdout()
		fmt.Fprintf(out, "Validated %d steps\n", len(results))
		for _, r := range results {
			fmt.Fprintf(out, "Step %s [%s] %s", r.StepID, r.ToolKind.Name(), r.Status)
			if r.Error != "" {
				fmt.Fprintf(out, " (%s)", r.Error)
			}
			fmt.Fprintln(out)
		}
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
