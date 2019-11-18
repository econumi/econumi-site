package cli

import (
  "context"
  "flag"
  "fmt"
  "os"
  log "github.com/Sirupsen/logrus"

  "econumi.org/coin/blockchain"
  "econumi.org/coin/wallet"
  "github.com/google/subcommands"
)

type startNodeCmd struct {
  capitalize bool
}

func (*startNodeCmd) Name() string     { return "startnode" }
func (*startNodeCmd) Synopsis() string { return "Start blockchain node" }
func (*startNodeCmd) Usage() string {
  return `startnode :
  Start blockchain node.
`
}

func (p *startNodeCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *startNodeCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  p.startNode()
  return subcommands.ExitSuccess
}

// TODO: https://docs.docker.com/engine/security/https/
func (p *startNodeCmd) startNode() {
  nodeID := os.Getenv("NODE_ID")
  wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
  minerAddress := wallets.GetDefaultWalletAddress()
  fmt.Printf("address = %s", minerAddress)

  fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if wallet.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	blockchain.StartServer(nodeID, minerAddress)
}
