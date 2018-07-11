package main

//var branch string
import (
	"fmt"
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

	//dir is a git repository
	localBranches, err := rep.ListLocalBranches(dir)
	if err != nil {
		return err
	}

	fmt.Printf("%s\n", localBranches)

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
