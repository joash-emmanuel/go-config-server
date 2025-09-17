package fetch

import (
	"fmt"

	"go-git/clone"

	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/config"
	"github.com/go-git/go-git/v6/plumbing"
	"github.com/go-git/go-git/v6/plumbing/transport/http"
)

func Fetch_configs() plumbing.Hash {

	for {
		// We instantiate a new repository targeting the given path (the .git folder)
		repo, err := git.PlainOpen(clone.Path)
		if err != nil {
			panic(err)
		}

		// List remotes from a repository
		list_remotes, err := repo.Remotes()
		if err != nil {
			panic(err)
		}

		//iterate through the remotes and do a fetch for the specified branch
		for _, r := range list_remotes {
			// fmt.Println(r.Config().Name, r.Config().URLs)
			err = r.Fetch(&git.FetchOptions{
				RemoteURL: clone.Url,
				Auth: &http.BasicAuth{
					Username: clone.Username, // anything but empty string
					Password: clone.Password,
				},
				RefSpecs: []config.RefSpec{
					config.RefSpec("+refs/heads/" + clone.Branch + ":/refs/remotes/origin/" + clone.Branch),
					//Refspecs are always in the form [+]<source>:<destination>/local, : Colon separates the source ref from the destination ref.
					// eg"+refs/heads/*:refs/remotes/origin/*" --* stands for the branch
				},
			})
			if err != nil {
				if err.Error() == "already up-to-date" {
					fmt.Println("Fetching up to date")
				} else {
					panic(err)
				}
			}

			//ListOptions describes how a remote list should be performed
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
				if r.Type() == plumbing.HashReference {
					fmt.Println(r.Hash())
					// return r.Hash()
					Remote_Hash := r.Hash()
					return Remote_Hash
				}

			}

		}
	}

}
