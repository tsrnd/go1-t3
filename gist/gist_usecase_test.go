package gist

import (
	"errors"
	"io"
	"strings"
	"testing"
)

var requestedUser string

type dummyGistRepository struct{}

func (d *dummyGistRepository) doGistsRequest(user string) (io.Reader, error) {
	requestedUser = user
	return strings.NewReader(`
		[
			{"html_url": "https://gist.github.com/1a861e999f3c73f83146"},
			{"html_url": "https://gist.github.com/653d65c17d8fe1d372af"}				
		]
		`), nil
}

func TestListGists(t *testing.T) {
	// save original value, and recover it at the end.
	saved := requestedUser
	defer func() { requestedUser = saved }()

	// use dummy object
	uc := &gistUsecase{
		Gister: &dummyGistRepository{},
	}

	const user = "test"
	urls, err := uc.ListGists(user)
	if err != nil {
		t.Errorf("list gets caused error: %s", err)
	}

	if requestedUser == "" {
		t.Fatalf("doGistsRequest not called")
	}

	if requestedUser != user {
		t.Errorf("wrong user (%s) requested, wants %s", requestedUser, user)
	}

	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}

type dummyGistRepositoryHTTPError struct{}

func (d *dummyGistRepositoryHTTPError) doGistsRequest(user string) (io.Reader, error) {
	return nil, errors.New("not found")
}

func TestListGistsHTTPError(t *testing.T) {
	// use dummy object
	uc := &gistUsecase{
		Gister: &dummyGistRepositoryHTTPError{},
	}

	const user = "test"
	_, err := uc.ListGists(user)
	if err == nil {
		t.Errorf("expected error. but succeed")
	}
}

type dummyGistRepositoryJSONDecodeError struct{}

func (d *dummyGistRepositoryJSONDecodeError) doGistsRequest(user string) (io.Reader, error) {
	requestedUser = user
	return strings.NewReader(`
		[
			"html_url": "https://gist.github.com/1a861e999f3c73f83146"},
			{"html_url": "https://gist.github.com/653d65c17d8fe1d372af"}				
		]
		`), nil
}

func TestListGistsJsonDecodeError(t *testing.T) {
	// use dummy object
	uc := &gistUsecase{
		Gister: &dummyGistRepositoryJSONDecodeError{},
	}

	_, err := uc.ListGists("test")
	if err == nil {
		t.Errorf("expected error. but succeed")
	}
}
