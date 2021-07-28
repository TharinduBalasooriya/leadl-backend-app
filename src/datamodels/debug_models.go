package datamodels

type DebugRequest struct {
	LDELScript string `json:"ldel_script"`
	LogFile    string `json:"log_file"`
}

type DebugResponse struct {
	Result string `json:"response"`
}
