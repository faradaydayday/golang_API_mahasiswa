package controllers

import (
	"encoding/json"
	"fiqri_muhammad/student-api/models"
	"fiqri_muhammad/student-api/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = student.Create()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusCreated, utils.Response{
		Status:  "success",
		Message: "Student created successfully",
		Data:    student,
	})
}

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := models.GetAllStudents()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, utils.Response{
		Status: "success",
		Data:   students,
	})
}

func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	student, err := models.GetStudentByID(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusNotFound, "Student not found")
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, utils.Response{
		Status: "success",
		Data:   student,
	})
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	var student models.Student
	err = json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	student.ID = id
	err = student.Update()
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Student updated successfully",
		Data:    student,
	})
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.ResponseWithError(w, http.StatusBadRequest, "Invalid student ID")
		return
	}

	err = models.DeleteStudent(id)
	if err != nil {
		utils.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.ResponseWithJSON(w, http.StatusOK, utils.Response{
		Status:  "success",
		Message: "Student deleted successfully",
	})
}
