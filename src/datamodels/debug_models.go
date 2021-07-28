package datamodels

type DebugRequest struct {
	ProjectId string `json:"project_id"`
	LDELScript string `json:"ldel_script"`
	LogFile    string `json:"log_file"`

}

type DebugResponse struct {
	Result string `json:"response"`
}
