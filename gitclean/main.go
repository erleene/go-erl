package main

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
	// for _, v := range strings.TrimSpace(string(localBranches), "* master") {
	// 	//fmt.Println(k)
	// 	fmt.Println("%v", v)
	// }

	outputs = append(outputs, strings.TrimSpace(string(localBranches)))
	//store the results in one string so do a concatenation
	//toDelete := strings.TrimSuffix(localBranches, "* master")
	for _, v := range outputs {
		//if master then delete
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
