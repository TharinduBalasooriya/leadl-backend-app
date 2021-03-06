package routes

import (
	"fmt"
	"net/http"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/api"
	"github.com/gorilla/mux"

	//"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/middleware"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/websocket"
)

func LogRoutes() *mux.Router {
	var router = mux.NewRouter()
	router = mux.NewRouter().StrictSlash(true)


	//router.Use(middleware.LoggingMiddleware)
	//Get All Log files

	router.HandleFunc("/",func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw,"HomeRoute")
	})
	router.HandleFunc("/api/logs/{user}/", api.GetAllLog).Methods("GET")

	//getAllProjetcs
	router.HandleFunc("/api/projects/{user}/", api.GetAllProjects).Methods("GET")

	//upload file
	router.HandleFunc("/api/uploads/", api.HandleLogFileUpload).Methods("POST")

	router.HandleFunc("/api/uploadSripts/", api.HandleSrciptUpload).Methods("POST")

	//get log file content v2

	router.HandleFunc("/api/v2/content/{fileId}", api.GetLogFileContentv2).Methods("GET")

	//catch the log file updates

	router.HandleFunc("/api/updates", api.HandleFileUpdates).Methods("POST")

	//GetLogsByUserandProject
	router.HandleFunc("/api/logs/getByProject/{id}/", api.GetLogListByProjectID).Methods("GET")
	router.HandleFunc("/ws", websocket.WSPage).Methods("GET")

	//Invoke Interpreter
	router.HandleFunc("/api/executeLDEL/{fileId}", api.HandleInvokeELInterpreter).Methods("GET")

	

	//Craete a project
	router.HandleFunc("/api/project", api.HandleProject).Methods("POST")
	router.HandleFunc("/api/project/{id}", api.GetProjectDetails).Methods("GET")
	//fetch a project by userId
	router.HandleFunc("/api/projectV2/{user}", api.GetAllProjectsV2).Methods("Get")

	//update project
	router.HandleFunc("/api/project/update", api.HandleUpdateProjects).Methods("PUT")

	//delete project
	router.HandleFunc("/api/project/delete/{projectID}", api.HandleDeleteProjects).Methods("DELETE")

	//check project existance 
	router.HandleFunc("/api/project/check/{userId}/{projectName}", api.HandleExistProjects).Methods("GET")

	router.HandleFunc("/api/logs/activateLog/{fileId}", api.HandleActiavetLogFile).Methods("GET")

	router.HandleFunc("/api/debug/{projectId}", api.HandelDebugLDEL).Methods("GET")

	router.HandleFunc("/api/debug_save", api.HandleDebugProject).Methods("POST")
	

	return router
}
