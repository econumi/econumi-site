package cli

import (
  "context"
  "flag"
  "fmt"
  "os"

  "econumi.org/coin/blockchain"
  "github.com/google/subcommands"
)

type reindexUtxCmd struct {
  capitalize bool
}

func (*reindexUtxCmd) Name() string     { return "reindexutx" }
func (*reindexUtxCmd) Synopsis() string { return "Reindex UTXO set" }
func (*reindexUtxCmd) Usage() string {
  return `reindexutx :
  Reindex UTXO set.
`
}

func (p *reindexUtxCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *reindexUtxCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  p.ReindexUTX()
  return subcommands.ExitSuccess
}

// TODO: https://docs.docker.com/engine/security/https/
func (p *reindexUtxCmd) ReindexUTX() {
  nodeID := os.Getenv("NODE_ID")
  bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in the UTXO set.\n", count)
}
