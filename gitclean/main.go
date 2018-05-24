package main

//var branch string
import (
	"fmt"
	"log"
	"os"

	rep "github.com/erleene/go-erl/gitclean/repository"
	git "gopkg.in/src-d/go-git.v4"
)

//RunGitClean function to delete local branches
func RunGitClean(dir string) error {
	dir, err := rep.CheckRepository() //path
	if err != nil {
		return err
	}

	repo, err := git.PlainOpen(dir) // *Repository
	if err != nil {
		return err
	}

	config, _ := repo.Config()
	for key, value := range config.Branches {
		fmt.Printf("key: %s, value: %s\n", key, value.Name)
	}

	if branch, exists := config.Branches["local1"]; exists {
		fmt.Printf("branch %s exists", "local1")
		if err := repo.DeleteBranch(branch.Name); err != nil {
			fmt.Println(err.Error())
		}

		if err := repo.DeleteRemote(branch.Name); err != nil {
			fmt.Println(err.Error())
		}
	}

	// config := rep.GetConfiguration(repo) //config.Config
	//
	// rep.DeleteLocalBranches(repo, config)
	//
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
