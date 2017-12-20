package gist

import "encoding/json"

type GistUsecase interface {
	ListGists(user string) ([]string, error)
}

type gistUsecase struct {
	Gister GistRepository
}

// ListGists return url list
func (c *gistUsecase) ListGists(user string) ([]string, error) {
	r, err := c.Gister.doGistsRequest(user)
	if err != nil {
		return nil, err
	}

	var gists []Gist
	if err := json.NewDecoder(r).Decode(&gists); err != nil {
		return nil, err
	}

	urls := make([]string, 0, len(gists))
	for _, u := range gists {
		urls = append(urls, u.Rawurl)
	}

	return urls, nil
}

func NewGistUsecase(g GistRepository) GistUsecase {
	return &gistUsecase{g}
}
