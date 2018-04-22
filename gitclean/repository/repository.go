package repository

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	git "gopkg.in/src-d/go-git.v4"
	config "gopkg.in/src-d/go-git.v4/config"
	//cli
)

//this package will contain the building blocks of the git-clean program
// it will outline its specifications, etc.

//the idea is to have the program check if the current directory we're on is a Repository
//if the current directory is a repository, we then want to look at the Branches
//we want to delete the local branches, except for MASTER

func CheckRepository() (string, error) {

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	//if dir contains .git folder then return it
	err1 := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {

		//this is a github repository
		if info.IsDir() && info.Name() == ".git" {
			return nil
		}
		return nil
	})
	return dir, err1
}

func GetConfiguration(dir git.Repository) *config.Config {

	//get Config
	conf, err := dir.Config()
	if err != nil {
		log.Fatal(err)
	}

	return conf
}

//TO IMPLEMENT
func DeleteLocalBranches(dir git.Repository, conf config.Config) {

	//list branchs
	br := conf.Branches //Config struct with Branches (map[string]*Branch)

	for brName, v := range br {
		fmt.Printf("\n")
		fmt.Println("==========")
		fmt.Println("Branch Name:", brName)
		fmt.Printf("\n")
		fmt.Println("Branch:", *v)
		fmt.Printf("\n")
		refSpec := v.Merge
		fmt.Println("Branch refspec value: ", refSpec)
		fmt.Println("==========")
		fmt.Printf("\n")

		//TO DELETE A BRANCH FROM THE REPOSITORY: dir.DeleteBranch(v.Name)

		// switch {
		//
		// case BranchName == "master":
		// 	fmt.Println("\n")
		// 	fmt.Println("This is the master branch")
		//
		// case BranchName == "local2":
		// 	//delete
		// 	fmt.Println("DELETING THIS BRANCH...")
		// 	err := dir.DeleteBranch(v.Name)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// 	er := dir.DeleteRemote(v.Name)
		// 	if er != nil {
		// 		log.Fatal(er)
		// 	}
		// default:
		// 	//DELETE
		// 	fmt.Println("====")
		// }
	}

	//localBr := br.(storer.ReferenceIter)

	//get the branches of dir is a Repository that belongs to a struct with field

	// branch, err := dir.Config()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// bra := branch.Branches
	//
	// fmt.Println(bra)

	// 	localBr.ForEach(func(ref *plumbing.Reference) error {
	// 	refName := ref.Name()
	// 	//
	// 	switch refName.IsBranch() {
	// 	case refName == "refs/heads/master":
	// 		fmt.Printf("MASTER BRANCH:%s \n", refName)
	// 	default:
	// 		fmt.Printf("DELETE THIS branch:%s \n", refName)
	//
	// 		//to delete the branch we need a string of the branch name DeleteBranch(refName)
	// 		//branchToDelete := refName.String()
	// 		//dir.DeleteBranch(branchToDelete)
	// 	}
	// 	return nil
	// })

}
