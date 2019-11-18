package cli

import (
  "flag"
  "os"

  "github.com/google/subcommands"
  context "golang.org/x/net/context"
)

type CLI struct {}

func (cli *CLI) Run() {
	subcommands.Register(subcommands.HelpCommand(), "")
  subcommands.Register(subcommands.FlagsCommand(), "")
  subcommands.Register(subcommands.CommandsCommand(), "")
  subcommands.Register(&startCmd{}, "")
	subcommands.Register(&psCmd{}, "")
  subcommands.Register(&imagesCmd{}, "")
  subcommands.Register(&createWalletCmd{}, "")
  subcommands.Register(&createBlockchainCmd{}, "")
  subcommands.Register(&startNodeCmd{}, "")
  subcommands.Register(&printChainCmd{}, "")
  subcommands.Register(&getBalanceCmd{}, "")
  subcommands.Register(&listAddressesCmd{}, "")
  subcommands.Register(&sendCmd{}, "")

	flag.Parse()
  ctx := context.Background()
  os.Exit(int(subcommands.Execute(ctx)))
}
