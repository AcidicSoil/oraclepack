package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/config"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/types"
)

var verifyOutputsEnabled bool

var verifyOutputsCmd = &cobra.Command{
	Use:   "verify-outputs [pack.md]",
	Short: "Verify --write-output files without executing steps",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		out := cmd.OutOrStdout()
		data, err := os.ReadFile(args[0])
		if err != nil {
			return err
		}

		p, err := pack.Parse(data)
		if err != nil {
			return err
		}
		if err := pack.Validate(p); err != nil {
			return err
		}

		verifyEnabled, err := config.ResolveOutputVerify(verifyOutputsEnabled, cmd.Flags().Changed("output-verify"))
		if err != nil {
			return err
		}
		if !verifyEnabled {
			fmt.Fprintln(out, "Output verification disabled (ORACLEPACK_OUTPUT_VERIFY=false).")
			return nil
		}

		report := pack.VerifyReport{
			TotalSteps: len(p.Steps),
		}

		for i := range p.Steps {
			step := &p.Steps[i]
			expectations := pack.StepOutputExpectations(step)
			if len(expectations) == 0 {
				continue
			}
			report.CheckedSteps++
			for path, required := range expectations {
				if _, err := os.Stat(path); err != nil {
					report.Failures = append(report.Failures, types.OutputFailure{
						StepID: step.ID,
						Path:   path,
						Error:  err.Error(),
					})
					continue
				}
				ok, failure := pack.ValidateOutputFile(path, required)
				if !ok {
					failure.StepID = step.ID
					report.Failures = append(report.Failures, failure)
				}
			}
		}

		fmt.Fprint(out, pack.FormatVerifyReport(report))
		if len(report.Failures) > 0 {
			return fmt.Errorf("output verification failed")
		}
		return nil
	},
}

func init() {
	verifyOutputsCmd.Flags().BoolVar(&verifyOutputsEnabled, "output-verify", config.DefaultOutputVerify, "Verify --write-output files contain required answer sections")
	rootCmd.AddCommand(verifyOutputsCmd)
}
