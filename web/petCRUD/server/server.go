package server

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/logbook/web/petCRUD/utils"
)

func connect() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "logbook_go"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+"?parseTime=true")
	utils.LogError(err)

	return db
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := connect()

	if r.Method == "POST" {

		name := r.FormValue("name")
		owner := r.FormValue("owner")
		species := r.FormValue("species")
		sex := r.FormValue("sex")
		birth := utils.ConvertDateToTime(r.FormValue("birth"))
		death := utils.ConvertDateToTime(r.FormValue("death"))

		insForm, err := db.Prepare("INSERT INTO pet(name, owner, species, sex, birth, death) VALUES(?,?,?,?,?,NULLIF(?,'0000-00-00'))")

		utils.LogError(err)

		sqlResult, err := insForm.Exec(name, owner, species, sex, birth, death)

		utils.LogError(err)
		log.Println("INSERT: name: " + name +
			" | owner: " + owner +
			" | species: " + species +
			" | sex: " + sex +
			" | birth: " + birth.Format(utils.LayoutDate) +
			" | death: " + death.Format(utils.LayoutDate))
		log.Println(sqlResult)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func QueryPetId(r *http.Request) *sql.Rows {
	db := connect()
	petId := r.URL.Query().Get("id")

	selDb, err := db.Query("SELECT * FROM pet WHERE id = ?", petId)

	utils.LogError(err)

	defer db.Close()
	return selDb
}

func QueryAllPets() *sql.Rows {
	db := connect()

	selDb, err := db.Query("SELECT * FROM pet")

	utils.LogError(err)

	defer db.Close()
	return selDb
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := connect()

	if r.Method == "POST" {
		name := r.FormValue("name")
		owner := r.FormValue("owner")
		species := r.FormValue("species")
		sex := r.FormValue("sex")
		birth := utils.ConvertDateToTime(r.FormValue("birth"))
		death := utils.ConvertDateToTime(r.FormValue("death"))
		id := r.FormValue("uid")

		insForm, err := db.Prepare("UPDATE pet SET name=?, owner=?, species=?, sex=?, birth=?, death=NULLIF(?,'0000-00-00') WHERE id=?")

		utils.LogError(err)

		sqlResult, err := insForm.Exec(name, owner, species, sex, birth, death, id)
		utils.LogError(err)
		log.Println("UPDATE: name: " + name +
			" | owner: " + owner +
			" | species: " + species +
			" | sex: " + sex +
			" | birth: " + birth.Format(utils.LayoutDate) +
			" | death: " + death.Format(utils.LayoutDate))
		log.Println(sqlResult)
	}
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := connect()
	petId := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM pet WHERE id=?")

	utils.LogError(err)

	delForm.Exec(petId)
	log.Println("DELETE: pet id: " + petId)
	defer db.Close()
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
