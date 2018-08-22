package main

import (
	"log"
	"os"

	rep "github.com/erleene/go-erl/gitclean/repository"
)

//RunGitClean function to delete local branches
func RunGitClean(dir string) error {
	dir, err := rep.CheckRepository() //path
	if err != nil {
		return err
	}

	branches, err := rep.ListLocalBranches(dir)
	if err != nil {
		return err
	}

	rep.DeleteLocalBranches(dir, branches)

	return nil
}

func main() {
	workingDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	err = RunGitClean(workingDir)
	if err != nil {
		log.Fatal(err)
	}

}
