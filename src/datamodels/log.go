package datamodels

type Log struct {
	Username    string `json:"username"`
	ProjectId string `json:"projectId"`
	LogFileName string `json:"logfilename"`
	LastUpdate  string `json:"lastupdate"`
	FileId  	string `json:"fileId"`
}
