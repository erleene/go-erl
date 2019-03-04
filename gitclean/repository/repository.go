package repository

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
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
	log.Info("Checking directory for git branches...")
	dir, geterr := os.Getwd()
	if geterr != nil {
		return "", geterr
	}

	//if dir contains .git folder then return it
	//REIMPLEMENT THIS BY TAKING ON THE DFS algorith
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

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

func ListBranches(path string) (*Repository, error) {
	log.Info("List Branches to delete...")
	os.Chdir(path)
	localStderr, err := exec.Command("git", "branch", "--list").CombinedOutput() //[]byte
	remoteStderr, err := exec.Command("git", "branch", "-r").CombinedOutput()

	if err != nil {
		log.Fatalf("Unabled to fetch branches %v", err)
	}

	localOutput := strings.Fields(string(localStderr))
	remoteOutput := strings.Fields(string(remoteStderr))

	var localBranches []string
	var remoteBranches []string

	for i := 0; i < len(localOutput); i++ {
		localBranches = append(localBranches, localOutput[i])
	}
	for j := 0; j < len(remoteOutput); j++ {
		remoteBranches = append(remoteBranches, remoteOutput[j])
	}
	log.Info("Completed fetching branches to delete...")
	return &Repository{
		LocalBranches:  localBranches,
		RemoteBranches: remoteBranches,
	}, err
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

func DeleteLocalBranches(path string, branches *Repository) error {
	log.Info("Preparing to delete branches")
	os.Chdir(path)

	//delete all except MASTER
	for i := 0; i < len(branches.LocalBranches); i++ {
		log.Info("preparing local...")
		if branches.LocalBranches[i] != "master" {
			_, err := exec.Command("git", "branch", "-D", branches.LocalBranches[i]).Output()
			if err != nil {
				log.Fatalf("unable to delete local branches: %v", err)
				os.Exit(1)
			}
			log.Info("completed preparing local branches")
		}
	}

	for j := 0; j < len(branches.RemoteBranches); j++ {
		log.Info("preparing remote...")
		if branches.RemoteBranches[j] != "master" {
			_, err := exec.Command("git", "branch", "-D", branches.RemoteBranches[j]).Output()
			if err != nil {
				log.Fatal(err)
			}
		}
		log.Info("Compelted preparing for remote branches")
	}
	return nil
}
