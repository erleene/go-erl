package gogit

import (
	"fmt"
	"os/exec"
	"strings"
)

//struct object
type Repository struct {
	name     string
	path     string
	branches map[string]Branch
}

type Branch struct {
	name   string
	remote string
}



func LoadRepository(path string) {
	repoName := getNameFromPath(path)

	r := &Repository{
		name: repoName, //to take from the path
		path: path,
	}

	r.branches = make(map[string]Branch) //branch
	branchNames := getRepositoryBranchNames(path)

  for _, name := branchNames {
    r.branches[name] = Branch{name: name,}
  }
}

func getNameFromPath(path string) string {

	s := strings.Split(path, "/")

	//get the last element of s
	name := s[-1]

	return name

}

func getRepositoryBranchNames(path string) map[string]Branch {
	//from the path, run the command to collect the branch branchNames

	//make sure you are in the PATH
	os.Chdir(path)
	out, err := exec.Command("git ", "branch", "--list").Output()

	if err != nil {
		return err
	}
	//create a new map

	branches := make(map[string]Branch)

	for i := 0; i < len(out); i++ {
		branches[] = out[i]
	}

}

func (r *Repository) CreateBranch(branchName, branchRemote string) error {

	b := Branch{name: branchName, remote: branchRemote}

	out, err := exec.Command("git ", "checkout", "-b", branchName).Output()

	if err != nil {
		return err
	}

	fmt.Println(string(out))

	r.branches[branchName] = b

	return nil
}

func (r *Repository) DeleteBranch(branchName string) error {

	if _, ok := r.branches[branchName]; ok {
		out, err := exec.Command("git ", "branch", "-D", branchName).Output()

		if err != nil {
			return err
		}

		fmt.Println(string(out))

		delete(r.branches, branchName)

		return nil
	}
	return nil

}

func (r *Repository) ListLocalBranches() error {

	out, err := exec.Command("git ", "branch", "--list").Output()

	if err != nil {
		return err
	}
	return out

}

func (r *Repository) ListRemoteBranches() error {
	out, err := exec.Command("git ", "branch", "--remote").Output()

	if err != nil {
		return err
	}
	return out

}
