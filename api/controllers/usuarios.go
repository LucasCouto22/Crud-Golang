package controllers

import (
	"api/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsUsuarios := models.BuscaTodosUsuarios()
	/*encoder := json.NewEncoder(w)
	encoder.Encode(todosOsUsuarios)*/
	temp.ExecuteTemplate(w, "index", todosOsUsuarios)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		nome := r.FormValue("nome")
		nivel := r.FormValue("nivel")
		idade := r.FormValue("idade")
		status := r.FormValue("status")

		idadeConvertida, err := strconv.Atoi(idade)
		if err != nil {
			log.Println("Erro na conversão de tipo de idade de string para int!")
		}

		models.CriaNovoUsuario(nome, nivel, status, idadeConvertida)
	}
	http.Redirect(w, r, "/", 301)

}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoUsuario := r.URL.Query().Get("id")
	models.DeletaUsuario(idDoUsuario)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoUsuario := r.URL.Query().Get("id")
	usuario := models.EditaUsuario(idDoUsuario)
	temp.ExecuteTemplate(w, "Edit", usuario)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		nivel := r.FormValue("nivel")
		idade := r.FormValue("idade")
		status := r.FormValue("status")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão de tipo de id de string para int!")
		}

		idadeConvertida, err := strconv.Atoi(idade)
		if err != nil {
			log.Println("Erro na conversão de tipo de idade de string para int!")
		}

		models.AtualizaUsuario(idConvertido, nome, nivel, idadeConvertida, status)
	}
	http.Redirect(w, r, "/", 301)
}
