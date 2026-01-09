package tui

import (
	"context"
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/oraclepack/internal/exec"
	"github.com/user/oraclepack/internal/overrides"
	"github.com/user/oraclepack/internal/types"
)

type OverridesStep int

const (
	OverridesFlags OverridesStep = iota
	OverridesSteps
	OverridesConfirm
)

type OverridesStartedMsg struct{}

type OverridesAppliedMsg struct {
	Overrides overrides.RuntimeOverrides
}

type OverridesCancelledMsg struct{}

type OverridesFlowModel struct {
	step    OverridesStep
	flags   FlagsPickerModel
	steps   StepsPickerModel
	confirm OverridesConfirmModel

	packSteps        []types.Step
	baseline         []string
	runnerOpts       exec.RunnerOptions
	pendingOverrides overrides.RuntimeOverrides
}

func NewOverridesFlowModel(steps []types.Step, baseline []string, opts exec.RunnerOptions) OverridesFlowModel {
	return OverridesFlowModel{
		step:       OverridesFlags,
		flags:      NewFlagsPickerModel(nil),
		steps:      NewStepsPickerModel(steps),
		confirm:    OverridesConfirmModel{},
		packSteps:  steps,
		baseline:   exec.ApplyChatGPTURL(baseline, opts.ChatGPTURL),
		runnerOpts: opts,
	}
}

func (m OverridesFlowModel) Init() tea.Cmd {
	return nil
}

func (m OverridesFlowModel) Update(msg tea.Msg) (OverridesFlowModel, tea.Cmd) {
	var cmd tea.Cmd
	if m.step == OverridesFlags {
		m.flags, cmd = m.flags.Update(msg)
	}
	if m.step == OverridesSteps {
		m.steps, cmd = m.steps.Update(msg)
	}
	if m.step == OverridesConfirm {
		switch v := msg.(type) {
		case ValidationResultMsg:
			m.confirm.validating = false
			m.confirm.errors = v.Errors
			if v.Err != nil {
				m.confirm.errMsg = v.Err.Error()
				return m, nil
			}
			if len(v.Errors) > 0 {
				m.confirm.errMsg = fmt.Sprintf("%d validation errors detected.", len(v.Errors))
				return m, nil
			}
			return m, func() tea.Msg { return OverridesAppliedMsg{Overrides: m.pendingOverrides} }
		}
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return OverridesCancelledMsg{} }
		case "shift+tab", "backspace":
			if m.step > OverridesFlags {
				m.step--
			}
		case "enter", "tab":
			if m.step == OverridesConfirm {
				if m.confirm.validating {
					return m, nil
				}
				m.pendingOverrides = m.currentOverrides()
				m.confirm.validating = true
				m.confirm.errMsg = ""
				m.confirm.errors = nil
				return m, m.validateCmd(m.pendingOverrides)
			}
			m.step++
		}
	}

	return m, cmd
}

func (m OverridesFlowModel) View(width, height int) string {
	title := lipgloss.NewStyle().Bold(true).Render("Overrides Wizard")
	step := fmt.Sprintf("Step %d/3", int(m.step)+1)
	body := fmt.Sprintf("Current step: %s\n\n[Enter] Next  [Esc] Cancel", overridesStepName(m.step))

	var content string
	if m.step == OverridesFlags {
		m.flags.SetSize(width-4, height-8)
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.flags.View(),
			"",
			body,
		)
	} else if m.step == OverridesSteps {
		m.steps.SetSize(width-4, height-8)
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.steps.View(),
			"",
			body,
		)
	} else if m.step == OverridesConfirm {
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			m.confirm.View(m.currentOverrides(), m.baseline),
		)
	} else {
		content = lipgloss.JoinVertical(lipgloss.Left,
			title,
			step,
			"",
			body,
		)
	}

	return lipgloss.Place(width, height, lipgloss.Center, lipgloss.Center, content)
}

func (m OverridesFlowModel) currentOverrides() overrides.RuntimeOverrides {
	return overrides.RuntimeOverrides{
		AddedFlags:   m.flags.SelectedFlags(),
		RemovedFlags: nil,
		ApplyToSteps: m.steps.SelectedSteps(),
	}
}

func (m OverridesFlowModel) validateCmd(over overrides.RuntimeOverrides) tea.Cmd {
	return func() tea.Msg {
		errs, err := exec.ValidateOverrides(context.Background(), m.packSteps, &over, m.baseline, m.runnerOpts)
		return ValidationResultMsg{Errors: errs, Err: err}
	}
}

func overridesStepName(step OverridesStep) string {
	switch step {
	case OverridesFlags:
		return "Flags"
	case OverridesSteps:
		return "Target Steps"
	case OverridesConfirm:
		return "Confirm"
	default:
		return "Unknown"
	}
}
