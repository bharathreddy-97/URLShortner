package httphandlers

import "net/http"

func ShortenURL(rw http.ResponseWriter, r *http.Request) {
	idValue := r.PathValue("id")
	rw.Write([]byte(idValue))
	rw.WriteHeader(200)
}
