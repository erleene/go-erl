package repository

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
	config "gopkg.in/src-d/go-git.v4/config"
	"gopkg.in/src-d/go-git.v4/plumbing"
	storer "gopkg.in/src-d/go-git.v4/plumbing/storer"
	//cli
)

//this package will contain the building blocks of the git-clean program
// it will outline its specifications, etc.

//the idea is to have the program check if the current directory we're on is a Repository
//if the current directory is a repository, we then want to look at the Branches
//we want to delete the local branches, except for MASTER

func CheckRepository() string {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//if dir contains .git folder then return it
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		//this is a github repository
		if info.IsDir() && info.Name() == ".git" {
			return nil
		}
		return nil
	})
	return dir
}

// func GetLocalBranches(dir git.Repository) interface{} {
//
// 	localBr, _ := dir.Branches()
// 	return localBr
// }
func GetLocalBranches(dir git.Repository) *config.Config {

	//list all branches using command line git branch

	// localBr, _ := dir.Branches() //storer.ReferenceIter
	// return localBr

	//get Config
	conf, err := dir.Config()
	if err != nil {
		log.Fatal(err)
	}

	return conf
}

//TO IMPLEMENT
func DeleteLocalBranches(dir git.Repository, br interface{}) {

	localBr := br.(storer.ReferenceIter)

	//get the branches of dir is a Repository that belongs to a struct with field

	// branch, err := dir.Config()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// bra := branch.Branches
	//
	// fmt.Println(bra)

	localBr.ForEach(func(ref *plumbing.Reference) error {
		refName := ref.Name()
		//
		switch refName.IsBranch() {
		case refName == "refs/heads/master":
			fmt.Printf("MASTER BRANCH:%s \n", refName)
		default:
			fmt.Printf("DELETE THIS branch:%s \n", refName)

			//to delete the branch we need a string of the branch name DeleteBranch(refName)
			//branchToDelete := refName.String()
			//dir.DeleteBranch(branchToDelete)
		}
		return nil
	})

}
