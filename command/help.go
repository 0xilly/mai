package command

import "fmt"

type HelpCommand struct {
	commands map[string]Command
}

func NewHelpCommand(commands map[string]Command) *HelpCommand {
	return &HelpCommand{commands: commands}
}

func (h *HelpCommand) Name() string        { return "help" }
func (h *HelpCommand) ShortName() string   { return "h" }
func (h *HelpCommand) Description() string { return "Show help for commands" }
func (h *HelpCommand) Usage() string {
	return "help [<command>] - Show general help or help for a specific command"
}

func (h *HelpCommand) Run(args []string) error {
	if len(args) == 0 {
		fmt.Println("Available commands:")
		seen := make(map[string]bool)
		for _, cmd := range h.commands {
			// Avoid duplicate short/long name entries
			name := cmd.Name()
			if seen[name] {
				continue
			}
			seen[name] = true
			fmt.Printf("  %-12s (%s) - %s\n", name, cmd.ShortName(), cmd.Description())
		}
		return nil
	}

	name := args[0]
	cmd, ok := h.commands[name]
	if !ok {
		// Attempt to resolve by short name
		for _, c := range h.commands {
			if c.ShortName() == name {
				cmd = c
				ok = true
				break
			}
		}
	}
	if !ok {
		return fmt.Errorf("no such command: %s", name)
	}

	fmt.Printf("Usage: %s\n\n", cmd.Usage())
	if desc := cmd.Description(); desc != "" {
		fmt.Println(desc)
	}
	return nil
}
