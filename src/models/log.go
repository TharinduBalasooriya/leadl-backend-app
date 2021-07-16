package models

import (

	//Importing file storage utility
	"archive/zip"
	"fmt"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
	filestorageHandler "github.com/TharinduBalasooriya/LogAnalyzerBackend/src/util/filestorage"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	fclLib "github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
)

/*
This package containes all business logic log file

*/

var logrepo repository.LogRepository

func unzipLogfile(Logs string) {

	fmt.Println("temp/" + Logs + os.Getenv("BUCKET_ITEM_EXT"))

	zipReader, err := zip.OpenReader("temp/" + Logs + os.Getenv("BUCKET_ITEM_EXT"))
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// Iterate through each file/dir found in
	for _, file := range zipReader.Reader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		// Specify what the extracted file name should be.
		// You can specify a full path or a prefix
		// to move it to a different directory.
		// In this case, we will extract the file from
		// the zip to a file of the same name.
		targetDir := "./temp"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since not a directory
			log.Println("Extracting file:", file.Name)

			// Open an output file for writing
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}


func unzipLogfilev2(logfilename string) {

	fmt.Println("temp/" + logfilename + os.Getenv("ARCHIVED_EXT"))

	zipReader, err := zip.OpenReader("temp/" + logfilename + os.Getenv("ARCHIVED_EXT"))
	if err != nil {
		log.Fatal(err)
	}
	defer zipReader.Close()

	// Iterate through each file/dir found in
	for _, file := range zipReader.Reader.File {
		// Open the file inside the zip archive
		// like a normal file
		zippedFile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		// Specify what the extracted file name should be.
		// You can specify a full path or a prefix
		// to move it to a different directory.
		// In this case, we will extract the file from
		// the zip to a file of the same name.
		targetDir := "./temp"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		// Extract the item (or create directory)
		if file.FileInfo().IsDir() {
			// Create directories to recreate directory
			// structure inside the zip archive. Also
			// preserves permissions
			log.Println("Creating directory:", extractedFilePath)
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			// Extract regular file since not a directory
			log.Println("Extracting file:", file.Name)

			// Open an output file for writing
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()

			// "Extract" the file by copying zipped file
			// contents to the output file
			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

/*
	Upload a file
*/
func Log_uploadFiles(fs filestorageHandler.FileStorage) {

	
	err := fs.AddFiles() // calling add files function of the file storage
	if err != nil {
		log.Fatal(err)
	}

}

//Save Log Details in mongo db
func Log_Save_Details(log datamodels.Log)(interface{},error){


	


	resultID,err :=logrepo.SaveLog(log);
	return resultID,err;
	

}

func Log_GetContent(file_object filestorageHandler.File, logfileName string) []byte {

	//fileExtension := os.Getenv("FILE_EXT")
	fileExtension := ".txt"

	err := file_object.GetContent()
	if err != nil {
		log.Fatal(err)
	}
	unzipLogfile(logfileName)

	data, err := ioutil.ReadFile("temp/" + logfileName + fileExtension)
	if err != nil {
		panic(err)
	}

	return data

}


func Log_GetContentV2(file_object filestorageHandler.File, logfileName string) []byte {

	//fileExtension := os.Getenv("FILE_EXT")
	//fileExtension := ".txt"

	err := file_object.GetContent()
	if err != nil {
		log.Fatal(err)
	}
	unzipLogfilev2(logfileName)

	data, err := ioutil.ReadFile("temp/" + logfileName)
	if err != nil {
		panic(err)
	}

	return data

}


//Create local storage derectories

func Log_CreateDirectory(fileId string){

	path := "localstorage/" + fileId
	err := os.MkdirAll(path,0755);

	if err != nil{
		log.Fatal(err)
	}
}


func Log_GetDefFileTempalte(fileId string){

	//Open DefFile template

	defFileTemplate, err := os.Open("util/templates/Defs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer defFileTemplate.Close();


	//Create New File

	newFilePath := "localstorage/" + fileId +"/Defs.txt"
	newFile, err := os.Create(newFilePath )
	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	//Copy bytes create a new Template


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

func Log_Execute_LDEL(fileId string){

	defFilePath := "localstorage/"  + fileId + "/Defs.txt";

	fclLib.NewELInterpretterWrapper().RunELInterpretter(defFilePath);



}


func Log_Read_Result(fileId string)(interface{}){
	resultFilePath := "localstorage/"  + fileId + "/result.txt";

	// Open file for reading
    file, err := os.Open(resultFilePath)
    if err != nil {
        log.Fatal(err)
    }
	defer file.Close();

	data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatal(err)
    }

	type Response struct{
		FileId string `json:"fileId"`
		Result string `json:"result"`
	}

	response := Response{}

	response.FileId = fileId;
	response.Result = string(data)



	return response;






}


func Log_Append_LDEL_ScriptLocation(fileId string){

	defFileLocation := "localstorage/"+fileId+"/Defs.txt"
	newDef:= "DEF	LDEL_SCRIPT_FILE			../src/localstorage/" + fileId + "/script.txt\n"


	defFile,err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644) 

	if err != nil{
		log.Println(err)
	}

	defer defFile.Close();

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}



}


func Log_Append_LDEL_LogFileLocation(fileId string, fileName string){


	defFileLocation := "localstorage/"+fileId+"/Defs.txt"
	newDef:= "DEF	LDEL_LOG_FILE				../src/localstorage/" + fileId + "/" + fileName + "\n";


	defFile,err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644) 

	if err != nil{
		log.Println(err)
	}

	defer defFile.Close();

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}

}

func Log_Append_LDEL_ResultLocation(fileId string){


	defFileLocation := "localstorage/"+fileId+"/Defs.txt"
	newDef:= "DEF	LDEL_RESULT_FILE			../src/localstorage/" + fileId + "/result.txt\n"


	defFile,err := os.OpenFile(defFileLocation,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,0644) 

	if err != nil{
		log.Println(err)
	}

	defer defFile.Close();

	if _, err := defFile.WriteString(newDef); err != nil {
		log.Println(err)
	}



}