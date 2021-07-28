package service

import (

	//Importing file storage utility

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"
)

/*
This package containes all business logic log file

*/

var projectrepo repository.ProjectRepository

//Save Log Details in mongo db
func Project_Save_Details(project datamodels.Project) (interface{}, error) {
	resultID, err := projectrepo.SaveProject(project)
	return resultID, err

}
