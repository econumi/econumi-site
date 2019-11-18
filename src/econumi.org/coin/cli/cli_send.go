package cli

import (
  "context"
  "flag"
  "fmt"
  "os"
  "strconv"
  log "github.com/Sirupsen/logrus"

  "econumi.org/coin/blockchain"
  "github.com/google/subcommands"
  "econumi.org/coin/wallet"
)

type sendCmd struct {
}

func (*sendCmd) Name() string     { return "send" }
func (*sendCmd) Synopsis() string { return "Send tokens from one address to another" }
func (*sendCmd) Usage() string {
  return `send :
  Send tokens from one address to another.`
}

func (p *sendCmd) SetFlags(f *flag.FlagSet) {

}

func (p *sendCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  //TODO: Check arguments for from, to, and amount
  args := f.Args()
  if(len(args) == 3) {
    p.send(args[0], args[1], args[1], false)
  } else  {
    fmt.Println("Please enter correct number of arguments")
    return subcommands.ExitUsageError
  }
  return subcommands.ExitSuccess
}

func (p *sendCmd) send(to, from, amount string, mineNow bool) {
  nodeID := os.Getenv("NODE_ID")

  if !wallet.ValidateAddress(from) {
		log.Panic("ERROR: Sender address is not valid")
	}
	if !wallet.ValidateAddress(to) {
		log.Panic("ERROR: Recipient address is not valid")
	}

  bc := blockchain.NewBlockchain(nodeID)
	UTXOSet := blockchain.UTXOSet{bc}
	defer bc.Db.Close()

  wallets, err := wallet.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	wallet := wallets.GetWallet(from)

  // floatAmount, _ := strconv.ParseFloat(amount, 64)
  intAmount, _ := strconv.ParseInt(amount, 10, 64)
	tx := blockchain.NewUTXOTransaction(&wallet, to, int(intAmount), &UTXOSet)

  if mineNow {
		cbTx := blockchain.NewCoinbaseTX(from, "")
		txs := []*blockchain.Transaction{cbTx, tx}

		newBlock := bc.MineBlock(txs)
		UTXOSet.Update(newBlock)
	} else {
		blockchain.SendTx(blockchain.KnownNodes[0], tx)
	}
	fmt.Println("Success!")
}
