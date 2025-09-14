package clone

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

var Url = os.Getenv("GIT-URL")
var Username = os.Getenv("GIT-USERNAME")
var Password = os.Getenv("GIT-PASSWORD")
var Branch = os.Getenv("GIT-BRANCH")
var Path = "./repo"

func Clone_configs() {

	//plain clone specifies the directory to clone the repo into
	//git clone specifies the configs to pass to enable you clone the repo

	r, err := git.PlainClone(Path, &git.CloneOptions{
		URL: Url, // The (possibly remote) repository URL to clone from.
		Auth: &http.BasicAuth{
			Username: Username, // anything except an empty string
			Password: Password, //set the access token or password
		},
		ReferenceName: plumbing.ReferenceName("refs/heads/" + Branch), // Remote branch to clone.
		SingleBranch:  true,                                           // Fetch only ReferenceName if true. A single branch
	})

	if err != nil {
		fmt.Println("Error cloning repository:", err)
	} else {
		fmt.Println("Repository cloned successfully!")
	}

	branches, _ := r.Branches()
	branches.ForEach(func(branch *plumbing.Reference) error {
		fmt.Println(branch.Hash().String(), branch.Name())
		return nil
	})
}
