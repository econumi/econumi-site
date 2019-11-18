package cli

import (
  "context"
  "flag"
  "fmt"
  "os"
  // log "github.com/Sirupsen/logrus"

  "github.com/google/subcommands"
  "econumi.org/coin/wallet"
)

type createWalletCmd struct {
  defaultWallet bool
}

func (*createWalletCmd) Name() string     { return "createwallet" }
func (*createWalletCmd) Synopsis() string { return "Create wallet on blockchain" }
func (*createWalletCmd) Usage() string {
  return `createwallet :
  Create a wallet on the blockchain.
`
}

func (p *createWalletCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.defaultWallet, "default", false, "set as default wallet")
}

func (p *createWalletCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  p.createWallet()
  return subcommands.ExitSuccess
}

// TODO: https://docs.docker.com/engine/security/https/
func (p *createWalletCmd) createWallet() {
  nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Printf("NODE_ID env. var is not set!")
		os.Exit(1)
	}

  wallets, _ := wallet.NewWallets(nodeID)
  var address string
  // TODO: save to wallet file
  if p.defaultWallet {
    address = wallets.CreateDefaultWallet()
    fmt.Printf("Your new default address: %s\n", address)
  } else {
    address = wallets.CreateWallet()
    fmt.Printf("Your new address: %s\n", address)
  }
	wallets.SaveToFile(nodeID)
}
