package datamodels

type Project struct {
	
	ProjectName  string `json:"projectName"`
	Location     string `json:"location"`
	UserId       string `json:"userId"`
	ProjectId    string   `json:"ProjectId"` 
}
