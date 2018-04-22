package main

//var branch string
import (
	"log"

	rep "github.com/go-erl/gitclean/repository"
	git "gopkg.in/src-d/go-git.v4"
)

func main() {

	dir, err := rep.CheckRepository() //path
	if err != nil {
		log.Fatal(err)
	}

	repo, err := git.PlainOpen(dir) // *Repository
	if err != nil {
		log.Fatal(err)
	}

	config := rep.GetConfiguration(*repo) //config.Config

	//br := config.Branches

	//fmt.Println("Br:", &br) //map[string]Branch

	//branches := rep.GetLocalBranches(dir)
	//STILL TO IMPLEMENT DELETE FUNCTION

	rep.DeleteLocalBranches(*repo, *config)

}
