package cli

import (
	"context"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/oraclepack/internal/app"
	"github.com/user/oraclepack/internal/tui"
)

var (
	yes          bool
	resume       bool
	stopOnFail   bool
	roiThreshold float64
	roiMode      string
	runAll       bool
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
			PackPath:     packPath,
			StatePath:    statePath,
			ReportPath:   reportPath,
			Resume:       resume,
			StopOnFail:   stopOnFail,
			WorkDir:      ".",
			ROIThreshold: roiThreshold,
			ROIMode:      roiMode,
		}

		a := app.New(cfg)
		if err := a.LoadPack(); err != nil {
			return err
		}
		if err := a.LoadState(); err != nil {
			return err
		}

		if noTUI {
			return a.RunPlain(context.Background(), os.Stdout)
		}

		m := tui.NewModel(a.Pack, a.Runner, a.State, cfg.ROIThreshold, cfg.ROIMode, runAll)
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
	rootCmd.AddCommand(runCmd)
}
