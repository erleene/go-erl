package main

//var branch string
import (
	"log"

	rep "github.com/go-erl/gitclean/repository"
	git "gopkg.in/src-d/go-git.v4"
)

func main() {
	//branch := flag.String("branch", "", "the type of branch to list")

	//flag.Parse()

	//if *branch == "" {
	//	fmt.Println("Please provide the type of branch to list")
	//}

	dir := rep.CheckRepository()

	repo, err := git.PlainOpen(dir)
	if err != nil {
		log.Fatal(err)
	}

	branches := rep.GetLocalBranches(*repo)

	//STILL TO IMPLEMENT DELETE FUNCTION

	rep.DeleteLocalBranches(*repo, branches)

}
