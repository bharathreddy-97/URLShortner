package httphandlers

import (
	"net/http"

	persistence "github.com/bharathreddy-97/URLShortner/Persistence"
	"github.com/bharathreddy-97/URLShortner/utils"
)

func (h *Handlers) RedirectURL(rw http.ResponseWriter, r *http.Request) {
	shortURL := r.PathValue("id")
	if !utils.CheckForValidShortURL(shortURL) {
		handleError(rw, utils.E004)
		return
	}

	originalURL := persistence.ReturnOriginalURL(h.MapData, shortURL)
	if originalURL == "" {
		handleError(rw, utils.E005)
		return
	}

	http.Redirect(rw, nil, originalURL, http.StatusSeeOther)
}
