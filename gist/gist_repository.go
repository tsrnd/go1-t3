package gist

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

type GistRepository interface {
	doGistsRequest(user string) (io.Reader, error)
}

type gistRepository struct{}

func (g *gistRepository) doGistsRequest(user string) (io.Reader, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/gists", user))

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, resp.Body); err != nil {
		return nil, err
	}

	return &buf, nil
}

func NewGistRepository() GistRepository {
	return &gistRepository{}
}
