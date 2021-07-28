package controller

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/service"
)

func GetLDELDebugResult(request datamodels.DebugRequest) datamodels.DebugResponse {




	 service.WriteDebugLogFile(request.LogFile)
	 service.WriteDebugScriptFile(request.LDELScript)
	 service.CreateDebugDefFile()
	 service.ConfigDebugDefsFile()

	var response datamodels.DebugResponse
	response = service.GetDebugResult()
	return response

}