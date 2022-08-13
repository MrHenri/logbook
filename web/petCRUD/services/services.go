package services

import (
	"database/sql"

	"github.com/logbook/web/petCRUD/model"
	"github.com/logbook/web/petCRUD/utils"
)

func ScanPet(selDB *sql.Rows, pet model.Pet) model.Pet {
	var id int
	var name, owner, species, sex sql.NullString
	var birth, death sql.NullTime

	err := selDB.Scan(&id, &name, &owner, &species, &sex, &birth, &death)

	utils.LogError(err)

	pet.Id = id
	pet.Name = name.String
	pet.Owner = owner.String
	pet.Species = species.String
	pet.Sex = sex.String
	pet.Birth = utils.ConvertTimeToDate(birth)
	pet.Death = utils.ConvertTimeToDate(death)

	return pet
}
