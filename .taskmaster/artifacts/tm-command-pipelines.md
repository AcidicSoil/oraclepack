~~~
task-master use-tag bugfix-tui → task-master list --tag=bugfix-tui
~~~
~~~
if [ -f .taskmaster/tasks/tasks.json ]; then task-master list --tag=bugfix-tui; else task-master init; fi
~~~
~~~
task-master show 7 --tag=bugfix-tui → ck --sem "textinput validation" internal/tui/overrides_url.go → go test ./internal/tui → task-master update-subtask --id=7.2 --prompt="Verified URL input validation" --tag=bugfix-tui
~~~
~~~
task-master show 8 --tag=bugfix-tui → ck --regex "OverridesAppliedMsg" internal/tui/overrides_confirm.go → go test ./internal/tui → task-master set-status --id=8.1 --status=done --tag=bugfix-tui
~~~
~~~
task-master show 9 --tag=bugfix-tui → ck --sem "EffectiveFlags" internal/exec/runner.go → go test ./internal/exec → task-master update-subtask --id=9.1 --prompt="Integrated effective flags into execution loop" --tag=bugfix-tui
~~~
~~~
for id in 10.1 10.2 10.3 10.4 10.5; do task-master show $id --tag=bugfix-tui → task-master set-status --id=$id --status=done --tag=bugfix-tui; done
~~~
~~~
go test ./internal/tui || task-master show 7 --tag=bugfix-tui → task-master update-subtask --id=7.1 --prompt="Tests failed, investigation needed" --tag=bugfix-tui
~~~
~~~
task-master analyze-complexity --tag=bugfix-tui & task-master validate-dependencies --tag=bugfix-tui → wait → task-master generate --tag=bugfix-tui
~~~
~~~
task-master list --status=done --tag=bugfix-tui → task-master sync-readme --with-subtasks --tag=bugfix-tui
~~~
~~~
ck --hybrid "OverridesActive" internal/tui/tui.go → task-master show 10 --tag=bugfix-tui → task-master set-status --id=10.2 --status=done --tag=bugfix-tui
~~~
~~~
ast-grep --pattern 'struct Model { $$$ autoRun bool $$$ }' internal/tui/model.go → task-master update-subtask --id=14.2 --prompt="Verified autoRun field structure" --tag=bugfix-tui
~~~
~~~
task-master --help → task-master show 1 --tag=bugfix-tui
~~~
~~~
task-master fix-dependencies --tag=bugfix-tui → task-master generate --tag=bugfix-tui
~~~
~~~
go build ./cmd/oraclepack → ./oraclepack run --help → task-master show 10 --tag=bugfix-tui
~~~
~~~
cat .taskmaster/state.json → task-master list --tag=bugfix-tui
~~~
~~~
ck --regex "EffectiveFlags" internal/overrides/merge.go → go test ./internal/overrides → task-master list --tag=bugfix-tui
~~~
~~~
ck --regex "ExtractOracleInvocations" internal/exec/oracle_scan.go → go test ./internal/exec → task-master list --tag=bugfix-tui
~~~
~~~
task-master research "Bubble Tea text input validation best practices" --tag=bugfix-tui → task-master update-subtask --id=7.2 --prompt="Research validation patterns" --tag=bugfix-tui
~~~
~~~
task-master sync-readme --with-subtasks --tag=bugfix-tui → git status
~~~
~~~
task-master next --tag=bugfix-tui → task-master show 10 --tag=bugfix-tui
~~~