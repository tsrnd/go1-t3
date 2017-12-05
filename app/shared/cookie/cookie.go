
package cookie

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
	"strings"
)

/**
* Get message from cookie then clear them
*/
func GetMessage(w http.ResponseWriter, r *http.Request, name string) (message string) {
	c, err := r.Cookie(name)
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Println(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name:    name,
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		message = string(val)
	}
	return message
}

/**
* Get all message from cookie start with prefix then clear them
*/
func GetMessageStartWith(w http.ResponseWriter, r *http.Request, prefix string) (map[string]string) {
	messageMap :=  make(map[string]string)
	for _,cookie := range r.Cookies() {
		if strings.HasPrefix(cookie.Name, prefix) {
			messageMap[cookie.Name] = GetMessage(w, r, cookie.Name)
		}
	}
	return messageMap
}

/**
* Set message into cookie
*/
func SetMessage(w http.ResponseWriter, message string, name string) {
	msg := []byte(message)
	c := http.Cookie{
		Name:  name,
		Value: base64.URLEncoding.EncodeToString(msg)}
	http.SetCookie(w, &c)
}
