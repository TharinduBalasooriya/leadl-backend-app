package controller

import (
	"fmt"
	"log"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/repository"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var projectrepo repository.ProjectRepository

func ProjectSaveDetails(project datamodels.Project) {

	exist, res := projectrepo.CheckprojectExist(project)

	if exist {

		fmt.Println("Project Already Exist")
		fmt.Print(res)

	} else {

		results, err := models.Project_Save_Details(project)

		if err != nil {
			log.Fatal(err)

		}

		id := results.(primitive.ObjectID)
		fmt.Println("Successfully inserted" + id.String())

	}

}



func GetProjectsV2(user string) []datamodels.Project {

	projectList := projectrepo.GetProjectsByUserV2(user)

	return projectList
}
