package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/gorilla/mux"
)

func HandleProject(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var project datamodels.Project
	err := json.NewDecoder(r.Body).Decode(&project)

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	controller.ProjectSaveDetails(project)

	fmt.Print(r.Body)
	fmt.Println("project create Endpoint Hit\n")

}



func GetAllProjectsV2(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logs := controller.GetProjectsV2(params["user"])
	fmt.Print(logs)
	json.NewEncoder(w).Encode(logs)

}

func HandleUpdateProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var project datamodels.Project
	err := json.NewDecoder(r.Body).Decode(&project)
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		return
	}
	result := controller.UpdateProject(project)
	fmt.Print(result)
	json.NewEncoder(w).Encode(result)

}

func HandleDeleteProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	result := controller.DeleteProject(params["projectID"])
	fmt.Print(result)
}

func HandleExistProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	result := controller.CheckProject(params["userId"],params["projectName"])
	fmt.Print(result)
	json.NewEncoder(w).Encode(result)

}

func GetProjectDetails(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	result := controller.GetProjectDetails(params["id"])
	json.NewEncoder(w).Encode(result)
}

