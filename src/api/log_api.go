package api

//API to handle main log operation
//Get file name , Get content

import (
	"encoding/json"
	"fmt"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

/*
	Read the content of log file
*/
func GetLogFileContent(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logs := controller.LogGetFileContent(params["user"], params["project"], params["logfileName"])
	json.NewEncoder(w).Encode(logs)

}

func GetLogFileContentv2(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//logs := controller.LogGetFileContent(params["user"], params["project"], params["logfileName"])
	log := controller.LogGetFileContentv2(params["fileId"])
	json.NewEncoder(w).Encode(log)

}

func GetAllLog(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logs := controller.GetFileList(params["user"])
	json.NewEncoder(w).Encode(logs)

}

func GetAllProjects(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	logs := controller.GetProjects(params["user"])
	json.NewEncoder(w).Encode(logs)

}

//

func GetLogListByUsernProject(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//logs := controller.GetProjects(params["user"])
	logs := controller.GetLogListByUsernProject(params["user"], params["project"])
	json.NewEncoder(w).Encode(logs)

}

func HandleActiavetLogFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result := controller.GetToActiveDir(params["fileId"])
	json.NewEncoder(w).Encode(result)

}
func HandleLogFileUpload(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	fmt.Println("File Upload Endpoint Hit")

	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("logFile")
	userName := r.FormValue("userName")
	projectName := r.FormValue("projectName")
	fileName := r.FormValue("fileName")
	fileId := r.FormValue("fileId")

	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	//aws upload path
	fullFilePath := "logs/" + userName + "/" + projectName + "/" + fileName

	//controller.LogUploadFiles(fullFilePath, file)
	controller.LogUploadFiles(fullFilePath, file)
	controller.LogSaveDetails(userName, projectName, fileName, fileId)
	//controller.Config_LDEL_DEF(fileName, fileId)

}

type Update struct {
	UserName    string `json:"userName"`
	ProjectName string `json:"project"`
	Data        string `json:"data"`
}

func HandleFileUpdates(w http.ResponseWriter, r *http.Request) {

	var newupdate Update
	_ = json.NewDecoder(r.Body).Decode(&newupdate)
	controller.HandleUpdateData(controller.Update(newupdate))

	json.NewEncoder(w).Encode(newupdate)

}

func HandleSrciptUpload(w http.ResponseWriter, r *http.Request) {

	fmt.Println("File Upload Endpoint Hit")
	r.ParseMultipartForm(10 << 20)

	file, handler, err := r.FormFile("myFile")
	fileId := r.FormValue("fileId")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("localstorage/"+fileId+"/script.txt", fileBytes, 0777)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintf(w, "Successfully Uploaded File\n")

}

func HandleInvokeELInterpreter(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)


	result := controller.ExecuteLDEL(params["fileId"])

	json.NewEncoder(w).Encode(result)


}
