package repository

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	//cli
)

/**

this package will contain the building blocks of the git-clean program
the idea is to have the program check if the current directory we're on is a Repository
if the current directory is a repository, we then want to look at the Branches
we want to delete the local branches (and remote branches), except for MASTER
 **/

//struct object
type Repository struct {
	name     string
	path     string
	branches map[string]Branch
}

type Branch struct {
	name   string
	remote string
}

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

func ListLocalBranches(path string) ([]byte, error) {
	os.Chdir(path)
	stdoutStderr, err := exec.Command("git", "branch", "--list").Output()

	if err != nil {
		log.Fatal(err)
	}

	return stdoutStderr, err
}

func DeleteBranch(path string, branchName string) ([]byte, error) {
	os.Chdir(path)
	combOutput, err := exec.Command("git", "branch", "--delete", "--force", branchName).Output()
	if err != nil {
		return nil, err
	}
	return combOutput, err
}

//
// func ListRemoteBranches(path string) ([]byte, error) {
// 	os.Chdir(path)
// 	cmd := exec.Command("git", "branch", "--remote")
//
// 	stdoutStderr, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return stdoutStderr, err
// }
