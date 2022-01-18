package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/infortech07/crudApi/src/commons"
	"github.com/infortech07/crudApi/src/models"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	persons := []models.Person{}
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&persons)
	json, _ := json.Marshal(persons)
	commons.SendResponse(w, http.StatusOK, json)
}

func Get(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	id := mux.Vars(r)["id"]
	db := commons.GetConnection()
	defer db.Close()
	db.Find(&person, id)
	if person.ID > 0 {
		json, _ := json.Marshal(person)
		commons.SendResponse(w, http.StatusOK, json)
	} else {
		commons.SendError(w, http.StatusNotFound)
	}
}

func Save(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	db := commons.GetConnection()
	defer db.Close()

	error := json.NewDecoder(r.Body).Decode(&person)

	if error != nil {
		log.Fatal(error)
		commons.SendError(w, http.StatusBadRequest)
		return
	}

	error = db.Save(&person).Error
	if error != nil {
		log.Fatal(error)
		commons.SendError(w, http.StatusInternalServerError)
		return
	}

	json, _ := json.Marshal(person)
	commons.SendResponse(w, http.StatusOK, json)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	person := models.Person{}
	db := commons.GetConnection()
	defer db.Close()
	id := mux.Vars(r)["id"]
	db.Find(&person, id)

	if person.ID > 0 {
		db.Delete(person)
		commons.SendResponse(w, http.StatusOK, []byte(`{}`))
	} else {
		commons.SendError(w, http.StatusNoContent)
	}

}
