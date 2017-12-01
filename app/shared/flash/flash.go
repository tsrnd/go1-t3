package flash

import (
	"encoding/base64"
	"net/http"
	"time"
	"encoding/json"
	"errors"
)

const (
	// FlashError is a bootstrap class
	FlashError = "alert-danger"
	// FlashSuccess is a bootstrap class
	FlashSuccess = "alert-success"
	// FlashNotice is a bootstrap class
	FlashNotice = "alert-info"
	// FlashWarning is a bootstrap class
	FlashWarning = "alert-warning"
	Name = "_flash"
)

// Flash Message
type Flash struct {
	Message string
	ClassCss   string
}

/**
*
* Set flassh message
**/
func SetFlash(w http.ResponseWriter, flash Flash) {
	value, _ := json.Marshal(flash)
	c := &http.Cookie{Name: Name, Value: encode(value)}
	http.SetCookie(w, c)
}

/**
*
* Get flassh message
*
**/
func GetFlash(w http.ResponseWriter, r *http.Request) (Flash, error) {
	flash := Flash{}
	c, err := r.Cookie(Name)
	if err != nil {
		switch err {
			case http.ErrNoCookie:
				return flash, nil
			default:
				return flash, err
		}
	}
	value, err := decode(c.Value)
	if value == nil {
		return flash, errors.New("")
	}
	if err != nil {
		return flash, err
	}
	json.Unmarshal([]byte(value), &flash)
	dc := &http.Cookie{Name: Name, MaxAge: -1, Expires: time.Unix(1, 0)}
	http.SetCookie(w, dc)
	return flash, nil
}

/**
*
* Encode String to byte
*
**/  
func encode(src []byte) string {
	return base64.URLEncoding.EncodeToString(src)
}

/**
*
* Decode String to byte
*
**/
func decode(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}