package controller

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/models"
)

func GetLDELDebugResult(request datamodels.DebugRequest) datamodels.DebugResponse {




	 models.WriteDebugLogFile(request.LogFile)
	 models.WriteDebugScriptFile(request.LDELScript)
	 models.CreateDebugDefFile()
	 models.ConfigDebugDefsFile()

	var response datamodels.DebugResponse
	response = models.GetDebugResult()
	return response

}