package main

import (
	"log"
	"os"

	rep "github.com/erleene/go-erl/gitclean/repository"
)

//RunGitClean function to delete local branches
func RunGitClean(dir string) error {
	//var outputs []string
	dir, err := rep.CheckRepository() //path
	if err != nil {
		return err
	}

	_, err = rep.ListLocalBranches(dir)
	if err != nil {
		return err
	}

	//}

	//out := string(localBranches)
	//newOut := strings.TrimPrefix(localBranches, "*")
	//outputs = append(outputs, strings.TrimSpace(newOut))
	//store the results in one string so do a concatenation
	//toDelete := strings.TrimSuffix(localBranches, "* master")
	// for _, v := range outputs {
	// 	if v == "master" {
	// 		fmt.Printf("%v", v)
	// 	}
	// }
	//fmt.Println(localBranches)
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
