package clone

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

var Url = os.Getenv("GIT_URL")
var Username = os.Getenv("GIT_USERNAME")
var Password = os.Getenv("GIT_PASSWORD")
var Branch = os.Getenv("GIT_BRANCH")
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
		log.Printf("Cloned %v, branch %v, with Commit ID %v", Url, branch.Name(), branch.Hash().String())
		return nil
	})

	// // ... retrieving the branch being pointed by HEAD
	// ref, err := r.Head()
	// if err != nil {
	// 	panic(err)
	// }
	// // ... retrieving the commit object
	// commit, err := r.CommitObject(ref.Hash())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(commit)

}
