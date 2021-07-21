package repository

import (
	"context"
	"fmt"
	"log"

	"time"

	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var project_collection = new(mongo.Collection)

const ProjectCollection = "Project"

/*
	Initalizing database configeration
*/

func init() {

	fmt.Println("Database Connection Established")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb+srv://tharindu:tharindu@cluster0.vnll5.mongodb.net/myFirstDB?retryWrites=true&w=majority")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	project_collection = client.Database("leadldb").Collection(ProjectCollection)

}

type ProjectRepository struct{}

func (l *ProjectRepository) SaveProject(project datamodels.Project) (interface{}, error) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := project_collection.InsertOne(ctx, project)
	fmt.Println("Inserted a single project: ", result.InsertedID)
	return result.InsertedID, err
}

func (l *ProjectRepository) CheckprojectExist(project datamodels.Project) (bool, string) {

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result := project_collection.FindOne(ctx, bson.M{"projectName": project.ProjectName})

	var resultLog bson.M
	result.Decode(&resultLog)

	/*
		check existences
	*/
	if len(resultLog) == 0 {

		return false, ""

	} else {
		stringObjectId := resultLog["_id"].(primitive.ObjectID).Hex()
		return true, stringObjectId
	}

}
