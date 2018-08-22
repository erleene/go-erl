package repository

import (
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

type Repository struct {
	LocalBranches  []string
	RemoteBranches []string
}

func ListBranches(path string) (r *Repository, error) {

	os.Chdir(path)
	localStderr, err := exec.Command("git", "branch", "--list").CombinedOutput() //[]byte
	remoteStderr, err := exec.Command("git", "branch", "-r").CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	var r Repository

	localOutput := strings.Fields(string(localStderr))
	remoteOutput := strings.Fields(string(remoteStderr))

	var localBranches []string
	var remoteBranches []string

	for i := 0; i < len(localOutput); i++ {
		localBranches := append(localBranches, localOutput[i])
	}

	for j := 0; j < len(remoteOutput); j++ {
		remoteBranches := append(remoteBranches, remoteOutput[j])

	}
	//now add the arrays to the struct
	r = Repository{
		LocalBranches:  localBranches,
		remoteBranches: remoteBranches,
	}
	// var branches []string
	//
	// for i := 0; i < len(output); i++ {
	// 	branches = append(branches, output[i])
	// }
	//
	// return branches, err
	return r
}

// func ListBranches(path string) ([]string, error) {
//
// 	os.Chdir(path)
// 	stdoutStderr, err := exec.Command("git", "branch", "--list").CombinedOutput() //[]byte
//
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	output := strings.Fields(string(stdoutStderr))
//
// 	var branches []string
//
// 	for i := 0; i < len(output); i++ {
// 		branches = append(branches, output[i])
// 	}
//
// 	return branches, err
// }

func DeleteBranches(path string, branches []string) error {

	os.Chdir(path)

	//delete all except MASTER
	for i := 0; i < len(branches); i++ {
		if branches[i] != "master" {
			_, err := exec.Command("git", "branch", "-D", branches[i]).Output()
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	return nil
}
