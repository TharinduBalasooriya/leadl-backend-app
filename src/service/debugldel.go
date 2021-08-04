package service

import (
	"encoding/base64"
	"fmt"
	fclLib "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/LogAnalyzer"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func WriteDebugLogFile(request datamodels.DebugRequest) {


	err := os.MkdirAll("debug_env/"+request.ProjectId, 0755)
	if err  != nil{
		fmt.Println(err)
	}
	// Open a new file for writing only
	file, err := os.OpenFile(
		"debug_env/"+request.ProjectId+"/log.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data ,err := base64.StdEncoding.DecodeString(request.LogFile)
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}

func WriteDebugScriptFile(request datamodels.DebugRequest) {

	err := os.MkdirAll("debug_env/"+request.ProjectId, 0755)
	if err  != nil{
		fmt.Println(err)
	}
	// Open a new file for writing only
	file, err := os.OpenFile(
		"debug_env/"+request.ProjectId+"/script.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data ,err := base64.StdEncoding.DecodeString(request.LDELScript)
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateDebugDefFile(projectId string) {
	defFileTemplate, err := os.Open("util/templates/Defs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer defFileTemplate.Close()
	//Create New File
	newFilePath := "debug_env/" +projectId+"/Defs.txt"
	newFile, err := os.Create(newFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	// Copy the bytes to destination from source
	bytesWritten, err := io.Copy(newFile, defFileTemplate)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Copied %d bytes.", bytesWritten)

	// Commit the file contents
	// Flushes memory to disk
	err = newFile.Sync()
	if err != nil {
		log.Fatal(err)
	}

}
func ConfigDebugDefsFile(projectID string){
	defFileLocation := "debug_env/" +projectID+"/Defs.txt"
	newDefScript:= "DEF	LDEL_SCRIPT_FILE			../src/debug_env/"+projectID+"/script.txt\n"
	newDefLogFile:= "DEF\tLDEL_LOG_FILE			../src/debug_env/"+projectID+"/log.txt\n"
	newDefResultLocation:= "DEF\tLDEL_RESULT_FILE			../src/debug_env/"+projectID+"/result.txt\n"


	defFile,err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644)

	if err != nil{
		log.Println(err)
	}

	defer defFile.Close();
	if _, err := defFile.WriteString(newDefScript); err != nil {
		log.Println(err)
	}

	if _, err := defFile.WriteString(newDefLogFile); err != nil {
		log.Println(err)
	}

	if _, err := defFile.WriteString(newDefResultLocation); err != nil {
		log.Println(err)
	}

}

func GetDebugResult(projectId string) (response datamodels.DebugResponse){
	defFilePath := "debug_env/"+projectId+"/Defs.txt"
	fclLib.NewELInterpretterWrapper().RunELInterpretter(defFilePath)

	// Open file for reading
	file, err := os.Open("debug_env/"+projectId+"/result.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	  response.Result = string(data)
	  defer cleanProjectFolder(projectId,file)
	  return response

}

func cleanProjectFolder(projectID string,file *os.File){
	file.Close()
	err := os.RemoveAll("debug_env/"+projectID)
    if err != nil {
        log.Fatal(err)
    }
}