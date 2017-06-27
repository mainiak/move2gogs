package main

import (
	"fmt"
	"github.com/mkideal/cli"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
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

	// FIXME
	if argv.Repo != "" && argv.Organization == "" {
		return fmt.Errorf("Sorry, at the moment use without --org is not supported.")
	}

	if argv.Repo != "" && argv.Organization == "" {
		return fmt.Errorf("Sorry, at the moment use of --create-org is not supported.")
	}
	// FIXME

	if argv.CreateOrg && argv.Organization == "" {
		return fmt.Errorf("You can not use --create-org without --org \"someorg\"")
	}

	// Force check without --create-org and do check if --repo is specified
	if !argv.CreateOrg || argv.Repo != "" {
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
	}

	return nil
}

func main() {
	cli.Run(new(ArgT), func(ctx *cli.Context) error {
		argv := ctx.Argv().(*ArgT)

		tokenFileContent, err := ioutil.ReadFile(argv.TokenFile)
		if err != nil {
			return err
		}

		// extract token without white spaces
		r, err := regexp.Compile("[a-z0-9]+")
		if err != nil {
			return err
		}

		bytes := r.Find(tokenFileContent)
		if bytes == nil {
			return fmt.Errorf("File %s does NOT contain token.", argv.TokenFile)
		}
		token := string(bytes)

		/*
			if argv.CreateOrg {
				ctx.String("createOrg(serverURI, token, orgName)")
			}
		*/

		if !argv.CreateOrg && argv.Repo != "" {
			repoPathClean, err := filepath.Abs(argv.Repo)
			if err != nil {
				return err
			}

			projectName := argv.Project
			if projectName == "" {
				projectName = filepath.Base(repoPathClean)
			}

			cloneURL, err := createRepo(argv.Server, token, argv.Organization, projectName)
			if err != nil {
				return err
			}

			fmt.Printf("cd %s\ngit remote add gogs %s\ngit push --mirror gogs\n", repoPathClean, cloneURL)
		}

		return nil
	}, "Create organization and mirror git repo to Gogs server")
}
