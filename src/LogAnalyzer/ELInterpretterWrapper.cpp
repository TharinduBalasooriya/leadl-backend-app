//
// Created by Isini_Danajana on 2021-07-12.
//

#include <string>
#include "ELInterpretterWrapper.h"
#include "ELInterpretter.h"

void ELInterpretterWrapper::RunELInterpretter(const char*  defFilepath) {
    ELInterpretter intp;
    intp.EvaluateCase(defFilepath);

}