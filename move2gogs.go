package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"os"
)

type ArgT struct {
	cli.Helper
	Server       string `cli:"s,server" usage:"Server URI - URL to your Gogs instance"`
	Organization string `cli:"o,org" usage:"Organization name"`
	CreateOrg    bool   `cli:"create-org" usage:"Create organization if it doesn't exist"`
	TokenFile    string `cli:"token-file" usage:"Path to file with API token for Gogs"`
	Repo         string `cli:"r,repo" usage:"Path to git repository"`
	Project      string `cli:"p,project" usage:"Project name - by default same as git directory name"`
}

func (argv *ArgT) Validate(ctx *cli.Context) error {
	if argv.Server == "" {
		return fmt.Errorf("Server URI must be specified")
	}

	if argv.TokenFile == "" {
		return fmt.Errorf("Path to file with API token for Gogs must be specified")
	}

	fi, err := os.Stat(argv.TokenFile)
	if os.IsNotExist(err) {
		return fmt.Errorf("Following path doesn't exist: %v", argv.TokenFile)
	}

	if !fi.Mode().IsRegular() {
		return fmt.Errorf("Following path isn't file: %v", argv.TokenFile)
	}

	if argv.Repo == "" {
		return fmt.Errorf("Path to git repository must be specified")
	}

	fi, err = os.Stat(argv.Repo)
	if os.IsNotExist(err) {
		return fmt.Errorf("Following path doesn't exist: %v", argv.Repo)
	}

	if !fi.Mode().IsDir() {
		return fmt.Errorf("Following path isn't directory: %v", argv.Repo)
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
