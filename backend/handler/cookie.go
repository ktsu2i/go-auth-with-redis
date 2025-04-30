package handler

import (
	"backend/pkg"
	"net/http"
	"time"
)

const (
	domain   = "localhost"
	secure   = false
	sameSite = http.SameSiteLaxMode
)

func SetAuthCookie(w http.ResponseWriter, pair *pkg.TokenPair) {
	http.SetCookie(w, &http.Cookie{
		Name:     "access",
		Value:    pair.AccessToken,
		Path:     "/",
		Domain:   domain,
		Secure:   secure,
		SameSite: sameSite,
		Expires:  time.Now().Add(15 * time.Minute),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh",
		Value:    pair.RefreshToken,
		Path:     "/",
		Domain:   domain,
		Secure:   secure,
		SameSite: sameSite,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
	})
}

func ClearAuthCookie(w http.ResponseWriter) {
	for _, name := range []string{"access", "refresh"} {
		http.SetCookie(w, &http.Cookie{
			Name:     name,
			Value:    "",
			Path:     "/",
			Domain:   domain,
			MaxAge:   -1,
			Secure:   secure,
			SameSite: sameSite,
		})
	}
}
