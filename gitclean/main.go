package main

//var branch string
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
	//dir is a repo
	//lets now list all the branches in this repo
	localBrances, err = rep.getLocalBranches(dir)
	if err != nil {
		return err
	}

	// repo, err := git.PlainOpen(dir) // *Repository
	// if err != nil {
	// 	return err
	// }
	//
	// config, _ := repo.Config()
	// for key, value := range config.Branches {
	// 	//print local branches with remote branches
	// 	fmt.Printf("key: %s, value: %s\n", key, value.Name)
	//
	// 	//check for remote branches of the coinfig
	// }

	// if branch, exists := config.Branches["local1"]; exists {
	// 	fmt.Printf("branch %s exists", "local1")
	// 	if err := repo.DeleteBranch(branch.Name); err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	//
	// 	if err := repo.DeleteRemote(branch.Name); err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// }

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
