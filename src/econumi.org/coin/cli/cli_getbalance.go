package cli

import (
  "context"
  "flag"
  "fmt"
  "os"
  log "github.com/Sirupsen/logrus"

  "github.com/google/subcommands"
  "econumi.org/coin/util"
	"econumi.org/coin/blockchain"
  "econumi.org/coin/wallet"
)

type getBalanceCmd struct {
  defaultWallet bool
}

func (*getBalanceCmd) Name() string     { return "getbalance" }
func (*getBalanceCmd) Synopsis() string { return "Get balance on the default wallet where the CLI is installed on" }
func (*getBalanceCmd) Usage() string {
  return `getbalance :
  Get the balance on the blockchain address.
`
}

func (p *getBalanceCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.defaultWallet, "default", false, "default address")
}

func (p *getBalanceCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  args := f.Args()
  if(len(args) > 0) {
    for _, arg := range args {
      p.getBalance(arg)
    }
  } else {
    nodeID := os.Getenv("NODE_ID")
    wallets, err := wallet.NewWallets(nodeID)
    if err != nil {
  		log.Panic(err)
  	}
    p.getBalance(wallets.GetDefaultWalletAddress())
  }
  return subcommands.ExitSuccess
}

func (p *getBalanceCmd) getBalance(address string) {
  nodeID := os.Getenv("NODE_ID")
  if !wallet.ValidateAddress(address) {
    log.Panic("ERROR: Address is not valid")
  }
  bc := blockchain.NewBlockchain(nodeID)
  UTXOSet := blockchain.UTXOSet{bc}
  defer bc.Db.Close()

  balance := 0
  pubKeyHash := util.Base58Decode([]byte(address))
  pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
  UTXOs := UTXOSet.FindUTXO(pubKeyHash)

  for _, out := range UTXOs {
    balance += out.Value
  }

  fmt.Printf("Balance of '%s': %d\n", address, balance)
}
