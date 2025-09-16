package pull

import (
	"fmt"
	"time"

	"go-git/clone"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

func Fetch_configs() {

	repo, err := git.PlainOpen(clone.Path)
	if err != nil {
		panic(err)
	}

	list, err := repo.Remotes()
	if err != nil {
		panic(err)
	}

	for _, r := range list {
		// fmt.Println(r.Config().Name, r.Config().URLs)
		// // Fetch
		err = r.Fetch(&git.FetchOptions{
			RemoteURL: clone.Url,
			Auth: &http.BasicAuth{
				Username: clone.Username, // anything but empty string
				Password: clone.Password,
			},
			// If you want only one branch:
			RefSpecs: []config.RefSpec{
				config.RefSpec("+refs/heads/" + clone.Branch + ":/refs/remotes/origin/" + clone.Branch),
				//Refspecs are always in the form [+]<source>:<destination>/local, : (Colon): This separates the source ref from the destination ref.
				// eg"+refs/heads/*:refs/remotes/origin/*" --* stands for the branch
			},
		})
		if err != nil && err != git.NoErrAlreadyUpToDate {
			fmt.Printf("failed to fetch %v in %v repo: %v", r.Config().Name, clone.Path, err)
			panic(err)

		}

		remoteRefs, err := r.List(&git.ListOptions{
			Auth: &http.BasicAuth{
				Username: clone.Username, // anything but empty string
				Password: clone.Password,
			},
		})
		if err != nil {
			panic(err)
		}

		for _, r := range remoteRefs {
			fmt.Println(r.Hash())
		}

	}

}

func Pull_configs() {

	forever := make(chan bool)

	go func() {
		for {

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

			// Pull the latest changes from the origin remote and merge into the current branch
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
			//The data from the step above gets sent to r

			if err != nil {
				if err.Error() == "already up-to-date" {
					fmt.Println("Fetching up to date")
				} else {
					panic(err)
				}
			}
			// Print the latest commit that was just pulled

			ref, err := r.Head()
			if err != nil {
				panic(err)
			}

			commit, err := r.CommitObject(ref.Hash())
			if err != nil {
				panic(err)
			}
			fmt.Println(commit)
			fmt.Println(commit.Hash)

			time.Sleep(2 * time.Second)
		}

	}()
	<-forever

}

// //Get the current commit hash in the folder
// local_commit_hash, err := r.Head()
// if err != nil {
// 	panic(err)
// }
// fmt.Printf("current hash in the folder - %v\n", local_commit_hash.Hash())
