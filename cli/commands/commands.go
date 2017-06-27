package commands

import (
	"github.com/sensu/sensu-go/cli"
	"github.com/sensu/sensu-go/cli/commands/asset"
	"github.com/sensu/sensu-go/cli/commands/check"
	"github.com/sensu/sensu-go/cli/commands/completion"
	"github.com/sensu/sensu-go/cli/commands/configure"
	"github.com/sensu/sensu-go/cli/commands/entity"
	"github.com/sensu/sensu-go/cli/commands/event"
	"github.com/sensu/sensu-go/cli/commands/handler"
	"github.com/sensu/sensu-go/cli/commands/organization"
	"github.com/sensu/sensu-go/cli/commands/user"
	"github.com/spf13/cobra"
)

// AddCommands adds management commands to given command
func AddCommands(rootCmd *cobra.Command, cli *cli.SensuCli) {
	rootCmd.AddCommand(
		configure.Command(cli),
		completion.Command(rootCmd),

		// Management Commands
		asset.HelpCommand(cli),
		check.HelpCommand(cli),
		configure.HelpCommand(cli),
		entity.HelpCommand(cli),
		event.HelpCommand(cli),
		handler.HelpCommand(cli),
		organization.HelpCommand(cli),
		user.HelpCommand(cli),
	)

	for _, cmd := range rootCmd.Commands() {
		rootCmd.ValidArgs = append(rootCmd.ValidArgs, cmd.Use)
	}
}
