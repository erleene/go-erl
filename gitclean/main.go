package main

//var branch string
import (
	rep "github.com/go-teach/git-clean/repository"
)

func main() {
	//branch := flag.String("branch", "", "the type of branch to list")

	//flag.Parse()

	//if *branch == "" {
	//	fmt.Println("Please provide the type of branch to list")
	//}

	dir := rep.CheckRepository()
	branches := rep.GetLocalBranches(dir)
	rep.DeleteLocalBranches(branches)

}
