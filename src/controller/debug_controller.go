package controller

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"
)

func GetLDELDebugResult(projectId string) datamodels.DebugResponse {





	 service.CreateDebugDefFile(projectId)
	 service.ConfigDebugDefsFile(projectId)

	var response datamodels.DebugResponse
	response = service.GetDebugResult(projectId)
	return response

}


func SaveDebugProject(request datamodels.DebugRequest){

	service.WriteDebugLogFile(request)
	service.WriteDebugScriptFile(request)

}
