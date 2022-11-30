package controller

import (
	"encoding/json"
	"es/helper"
	"es/model"
	"es/service"
	"net/http"
)

func InsertUserDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.UserDetails
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}
	if dataBody.FirstName == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "please enter firstName")
		return
	}
	if dataBody.LastName == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "please enter lastName")
		return
	}
	if dataBody.EmailId == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "please enter emailId")
		return
	}
	if dataBody.Role == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "please enter role")
		return
	}
	if result, err := service.Insert(dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record inserted successfully", helper.StatusCodeOK, result)
	}
}

func SearchByID(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	var dataBody model.UserDetails
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if dataBody.Id == "" && dataBody.FirstName == "" {
		helper.RespondWithError(w, http.StatusBadRequest, "please enter id for search")
		return
	}
	if result, err := service.FetchById(dataBody.Id, dataBody.FirstName); err != nil {
		helper.RespondWithError(w, http.StatusBadRequest, err.Error())
	} else {
		helper.RespondWithJson(w, http.StatusAccepted, "Record fetch successfully", helper.StatusCodeOK, result)
	}

}
