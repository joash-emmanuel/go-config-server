package pull

import (
	"fmt"

	"go-git/clone"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

func Pull_configs() {

	// Pull the latest changes from the origin remote and merge into the current branch

	// We instantiate a new repository targeting the given path (the .git folder)
	r, err := git.PlainOpen(clone.Path)
	if err != nil {
		panic(err)
	}

	// Get the working directory for the repository
	w, err := r.Worktree()
	if err != nil {
		panic(err)
	}

	err = w.Pull(&git.PullOptions{
		RemoteURL: clone.Url,
		Auth: &http.BasicAuth{
			Username: clone.Username, // anything except an empty string
			Password: clone.Password, //set the access token or password
		},
		// Remote branch to clone.
		ReferenceName: plumbing.ReferenceName("refs/heads/" + clone.Branch), //Remote branch to clone. If empty, uses HEAD.
		// Fetch only ReferenceName if true. A single branch
		SingleBranch: true,
	})

	ref, err := r.Head()
	if err != nil {
		panic(err)
	}
	commit, err := r.CommitObject(ref.Hash())
	if err != nil {
		panic(err)
	}
	fmt.Println(commit)

}
