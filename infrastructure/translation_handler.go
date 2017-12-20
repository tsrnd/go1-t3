package infrastructure

import (
	"path/filepath"

	"github.com/dalu/i18n"
)

const (
	// DirTranslation is translation path directory.
	DirTranslation = "translation"
)

// TranslationHandler struct.
type TranslationHandler struct {
	Middleware *i18n.I18nMiddleware
}

// NewTranslationHandler returns new TranslationHandler.
// repository: https://github.com/dalu/i18n
func NewTranslationHandler() *TranslationHandler {
	files, err := filepath.Glob(DirTranslation + "/*.json")
	if err != nil {
		panic(err)
	}
	c := i18n.Config{DefaultLanguage: GetConfigString("language.default"),
		Files: files,
		Debug: false,
	}
	return &TranslationHandler{Middleware: i18n.New(c)}
}
