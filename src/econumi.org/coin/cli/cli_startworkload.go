package cli

import (
  "context"
  "flag"
  "fmt"

  log "github.com/Sirupsen/logrus"

  "github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
  "github.com/google/subcommands"
)

type startCmd struct {
  capitalize bool
}

func (*startCmd) Name() string     { return "start" }
func (*startCmd) Synopsis() string { return "Start container in blockchain." }
func (*startCmd) Usage() string {
  return `start [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *startCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *startCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  for _, arg := range f.Args() {
    createStartContainer(arg)
  }
  fmt.Println()
  return subcommands.ExitSuccess
}

// TODO: https://docs.docker.com/engine/security/https/
func createStartContainer(image string) {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	r, errCli := cli.ContainerCreate(context.Background(), &container.Config{ Image: image }, nil, nil, "")
	if errCli != nil {
		log.Infof("%s", errCli)
	}
	if r.ID != "container_id" {
		log.Infof("expected `container_id`, got %s", r.ID)
	}

	if errStart := cli.ContainerStart(context.Background(), r.ID, types.ContainerStartOptions{}); err != nil {
		log.Infof("%s", errStart)
	}

}
