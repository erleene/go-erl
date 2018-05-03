package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/DATA-DOG/godog"
	homedir "github.com/mitchellh/go-homedir"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/config"
)

var gitWorkingDir string

func goPath() (string, error) {
	s := os.Getenv("GOPATH")
	if s != "" {
		return s, nil
	}
	return homedir.Expand("~/go")
}

func iAmInTheGoerlgitcleanProjectRepo() error {
	goPath, err := goPath()
	if err != nil {
		return err
	}

	gitWorkingDir = fmt.Sprintf("%s/%s", goPath, "src/github.com/erleene/go-erl")
	return nil
}

func iHaveMultipleLocalBranches() error {
	fmt.Println(gitWorkingDir)
	repo, err := git.PlainOpen(gitWorkingDir)
	if err != nil {
		return err
	}

	b1 := config.Branch{
		Name:   "b1",
		Remote: "some_remote_1",
		Merge:  "refs/heads/b1",
	}

	b2 := config.Branch{
		Name:   "b2",
		Remote: "some_remote_2",
		Merge:  "refs/heads/b2",
	}
	//IMPROVEMENT: if branch dont exist, create it.

	repo.CreateBranch(&b1)
	repo.CreateBranch(&b2)

	return nil
}

func iRunGitcleanInThatRepo() error {
	err := RunGitClean(gitWorkingDir)
	if err != nil {
		return err
	}

	return nil
}

func iShouldOnlyBeLeftWithTheMasterBranch() error {
	//list the branches
	//capture output as string from os/exec
	err := os.Chdir(gitWorkingDir)
	if err != nil {
		return err
	}

	bytes, err := exec.Command("git branch --list").Output()

	//convert bytes to string, then check for master string
	hasMaster := strings.Contains(string(bytes), "master")

	if !hasMaster {
		return fmt.Errorf("Expected Master Branch")
	}

	//check b1 and b2 no longer exist
	hasB1 := strings.Contains(string(bytes), "b1")
	hasB2 := strings.Contains(string(bytes), "b2")

	if hasB1 || hasB2 {
		return fmt.Errorf("Did not expect branches: b1 and b2")
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I am in the go-erl\/gitclean project repo$`, iAmInTheGoerlgitcleanProjectRepo)
	s.Step(`^I have multiple local branches$`, iHaveMultipleLocalBranches)
	s.Step(`^I run gitclean in that repo$`, iRunGitcleanInThatRepo)
	s.Step(`^I should only be left with the master branch$`, iShouldOnlyBeLeftWithTheMasterBranch)
}
