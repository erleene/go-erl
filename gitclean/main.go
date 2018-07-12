package main

//var branch string
import (
	"fmt"
	"log"
	"os"
	"strings"

	rep "github.com/erleene/go-erl/gitclean/repository"
)

//RunGitClean function to delete local branches
func RunGitClean(dir string) error {
	var outputs []string
	dir, err := rep.CheckRepository() //path
	if err != nil {
		return err
	}

	localBranches, err := rep.ListLocalBranches(dir)
	if err != nil {
		return err
	}

	outputs = append(outputs, strings.TrimSpace(string(localBranches)))
	//
	for _, v := range outputs {
		if v != "master" {
			fmt.Println("Deleting branch...", v)
			_, err := rep.DeleteBranch(dir, v)
			if err != nil {
				return err
			}
		}
	}
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
