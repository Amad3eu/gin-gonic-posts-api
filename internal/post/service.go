package post

import (
	"errors"
	"unicode/utf8"

	"github.com/Amad3eu/gin-gonic-posts-api/internal"
	"github.com/google/uuid"
)

var ErrPostBodyEmpty = errors.New("post body is empty :(")
var ErrPostBodyExceedsLimit = errors.New("post body exceeds 140 characters :(")
var ErrPostNotFound = errors.New("post not found :(")

type Service struct {
	Repository Repository
}

func (p Service) Create(post internal.Post) error {
	if post.Body == "" {
		return ErrPostBodyEmpty
	}

	if utf8.RuneCountInString(post.Body) > 140 {
		return ErrPostBodyExceedsLimit
	}

	return p.Repository.Insert(post)
}

func (s Service) Delete(id uuid.UUID) error {
	return s.Repository.Delete(id)
}

func (s Service) FindOneByID(ID uuid.UUID) (internal.Post, error) {
	return s.Repository.FindOneByID(ID)
}

func (s Service) FindAll() ([]internal.Post, error) {
	return s.Repository.FindAll()
}
