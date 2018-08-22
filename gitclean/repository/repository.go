package repository

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	//cli
)

/**

this package will contain the building blocks of the git-clean program
the idea is to have the program check if the current directory we're on is a Repository
if the current directory is a repository, we then want to look at the Branches
we want to delete the local branches (and remote branches), except for MASTER
 **/

const GitDirName = ".git"

func CheckRepository() (string, error) {

	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	//if dir contains .git folder then return it
	//REIMPLEMENT THIS BY TAKING ON THE DFS algorith
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		//this is a github repository
		if info.IsDir() && info.Name() == GitDirName {
			return nil
		}
		return nil
	})
	return dir, err
}

func ListLocalBranches(path string) ([]string, error) {

	os.Chdir(path)
	stdoutStderr, err := exec.Command("git", "branch", "--list").CombinedOutput() //[]byte

	if err != nil {
		log.Fatal(err)
	}

	output := strings.Fields(string(stdoutStderr))

	var branches []string

	for i := 0; i < len(output); i++ {
		branches = append(branches, output[i])
		fmt.Println(branches[i])
	}

	return branches, err
}

func DeleteLocalBranches(path string, branches []string) error {

	os.Chdir(path)

	//delete all except MASTER
	for i := 0; i < len(branches); i++ {
		if branches[i] != "master" {
			_, err := exec.Command("git", "branch", "-D", branches[i]).Output()

			if err != nil {
				return err
			}
		}
	}
	return nil
}
