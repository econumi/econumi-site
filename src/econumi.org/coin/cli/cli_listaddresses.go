package cli

import (
  "context"
  "flag"
  "fmt"
  "os"
  log "github.com/Sirupsen/logrus"

  "github.com/google/subcommands"
  "econumi.org/coin/wallet"
)

type listAddressesCmd struct {
}

func (*listAddressesCmd) Name() string     { return "listaddresses" }
func (*listAddressesCmd) Synopsis() string { return "Get the list of addresses on the blockchain" }
func (*listAddressesCmd) Usage() string {
  return `listaddresses :
  Get the list of addresses on the blockchain.`
}

func (p *listAddressesCmd) SetFlags(f *flag.FlagSet) {

}

func (p *listAddressesCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  p.getAddresses()
  return subcommands.ExitSuccess
}

func (p *listAddressesCmd) getAddresses() {
  nodeID := os.Getenv("NODE_ID")
  wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()
	for _, address := range addresses {
		fmt.Println(address)
	}
}
