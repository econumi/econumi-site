package cli

import (
  "context"
  "flag"
  "fmt"
  "strconv"
  "unicode/utf8"

  "golang.org/x/text/width"
  "github.com/docker/docker/pkg/stringid"
  "github.com/docker/docker/api/types"
  "github.com/docker/docker/client"
  "github.com/google/subcommands"
)

type psCmd struct {
  capitalize bool
}

func (*psCmd) Name() string     { return "ps" }
func (*psCmd) Synopsis() string { return "Show containers running." }
func (*psCmd) Usage() string {
  return `ps [-capitalize]:
  Show containers running
`
}

func (p *psCmd) SetFlags(f *flag.FlagSet) {
  f.BoolVar(&p.capitalize, "capitalize", false, "capitalize output")
}

func (p *psCmd) Execute(_ context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
  listContainers()
  return subcommands.ExitSuccess
}

func listContainers() {
  cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

  fmt.Println("CONTAINER ID\tIMAGE\t\tCOMMAND\t\t\t\tSTATUS")
	for _, container := range containers {
    fmt.Printf("%v\t%v\t\t%v\t\t%v\n", stringid.TruncateID(container.ID), container.Image, strconv.Quote(Ellipsis(container.Command, 20)), container.Status)
	}
}

// NOTE: from github.com/docker/cli
// charWidth returns the number of horizontal positions a character occupies,
// and is used to account for wide characters when displaying strings.
func charWidth(r rune) int {
	switch width.LookupRune(r).Kind() {
	case width.EastAsianWide, width.EastAsianFullwidth:
		return 2
	default:
		return 1
	}
}

// NOTE: from github.com/docker/cli
// Ellipsis truncates a string to fit within maxDisplayWidth, and appends ellipsis (…).
// For maxDisplayWidth of 1 and lower, no ellipsis is appended.
// For maxDisplayWidth of 1, first char of string will return even if its width > 1.
func Ellipsis(s string, maxDisplayWidth int) string {
	if maxDisplayWidth <= 0 {
		return ""
	}
	rs := []rune(s)
	if maxDisplayWidth == 1 {
		return string(rs[0])
	}

	byteLen := len(s)
	if byteLen == utf8.RuneCountInString(s) {
		if byteLen <= maxDisplayWidth {
			return s
		}
		return string(rs[:maxDisplayWidth-1]) + "…"
	}

	var (
		display      []int
		displayWidth int
	)
	for _, r := range rs {
		cw := charWidth(r)
		displayWidth += cw
		display = append(display, displayWidth)
	}
	if displayWidth <= maxDisplayWidth {
		return s
	}
	for i := range display {
		if display[i] <= maxDisplayWidth-1 && display[i+1] > maxDisplayWidth-1 {
			return string(rs[:i+1]) + "…"
		}
	}
	return s
}
