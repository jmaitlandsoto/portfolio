package handlers

import (
    "html/template"
    "net/http"
    "path/filepath"
)

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    tmplPath := filepath.Join("internal", "templates", tmpl)
    t, err := template.ParseFiles(tmplPath, "internal/templates/base.html")
    if err != nil {
        http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    t.ExecuteTemplate(w, "base", data)
}

func Home(w http.ResponseWriter, r *http.Request) {
    renderTemplate(w, "index.html", nil)
}

func Projects(w http.ResponseWriter, r *http.Request) {
    projects := []string{"Project A", "Project B", "Project C"}
    tmplPath := filepath.Join("internal", "templates", "partials", "projects.html")
    t, err := template.ParseFiles(tmplPath)
    if err != nil {
        http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
        return
    }
    t.Execute(w, projects)
}
