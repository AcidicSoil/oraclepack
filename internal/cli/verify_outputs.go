package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/config"
	"github.com/user/oraclepack/internal/pack"
)

var (
	verifyOutputsEnabled      bool
	verifyOutputsRequireHeads bool
	verifyOutputsChunkMode    string
)

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
		requireHeadings, err := config.ResolveOutputRequireHeadings(verifyOutputsRequireHeads, cmd.Flags().Changed("output-require-headings"))
		if err != nil {
			return err
		}
		chunkMode, err := config.ResolveOutputChunkMode(verifyOutputsChunkMode, cmd.Flags().Changed("output-chunk-mode"))
		if err != nil {
			return err
		}

		report := pack.VerifyReport{
			TotalSteps: len(p.Steps),
		}

		for i := range p.Steps {
			step := &p.Steps[i]
			failures := pack.VerifyStepOutputs(step, requireHeadings, chunkMode)
			if len(failures) == 0 {
				continue
			}
			report.CheckedSteps++
			for _, failure := range failures {
				failure.StepID = step.ID
				report.Failures = append(report.Failures, failure)
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
	verifyOutputsCmd.Flags().BoolVar(&verifyOutputsRequireHeads, "output-require-headings", config.DefaultOutputRequireHeadings, "Require strict output headings when verifying outputs")
	verifyOutputsCmd.Flags().StringVar(&verifyOutputsChunkMode, "output-chunk-mode", config.DefaultOutputChunkMode, "Output chunk verification mode: auto|single|multi")
	rootCmd.AddCommand(verifyOutputsCmd)
}
