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

type createBlockchainCmd struct {
  capitalize bool
}

func (*createBlockchainCmd) Name() string     { return "createblockchain" }
func (*createBlockchainCmd) Synopsis() string { return "Create wallet on blockchain" }
func (*createBlockchainCmd) Usage() string {
  return `createblockchain :
  Create a wallet on the blockchain.
`
}

func (p *createBlockchainCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *createBlockchainCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  createBlockchain()
  return subcommands.ExitSuccess
}

// TODO: https://docs.docker.com/engine/security/https/
func createBlockchain() {
  nodeID := os.Getenv("NODE_ID")
  wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
  address := wallets.GetDefaultWalletAddress()
  if !wallet.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := blockchain.CreateBlockchain(address, nodeID)
	defer bc.Db.Close()

	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
