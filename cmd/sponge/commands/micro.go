package commands

import (
	"github.com/zhufuyi/sponge/cmd/sponge/commands/generate"

	"github.com/spf13/cobra"
)

// MicroCommand micro commands
func MicroCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "micro",
		Short:         "Generate proto, model, cache, dao, service, rpc, rpc-gw, rpc-cli codes",
		Long:          "generate proto, model, cache, dao, service, rpc, rpc-gw, rpc-cli codes.",
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.AddCommand(
		generate.ProtoBufCommand(),
		generate.ModelCommand("micro"),
		generate.DaoCommand("micro"),
		generate.CacheCommand("micro"),
		generate.ServiceCommand(),
		generate.RPCCommand(),
		generate.RPCGwPbCommand(),
		generate.RPCPbCommand(),
		generate.RPCConnectionCommand(),
		generate.ConvertSwagJSONCommand("micro"),
	)

	return cmd
}
