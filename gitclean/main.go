package main

//var branch string
import (
	"log"
	"os"

	rep "github.com/erleene/go-erl/gitclean/repository"
	git "gopkg.in/src-d/go-git.v4"
)

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

//function to delete local branches
func RunGitClean(dir string) error {
	dir, err := rep.CheckRepository() //path
	if err != nil {
		return err
	}

	repo, err := git.PlainOpen(dir) // *Repository
	if err != nil {
		return err
	}

	config := rep.GetConfiguration(*repo) //config.Config

	rep.DeleteLocalBranches(*repo, *config)

	return nil
}
