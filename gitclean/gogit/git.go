package gogit

import (
	"fmt"
	"os/exec"
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
	name := getNameFromPath(path)

	r := &Repository{
		name: "", //to take from the path
		path: path,
	}

	r.branches = make(map[string]Branch) //branch
	branchNames := getRepositoryBranchNames(path)

  for _, name := branchNames {
    r.branches[name] = Branch{name: name,}
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

}

func (r *Repository) ListRemoteBranches() error {

}
