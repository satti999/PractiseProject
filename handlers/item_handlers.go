package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/satti999/PractiseProject/models"
	"github.com/satti999/PractiseProject/services"
	"github.com/satti999/PractiseProject/utils"

	"github.com/gorilla/mux"
)

var itemService = services.NewItemService()

func GetItems(w http.ResponseWriter, r *http.Request) {
	items, err := itemService.GetAllItems()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, items)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	item, err := itemService.GetItemByID(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, item)
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	createdItem, err := itemService.CreateItem(item)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, createdItem)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	updatedItem, err := itemService.UpdateItem(id, item)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, updatedItem)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid ID")
		return
	}

	err = itemService.DeleteItem(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
