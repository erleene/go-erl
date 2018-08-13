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

func ListLocalBranches(path string) (string, error) {
	//var out []string

	os.Chdir(path)
	stdoutStderr, err := exec.Command("git", "branch", "--list").CombinedOutput() //[]byte

	if err != nil {
		log.Fatal(err)
	}

	output := strings.Fields(string(stdoutStderr))

	//fmt.Println(strings.Join(output, ","))

	te := strings.Join(output, ",")

	//remove * and MASTER from te
	noM := strings.Trim(te, "master")
	fmt.Println(strings.TrimRight(noM, "*"))

	//fmt.Println(noM)
	//newOut := strings.Trim(output, "* & master")
	return strings.Join(output, ","), err
}

func DeleteBranch(path string, branchName string) error {

	os.Chdir(path)
	_, err := exec.Command("git", "branch", "-D", branchName).Output()

	if err != nil {
		return err
	}

	return err
}
