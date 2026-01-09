package cli

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/app"
	"github.com/user/oraclepack/internal/config"
	"github.com/user/oraclepack/internal/pack"
	"github.com/user/oraclepack/internal/tui"
)

var (
	yes           bool
	resume        bool
	stopOnFail    bool
	roiThreshold  float64
	roiMode       string
	runAll        bool
	outputVerify  bool
	outputRetries int
)

var runCmd = &cobra.Command{
	Use:   "run [pack.md]",
	Short: "Run an oracle pack",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		packPath := args[0]

		// Setup paths
		base := strings.TrimSuffix(filepath.Base(packPath), filepath.Ext(packPath))
		statePath := base + ".state.json"
		reportPath := base + ".report.json"

		resolvedVerify, err := config.ResolveOutputVerify(outputVerify, cmd.Flags().Changed("output-verify"))
		if err != nil {
			return err
		}
		resolvedRetries, err := config.ResolveOutputRetries(outputRetries, cmd.Flags().Changed("output-retries"))
		if err != nil {
			return err
		}

		cfg := app.Config{
			PackPath:      packPath,
			StatePath:     statePath,
			ReportPath:    reportPath,
			Resume:        resume,
			StopOnFail:    stopOnFail,
			WorkDir:       ".",
			OutDir:        outDir,
			ROIThreshold:  roiThreshold,
			ROIMode:       roiMode,
			OutputVerify:  resolvedVerify,
			OutputRetries: resolvedRetries,
		}

		a := app.New(cfg)
		// Prepare the application (loads pack, resolves out_dir, provisions env)
		if err := a.Prepare(); err != nil {
			return err
		}

		if err := a.LoadState(); err != nil {
			return err
		}

		findings, warning, err := pack.CheckPackScripts(a.Pack)
		if err != nil {
			return err
		}
		if warning != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "Warning: %s\n", warning)
		}
		if len(findings) > 0 {
			for _, finding := range findings {
				if finding.StepID != "" {
					fmt.Fprintf(cmd.OutOrStdout(), "Step %s line %d: %s\n", finding.StepID, finding.Line, finding.Message)
				} else {
					fmt.Fprintf(cmd.OutOrStdout(), "Line %d: %s\n", finding.Line, finding.Message)
				}
			}
			return fmt.Errorf("bash syntax validation failed")
		}

		if noTUI {
			out := cmd.OutOrStdout()
			fmt.Fprintf(out, "[Selected] %s\n", packPath)
			fmt.Fprintln(out, "[Ready] Parsed and validated pack")
			err := a.RunPlain(context.Background(), out)
			if err != nil {
				fmt.Fprintf(out, "[Completed] Failed: %v\n", err)
				return err
			}
			fmt.Fprintln(out, "[Completed] Success")
			return nil
		}

		m := tui.NewModel(a.Pack, a.Runner, a.State, cfg.StatePath, cfg.ROIThreshold, cfg.ROIMode, runAll, cfg.OutputVerify, cfg.OutputRetries)
		p := tea.NewProgram(m, tea.WithAltScreen())
		_, err = p.Run()
		return err
	},
}

func init() {
	runCmd.Flags().BoolVarP(&yes, "yes", "y", false, "Auto-approve all steps")
	runCmd.Flags().BoolVar(&resume, "resume", false, "Resume from last successful step")
	runCmd.Flags().BoolVar(&stopOnFail, "stop-on-fail", true, "Stop execution if a step fails")
	runCmd.Flags().Float64Var(&roiThreshold, "roi-threshold", 0.0, "Filter steps by ROI threshold")
	runCmd.Flags().StringVar(&roiMode, "roi-mode", "over", "ROI filter mode ('over' or 'under')")
	runCmd.Flags().BoolVar(&runAll, "run-all", false, "Automatically run all steps sequentially on start")
	runCmd.Flags().BoolVar(&outputVerify, "output-verify", config.DefaultOutputVerify, "Verify --write-output files contain required answer sections")
	runCmd.Flags().IntVar(&outputRetries, "output-retries", config.DefaultOutputRetries, "Retries for output verification failures")
	rootCmd.AddCommand(runCmd)
}
