package main

//var branch string
import (
	"fmt"
	"log"

	rep "github.com/go-erl/gitclean/repository"
	git "gopkg.in/src-d/go-git.v4"
)

func main() {

	dir := rep.CheckRepository() //path

	repo, err := git.PlainOpen(dir) // *Repository
	if err != nil {
		log.Fatal(err)
	}

	config := rep.GetLocalBranches(*repo) //storer.ReferenceIter

	fmt.Println("Br:", *config)
	//branches := rep.GetLocalBranches(dir)
	//STILL TO IMPLEMENT DELETE FUNCTION

	//rep.DeleteLocalBranches(*repo, branches)

}
