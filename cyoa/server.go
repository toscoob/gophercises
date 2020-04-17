package cyoa

import (
	"fmt"
	"net/http"
	"html/template"
)

type StoryHandler struct {
	Story map[string]StoryEntry
	CurrentArc string
	Tmpl template.Template
}

func (h *StoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	newArc := r.URL.Query().Get("arc")
	if arc, ok := h.Story[newArc]; ok {
		//_, _ = fmt.Fprintf(w, "%s arc: %s", h.CurrentArc, &arc)
		_ = h.Tmpl.Execute(w, arc)
	} else {
		_, _ = fmt.Fprintf(w, "Could not locate arc %s in story", newArc)
	}
}
