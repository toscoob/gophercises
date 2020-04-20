package cyoa

import (
	"net/http"
	"html/template"
)

type StoryHandler struct {
	Story Story
	Tmpl *template.Template
}

func (h *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newArc := r.URL.Query().Get("arc")
	if newArc == "" {
		newArc = "intro"
	}
	if arc, ok := h.Story[newArc]; ok {
		//_, _ = fmt.Fprintf(w, "%s arc: %s", h.CurrentArc, &arc)
		err := h.Tmpl.Execute(w, arc)
		if err != nil {
			http.Error(w, "Something went wrong...", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found.", http.StatusNotFound)
}
