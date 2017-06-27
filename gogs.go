package main

import (
	"fmt"
	"github.com/gogits/go-gogs-client"
)

func createOrg(serverURI, token, orgName string) error {
	gogsClient := gogs.NewClient(serverURI, token)
	_, err := gogsClient.GetOrg(orgName)
	if err != nil && err.Error() != "404 Not Found" {
		return err
	}

	// _, err := gogsClient.AdminCreateOrg()
	return fmt.Errorf("FIXME")
}

func createRepo(serverURI, token, orgName, projectName string) (string, error) {
	gogsClient := gogs.NewClient(serverURI, token)

	// TODO: find user name!!!
	var userName string

	var err error
	var repo *gogs.Repository

	if orgName == "" {
		repo, err = gogsClient.GetRepo(userName, projectName)
		if err != nil && err.Error() != "404 Not Found" {
			return "", err
		}

		// TODO
	} else {
		repo, err = gogsClient.GetRepo(orgName, projectName)
		if err != nil && err.Error() != "404 Not Found" {
			return "", err
		}

		if repo != nil && repo.ID != 0 {
			return "", fmt.Errorf("Repository %s under organization %s already exists.", projectName, orgName)
		}

		createOptions := gogs.CreateRepoOption{Name: projectName, Description: projectName, Private: true, AutoInit: false}
		repo, err = gogsClient.CreateOrgRepo(orgName, createOptions)
		if err != nil {
			return "", err
		}

		return repo.CloneURL, nil
	}

	return "FIXME", nil
}
