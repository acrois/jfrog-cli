package plugins

import (
	corecommon "github.com/jfrog/jfrog-cli-core/v2/docs/common"
	"github.com/jfrog/jfrog-cli/docs/common"
	installdocs "github.com/jfrog/jfrog-cli/docs/plugin/install"
	publishdocs "github.com/jfrog/jfrog-cli/docs/plugin/publish"
	uninstalldocs "github.com/jfrog/jfrog-cli/docs/plugin/uninstall"
	"github.com/jfrog/jfrog-cli/plugins/commands"
	"github.com/jfrog/jfrog-cli/utils/cliutils"
	"github.com/urfave/cli"
)

func GetCommands() []cli.Command {
	return cliutils.GetSortedCommands(cli.CommandsByName{
		{
			Name:         "install",
			Aliases:      []string{"i"},
			Usage:        installdocs.GetDescription(),
			HelpName:     corecommon.CreateUsage("plugin install", installdocs.GetDescription(), installdocs.Usage),
			UsageText:    installdocs.GetArguments(),
			ArgsUsage:    common.CreateEnvVars(installdocs.EnvVar...),
			BashComplete: corecommon.CreateBashCompletionFunc(),
			Action:       commands.InstallCmd,
		},
		{
			Name:         "uninstall",
			Aliases:      []string{"ui"},
			Usage:        uninstalldocs.GetDescription(),
			HelpName:     corecommon.CreateUsage("plugin uninstall", uninstalldocs.GetDescription(), uninstalldocs.Usage),
			UsageText:    uninstalldocs.GetArguments(),
			ArgsUsage:    common.CreateEnvVars(),
			BashComplete: corecommon.CreateBashCompletionFunc(),
			Action:       commands.UninstallCmd,
		},
		{
			Name:         "publish",
			Aliases:      []string{"p"},
			Flags:        cliutils.GetCommandFlags(cliutils.PluginPublish),
			Usage:        publishdocs.GetDescription(),
			HelpName:     corecommon.CreateUsage("plugin publish", publishdocs.GetDescription(), publishdocs.Usage),
			UsageText:    publishdocs.GetArguments(),
			ArgsUsage:    common.CreateEnvVars(publishdocs.EnvVar...),
			BashComplete: corecommon.CreateBashCompletionFunc(),
			Action:       commands.PublishCmd,
		},
	})
}
