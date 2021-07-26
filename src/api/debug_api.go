package api

import (
	"encoding/json"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/controller"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/datamodels"
	"net/http"
)

func HandelDebugLDEL(w http.ResponseWriter , r *http.Request ){

	decoder := json.NewDecoder(r.Body)
	var debugRequest datamodels.DebugRequest
	err := decoder.Decode(&debugRequest)
	if err != nil{
		println(err.Error())
	}
	result := controller.GetLDELDebugResult(debugRequest)
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		println(err.Error())
	}


}

