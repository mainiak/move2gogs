package main

import (
	"fmt"
	"github.com/mkideal/cli"
)

type ArgT struct {
	cli.Helper
	Organization string `cli:"o,org" usage:"Organization name"`
	CreateOrg    bool   `cli:"create-org" usage:"Create organization if it doesn't exist"`
	TokenFile    string `cli:"token-file" usage:"Path to file with API token for Gogs"`
	Repo         string `cli:"r,repo" usage:"Path to git repository"`
	Project      string `cli:"p,project" usage:"Project name - by default same as git directory name"`
}

func (argv *ArgT) Validate(ctx *cli.Context) error {
	if argv.TokenFile == "" {
		return fmt.Errorf("Path to file with API token for Gogs must be specified")
	}
	if argv.Repo == "" {
		return fmt.Errorf("Path to git repository must be specified")
	}
	if argv.CreateOrg && argv.Organization == "" {
		return fmt.Errorf("You can not use --create-org without --org \"someorg\"")
	}
	return nil
}

func main() {
	cli.Run(new(ArgT), func(ctx *cli.Context) error {
		return nil
	}, "Create organization and mirror git repo to Gogs server")
}
