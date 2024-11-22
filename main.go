package main

import (
    "html/template"
    "net/http"
    "log"
)

type PortfolioData struct {
    Name       string
    Title      string
    Summary    string
    Project1   string
    Project2   string
    Project3   string
    Image      string
    ThemeColor string
    Layout     string
    HoverColor string  // Adicionando HoverColor
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    t, err := template.ParseFiles("templates/" + tmpl)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func renderForm(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        portfolio := PortfolioData{
            Name:       r.FormValue("name"),
            Title:      r.FormValue("title"),
            Summary:    r.FormValue("summary"),
            Project1:   r.FormValue("project1"),
            Project2:   r.FormValue("project2"),
            Project3:   r.FormValue("project3"),
            Image:      r.FormValue("image"),
            ThemeColor: r.FormValue("themeColor"),
            Layout:     r.FormValue("layout"),
            HoverColor: r.FormValue("hoverColor"),  // Capturando HoverColor
        }

        if portfolio.Layout == "moderno" {
            renderTemplate(w, "portfolio-moderno.html", portfolio)
        } else {
            renderTemplate(w, "portfolio-classico.html", portfolio)
        }
    } else {
        renderTemplate(w, "form.html", nil)
    }
}

func main() {
    http.HandleFunc("/", renderForm)
    log.Println("Iniciando servidor na porta 8080...")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("Erro ao iniciar servidor: ", err)
    }
}


