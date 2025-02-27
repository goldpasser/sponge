// Package commands are subcommands of the sponge command.
package commands

import (
	"fmt"
	"os"

	"github.com/zhufuyi/sponge/cmd/sponge/commands/generate"

	"github.com/spf13/cobra"
)

var (
	version     = "v0.0.0"
	versionFile = GetSpongeDir() + "/.sponge/.github/version"
)

// NewRootCMD command entry
func NewRootCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sponge",
		Long: `sponge is a powerful tool for generating web and microservice code, a microservice framework
based on gin and grpc encapsulation, and an open source framework for rapid application development.`,
		SilenceErrors: true,
		SilenceUsage:  true,
		Version:       getVersion(),
	}

	cmd.AddCommand(
		InitCommand(),
		UpgradeCommand(),
		ToolsCommand(),
		NewWebCommand(),
		MicroCommand(),
		generate.ConfigCommand(),
		NewRunCommand(),
		generate.DeleteJSONOmitemptyCommand(),
		MergeCommand(),
	)

	return cmd
}

func getVersion() string {
	data, _ := os.ReadFile(versionFile)
	v := string(data)
	if v != "" {
		return v
	}
	return "unknown, execute command 'sponge init' to get version"
}

// GetSpongeDir get sponge home directory
func GetSpongeDir() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("can't get home directory'")
		return ""
	}

	return dir
}
