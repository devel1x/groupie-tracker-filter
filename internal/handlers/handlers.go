package handlers

import (
	"fmt"
	"groupie/internal/models"
	"groupie/internal/render"
	"groupie/internal/service"
	"net/http"
	"strings"
	"time"
)

var (
	t models.TemplateData
)
var StaticHandler = http.StripPrefix("/static/", preventDirListing(http.FileServer(http.Dir("./web/static"))))

func (h *handler) Cash() *handler {
	if h.useCashTimer == nil {
		// Set useCashTimer to a new timer that triggers every 20 minutes
		h.useCashTimer = time.NewTimer(1 * time.Second)

		// Start a goroutine to handle the timer expiration
		go func() {
			for {
				select {
				case <-h.useCashTimer.C:
					// Timer expired, set useCash to false
					h.useCash = false

					// Reset the timer for the next 20 minutes
					h.useCashTimer.Reset(1 * time.Second)
				}
			}
		}()
	}

	return h
}

func (h *handler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if h.useCash {
		fmt.Print("USED CASH")
		render.RenderTemplate(w, "index.page.html", &models.TemplateData{
			Data: t.Data,
		})
	} else {
		h.useCash = true
		if r.URL.Path != "/" {
			ErrorHandler(w, r, models.Errors[404])
			return
		}
		switch r.Method {
		case "GET":
			Artists, err := h.ServiceI.Artists()
	
			// fmt.Print(Artists)
			if err != nil {
				ErrorHandler(w, r, models.Errors[500])
				return
			}
			loc := service.GetLocF(Artists)
			for key, value := range loc {
				fmt.Println(key)
				fmt.Println(value)
				fmt.Println()
			}
			t = models.TemplateData{
				Data: make(map[string]interface{}),
			}

			t.Data["Artists"] = Artists

			render.RenderTemplate(w, "index.page.html", &t)

		case "POST":
			cDateF := r.URL.Query().Get("cDateF")
			cDateT := r.URL.Query().Get("cDateT")
			loc := r.URL.Query().Get("loc")
			members := r.URL.Query().Get("members")
			aDateF := r.URL.Query().Get("aDateF")
			aDateT := r.URL.Query().Get("aDateT")
			Artists, err := h.ServiceI.Artists()
			f := service.NewFilter(Artists, aDateF, aDateT, cDateF, cDateT, members, loc)
			Artists, err = f.Filter()

			if err != nil {
				ErrorHandler(w, r, models.Errors[500])
				return
			}
			t = models.TemplateData{
				Data: make(map[string]interface{}),
			}

			t.Data["Artists"] = Artists

			render.RenderTemplate(w, "index.page.html", &t)
		default:
			ErrorHandler(w, r, models.Errors[405])
			return
		}
	}
}

func preventDirListing(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, "/") || len(r.URL.Path) == 0 {
			ErrorHandler(w, r, models.Errors[404])
			return
		}
		handler.ServeHTTP(w, r)
	}
}

func (h *handler) ArtistHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/artist" {
		ErrorHandler(w, r, models.Errors[404])
		return
	}
	if r.Method != "GET" {
		ErrorHandler(w, r, models.Errors[405])
		return
	}
	artistID := r.URL.Query().Get("id")

	Artist, err := h.ServiceI.Artist(artistID)
	if err != nil {
		for code, msg := range models.Errors {
			if msg.Msg == err.Error() {
				ErrorHandler(w, r, models.Errors[code])
				return
			}
		}
	}
	Data := make(map[string]interface{})

	Data["Artist"] = Artist

	render.RenderTemplate(w, "artist.page.html", &models.TemplateData{
		Data: Data,
	})
}

func ErrorHandler(w http.ResponseWriter, r *http.Request, myErr models.Errs) {
	w.WriteHeader(myErr.Code)

	render.RenderTemplate(w, "error.page.html", &models.TemplateData{
		Error:   myErr.Code,
		Warning: myErr.Msg,
	})

}
