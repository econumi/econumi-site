package cli

import (
  "context"
  "flag"
  "fmt"

  // log "github.com/Sirupsen/logrus"
  // "github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
  "github.com/google/subcommands"
)

type stopCmd struct {
  capitalize bool
}

func (*stopCmd) Name() string     { return "stop" }
func (*stopCmd) Synopsis() string { return "Stop container in blockchain." }
func (*stopCmd) Usage() string {
  return `stop [-capitalize] <some text>:
  Print args to stdout.
`
}

func (p *stopCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *stopCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  for _, arg := range f.Args() {
    stopContainer(arg)
  }
  fmt.Println()
  return subcommands.ExitSuccess
}

// TODO: https://docs.docker.com/engine/security/https/
func stopContainer(containerId string) {
  ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	fmt.Print("Stopping container ", containerId, "... ")
	if err := cli.ContainerStop(ctx, containerId, nil); err != nil {
		panic(err)
	}
	fmt.Println("Success")

}
