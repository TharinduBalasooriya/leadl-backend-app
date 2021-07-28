package service

import (
	"encoding/base64"
	fclLib "github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func WriteDebugLogFile(logfile string) {

	// Open a new file for writing only
	file, err := os.OpenFile(
		"debug_env/log.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data ,err := base64.StdEncoding.DecodeString(logfile)
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}

func WriteDebugScriptFile(scriptFile string) {

	// Open a new file for writing only
	file, err := os.OpenFile(
		"debug_env/script.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0666,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data ,err := base64.StdEncoding.DecodeString(scriptFile)
	_, err = file.Write(data)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateDebugDefFile() {
	defFileTemplate, err := os.Open("util/templates/Defs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer defFileTemplate.Close()
	//Create New File
	newFilePath := "debug_env/Defs.txt"
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
func ConfigDebugDefsFile(){
	defFileLocation := "debug_env/Defs.txt"
	newDefScript:= "DEF	LDEL_SCRIPT_FILE			../src/debug_env/script.txt\n"
	newDefLogFile:= "DEF\tLDEL_LOG_FILE			../src/debug_env/log.txt\n"
	newDefResultLocation:= "DEF\tLDEL_RESULT_FILE			../src/debug_env/result.txt\n"


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

func GetDebugResult() (response datamodels.DebugResponse){
	defFilePath := "debug_env/Defs.txt"
	fclLib.NewELInterpretterWrapper().RunELInterpretter(defFilePath)

	// Open file for reading
	file, err := os.Open("debug_env/result.txt")
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	  response.Result = string(data)
	  return response

}
