package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
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
	fmt.Println("project create Endpoint Hit")
	
	

}