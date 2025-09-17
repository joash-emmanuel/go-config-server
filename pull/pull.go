package pull

import (
	"fmt"
	"log"
	"time"

	"go-git/clone"
	"go-git/fetch"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

func Pull_configs() {

	forever := make(chan bool)

	go func() {
		for {

			Remote_Hash := fetch.Fetch_configs()
			log.Printf("current-hash in the remote - %v\n", Remote_Hash)

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

			//Get the current commit hash in the folder
			local_commit_hash, err := r.Head()
			if err != nil {
				panic(err)
			}
			fmt.Printf("current hash in the folder - %v\n", local_commit_hash.Hash())

			//compare the remote hash and local hash. If they're different, a pull is done
			if Remote_Hash != local_commit_hash.Hash() {
				log.Println("New commit found")
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
				// The data from the step above gets sent to r

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
				fmt.Println(commit.Hash)
			} else {
				log.Println("No new commit")
			}

			time.Sleep(2 * time.Second)
		}

	}()
	<-forever

}
