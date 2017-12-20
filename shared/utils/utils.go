package utils

import (
	"encoding/hex"
	"os"
	"reflect"

	"io"

	"golang.org/x/crypto/scrypt"
)

// FileCopy file.
func FileCopy(dstPath, srcPath string) error {
	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	return nil
}

// GetTag get struct tag.
func GetTag(typeData interface{}, tagName string) string {
	field, ok := reflect.TypeOf(typeData).Elem().FieldByName(tagName)
	if !ok {
		return ""
	}
	return string(field.Tag)
}

// CreateHashFromPassword scrypt hash password.
func CreateHashFromPassword(salt, password string) string {
	converted, _ := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, 16)
	return hex.EncodeToString(converted[:])
}
