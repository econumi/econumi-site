package cli

import (
  "context"
  "flag"
  "fmt"
  // log "github.com/Sirupsen/logrus"

  "github.com/docker/docker/pkg/stringid"
  "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
  "github.com/google/subcommands"
)

type imagesCmd struct {
  capitalize bool
}

func (*imagesCmd) Name() string     { return "images" }
func (*imagesCmd) Synopsis() string { return "List container images to run in blockchain" }
func (*imagesCmd) Usage() string {
  return `images :
  List images to run in the blockchain.
`
}

func (p *imagesCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *imagesCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  listImages()
  return subcommands.ExitSuccess
}

// TODO: https://docs.docker.com/engine/security/https/
func listImages() {
  cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

  fmt.Println("IMAGE ID\tIMAGE\t\tTAG")
	for _, image := range images {
		fmt.Println(stringid.TruncateID(image.ID))
	}
}
