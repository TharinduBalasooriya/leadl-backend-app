package test

import (
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/src/LogAnalyzer"
	"testing"
)

func connectQL(t *testing.T){

	Tests := LDELib.NewELInterpretterWrapper();
    Tests.RunELInterpretter("D:\\Tracified\\LogAnalyzer\\Leedl-backend\\src\\debug_env\\Defs.txt");
}




