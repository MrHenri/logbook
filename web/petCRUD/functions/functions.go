package functions

import (
	"html/template"
	"net/http"

	"github.com/logbook/web/petCRUD/model"
	"github.com/logbook/web/petCRUD/server"
	"github.com/logbook/web/petCRUD/services"
	"github.com/logbook/web/petCRUD/utils"
)

var tmpl = template.Must(template.ParseGlob("components/*"))

func Index(w http.ResponseWriter, r *http.Request) {

	selDB := server.QueryAllPets()

	pet := model.Pet{}
	res := []model.Pet{}

	for selDB.Next() {
		pet = services.ScanPet(selDB, pet)
		res = append(res, pet)
	}

	err := tmpl.ExecuteTemplate(w, "Index", res)

	utils.LogError(err)
}

func Show(w http.ResponseWriter, r *http.Request) {
	selDB := server.QueryPetId(r)

	pet := model.Pet{}

	for selDB.Next() {
		pet = services.ScanPet(selDB, pet)
	}

	tmpl.ExecuteTemplate(w, "Show", pet)
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	selDB := server.QueryPetId(r)

	pet := model.Pet{}

	for selDB.Next() {
		pet = services.ScanPet(selDB, pet)
	}

	tmpl.ExecuteTemplate(w, "Edit", pet)
}
