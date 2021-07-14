package test

import(
	"testing"
	"github.com/TharinduBalasooriya/LogAnalyzerBackend/LogAnalyzer"
	
)

func connectQL(t *testing.T){

	Tests := LDELib.NewELInterpretterWrapper();
    Tests.RunELInterpretter("/home/codebind/go/src/github.com/isini/Code/tests/LDEL_test1/Defs.txt");




}


