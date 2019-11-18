package cli

import (
  "context"
  "flag"
  "fmt"
  "os"
  "strconv"
  // log "github.com/Sirupsen/logrus"

  "econumi.org/coin/blockchain"
  "github.com/google/subcommands"
)

type printChainCmd struct {
  capitalize bool
}

func (*printChainCmd) Name() string     { return "printchain" }
func (*printChainCmd) Synopsis() string { return "Print the current blockchain" }
func (*printChainCmd) Usage() string {
  return `printchain :
  List images to run in the blockchain.
`
}

func (p *printChainCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *printChainCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  printChain()
  return subcommands.ExitSuccess
}

// TODO: https://docs.docker.com/engine/security/https/
func printChain() {
  nodeID := os.Getenv("NODE_ID")
  bc := blockchain.NewBlockchain(nodeID)
	defer bc.Db.Close()

	bci := bc.Iterator()

	for {
		block := bci.Next()

		fmt.Printf("============ Block %x ============\n", block.Hash)
		fmt.Printf("Prev. block: %x\n", block.PrevBlockHash)
		pow := blockchain.NewProofOfWork(block)
		fmt.Printf("PoW: %s\n\n", strconv.FormatBool(pow.Validate()))
		for _, tx := range block.Transactions {
			fmt.Println(tx)
		}
		fmt.Printf("\n\n")

		if len(block.PrevBlockHash) == 0 {
			break
		}
	}
}
