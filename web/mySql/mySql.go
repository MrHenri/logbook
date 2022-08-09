package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type Pet struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Owner   string `json:"owner"`
	Species string `json:"species"`
	Sex     string `json:"sex"`
	Birth   string `json:"birth"`
	Death   string `json:"death"`
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func dbConnect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "logbook_go"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	logError(err)

	return db
}

func scanPet(selDB *sql.Rows, pet Pet) Pet {
	var id int
	var name, owner, species, sex sql.NullString
	var birth, death sql.NullString

	err := selDB.Scan(&id, &name, &owner, &species, &sex, &birth, &death)

	logError(err)

	pet.Id = id
	pet.Name = name.String
	pet.Owner = owner.String
	pet.Species = species.String
	pet.Sex = sex.String
	pet.Birth = birth.String
	pet.Death = death.String

	return pet
}

func queryPetId(r *http.Request) *sql.Rows {
	db := dbConnect()
	petId := r.URL.Query().Get("id")

	selDb, err := db.Query("SELECT * FROM pet WHERE id = ?", petId)

	logError(err)

	defer db.Close() //ponto de anteção

	return selDb
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func index(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()

	selDB, err := db.Query("SELECT * FROM pet")

	logError(err)

	pet := Pet{}
	res := []Pet{}

	for selDB.Next() {
		pet = scanPet(selDB, pet)
		res = append(res, pet)
	}

	log.Println(res)

	err = tmpl.ExecuteTemplate(w, "Index", res)

	logError(err)
	defer db.Close()
}

func show(w http.ResponseWriter, r *http.Request) {
	selDB := queryPetId(r)

	pet := Pet{}

	for selDB.Next() {
		pet = scanPet(selDB, pet)
	}

	tmpl.ExecuteTemplate(w, "Show", pet)
}

//IS NOT WORKING
// func new(w http.ResponseWriter, r *http.Request) {
// 	tmpl.ExecuteTemplate(w, "New", nil)
// }

// func edit(w http.ResponseWriter, r *http.Request) {
// 	selDB := queryPetId(r)

// 	pet := Pet{}

// 	for selDB.Next() {
// 		scanPet(selDB, pet)
// 	}

// 	tmpl.ExecuteTemplate(w, "Edit", pet)
// }

// func insert(w http.ResponseWriter, r *http.Request) {
// 	db := dbConnect()

// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		owner := r.FormValue("owner")
// 		species := r.FormValue("species")
// 		sex := r.FormValue("sex")
// 		birth := r.FormValue("birth")
// 		death := r.FormValue("death")

// 		insForm, err := db.Prepare("INSERT INTO pet(name, owner, species, sex, birth, death) VALUES(?,?,?,?,?,?")

// 		logError(err)

// 		insForm.Exec(name, owner, species, sex, birth, death)
// 		log.Println("INSERT: name: " + name +
// 			" | owner: " + owner +
// 			" | species: " + species +
// 			" | sex: " + sex +
// 			" | birth: " + birth +
// 			" | death: " + death)
// 	}
// 	defer db.Close()
// 	http.Redirect(w, r, "/", 301)
// }

// func update(w http.ResponseWriter, r *http.Request) {
// 	db := dbConnect()

// 	if r.Method == "POST" {
// 		name := r.FormValue("name")
// 		owner := r.FormValue("owner")
// 		species := r.FormValue("species")
// 		sex := r.FormValue("sex")
// 		birth := r.FormValue("birth")
// 		death := r.FormValue("death")
// 		id := r.FormValue("uid")

// 		insForm, err := db.Prepare("UPDATE pet SET name=?, owner=?, species=?, sex=?, birth=?, death=? WHERE id=?")

// 		logError(err)

// 		insForm.Exec(name, owner, species, sex, birth, death, id)
// 		log.Println("UPDATE: name: " + name +
// 			" | owner: " + owner +
// 			" | species: " + species +
// 			" | sex: " + sex +
// 			" | birth: " + birth +
// 			" | death: " + death)
// 	}
// 	defer db.Close()
// 	http.Redirect(w, r, "/", 301)
// }

func delete(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()
	petId := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM pet WHERE id=?")

	logError(err)

	delForm.Exec(petId)
	log.Println("DELETE: pet id: " + petId)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", index)
	http.HandleFunc("/show", show)
	// http.HandleFunc("/new", new)
	// http.HandleFunc("/edit", edit)
	// http.HandleFunc("/insert", insert)
	// http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.ListenAndServe(":8080", nil)
}
