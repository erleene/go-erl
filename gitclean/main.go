package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	rep "github.com/erleene/go-erl/gitclean/repository"
)

//RunGitClean function to delete local branches
func RunGitClean(dir string) error {
	dir, err := rep.CheckRepository() //path
	if err != nil {
		log.Fatalf("Unabled to check repository: %v", err)
	}

	updateerr := rep.UpdateBranches(dir)
	if updateerr != nil {
		return updateerr
	}

	branches, err := rep.ListBranches(dir)
	if err != nil {
		return err
	}

	delErr := rep.DeleteLocalBranches(dir, branches)
	if delErr != nil {
		return err
	}

	return nil
}

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatalf("Unabled to get working directory %v", err)
	}
	log.Info("Fetched working directory to check...")

	Runerr := RunGitClean(workingDir)
	if Runerr != nil {
		log.Fatalf("Unabled to to delete branches: %v", err)
	}

}
