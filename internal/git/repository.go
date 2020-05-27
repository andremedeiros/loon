package git

import (
	"fmt"
	"os/user"
	"strings"

	gogit "github.com/go-git/go-git/v5"
)

type Repository interface {
	Clone(path string) error
	CheckoutURL(useSecure bool) string
	Host() string
	Name() string
	Owner() string
	URL() string
}

type GitHubRepository struct {
	name  string
	owner string
}

func NewRepository(nameWithOwner string) Repository {
	name := nameWithOwner
	owner := ""

	if strings.Contains(nameWithOwner, "/") {
		split := strings.SplitN(nameWithOwner, "/", 2)
		owner = split[0]
		name = split[1]
	} else {
		user, _ := user.Current()
		owner = user.Username
	}

	return &GitHubRepository{
		name:  name,
		owner: owner,
	}
}

func (ghr *GitHubRepository) Clone(path string) error {
	_, err := gogit.PlainClone(path, false, &gogit.CloneOptions{
		URL: ghr.CheckoutURL(true),
	})

	return err
}

func (ghr *GitHubRepository) Host() string {
	return "github.com"
}

func (ghr *GitHubRepository) Name() string {
	return ghr.name
}

func (ghr *GitHubRepository) Owner() string {
	return ghr.owner
}

func (ghr *GitHubRepository) URL() string {
	return fmt.Sprintf("https://github.com/%s/%s", ghr.owner, ghr.name)
}

func (ghr *GitHubRepository) CheckoutURL(useSecure bool) string {
	if useSecure {
		return fmt.Sprintf("git@github.com:%s/%s.git", ghr.owner, ghr.name)
	}

	return fmt.Sprintf("https://github.com/%s/%s.git", ghr.owner, ghr.name)
}
