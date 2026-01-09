package cli

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"github.com/user/oraclepack/internal/app"
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
			OutputVerify:  outputVerify,
			OutputRetries: outputRetries,
		}

		a := app.New(cfg)
		// Prepare the application (loads pack, resolves out_dir, provisions env)
		if err := a.Prepare(); err != nil {
			return err
		}

		if err := a.LoadState(); err != nil {
			return err
		}

		if noTUI {
			return a.RunPlain(context.Background(), os.Stdout)
		}

		m := tui.NewModel(a.Pack, a.Runner, a.State, cfg.StatePath, cfg.ROIThreshold, cfg.ROIMode, runAll, cfg.OutputVerify, cfg.OutputRetries)
		p := tea.NewProgram(m, tea.WithAltScreen())
		_, err := p.Run()
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
	runCmd.Flags().BoolVar(&outputVerify, "output-verify", true, "Verify --write-output files contain required answer sections")
	runCmd.Flags().IntVar(&outputRetries, "output-retries", 1, "Retries for output verification failures")
	rootCmd.AddCommand(runCmd)
}
