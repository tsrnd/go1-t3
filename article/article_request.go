package article

import "mime/multipart"

// ArticleGetRequest struct.
type ArticleGetRequest struct {
	Title string `validate:"required,max=45"`
}

// ArticleGetIDRequest struct.
type ArticleGetIDRequest struct {
	ID string `validate:"required,numeric"`
}

// ArticlePostAddRequest struct.
type ArticlePostAddRequest struct {
	Title   string `form:"title" validate:"required,max=45"`
	Content string `form:"content" validate:"required"`
}

// ArticleDeleteIDRequest struct.
type ArticleDeleteIDRequest struct {
	ID string `validate:"required,numeric"`
}

// ArticlePostVisenzeDiscoverSearchRequest struct.
type ArticlePostVisenzeDiscoverSearchRequest struct {
	Page        int            `form:"page" validate:"min=1"`
	ResultLimit int            `form:"result_limit" validate:"min=1"`
	File        multipart.File `form:"file" validate:"required"`
}
