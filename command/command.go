package command

import (
	"fmt"
	"strings"
)

type Command interface {
	Name() string
	ShortName() string
	Run(args []string) error
	Usage() string
	Description() string
}

type commandRegistry struct {
	commands map[string]Command
}

func InitCommandRegistry() *commandRegistry {
	return &commandRegistry{
		commands: make(map[string]Command),
	}
}

func (r *commandRegistry) Register(cmd Command) {
	r.commands[cmd.Name()] = cmd
	if short := cmd.ShortName(); short != "" {
		r.commands[short] = cmd
	}
}

func (r *commandRegistry) Run(name string, args []string) {
	cmd, ok := r.commands[name]
	if !ok {
		fmt.Printf("Unknown command: %s\n", name)
		return
	}

	normalizedArgs := normalizeFlags(args)
	if err := cmd.Run(normalizedArgs); err != nil {
		fmt.Printf("Error: %v\n", err)
		fmt.Println(cmd.Usage())
	}
}

func (r *commandRegistry) LoadCommands() {
	r.Register(NewHelpCommand(r.commands))
	r.Register(NewBuildCommand())
	r.Register(NewInitBuild())
}

func normalizeFlags(args []string) []string {
	var out []string
	for _, arg := range args {
		if strings.HasPrefix(arg, "-O") && len(arg) > 2 {
			out = append(out, "-O", arg[2:])
		} else {
			out = append(out, arg)
		}
	}
	return out
}
