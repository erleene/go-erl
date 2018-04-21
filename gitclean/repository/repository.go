package repository

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/plumbing/storer"
)

//this package will contain the building blocks of the git-clean program
// it will outline its specifications, etc.

//the idea is to have the program check if the current directory we're on is a Repository
//if the current directory is a repository, we then want to look at the Branches
//we want to delete the local branches, except for MASTER

//if dir contains .git folder then return true
func CheckRepository() string {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//if dir contains .git folder then return true
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() && info.Name() == ".git" {
			fmt.Println("This is a github repository")
		}
		return nil
	})
	return dir
}

func GetLocalBranches(dir string) interface{} {
	repo, err := git.PlainOpen(dir)
	if err != nil {
		return err
	}

	localBr, _ := repo.Branches()

	return localBr
}

func DeleteLocalBranches(branches interface{}) {

	//convert branches to storer.ReferenceIter object
	localBr := branches.(storer.ReferenceIter)

	localBr.ForEach(func(ref *plumbing.Reference) error {
		refName := ref.Name()
		//
		switch refName.IsBranch() {
		case refName != "refs/heads/master":
			fmt.Printf("Branch:%s \n", refName)

		default:
			fmt.Println("This is the master branch")
			fmt.Printf("Branch:%s \n", refName)
		}
		return nil
	})

}
