%module LDELib
%{
#include "ELInterpretterWrapper.h"
%}

%include <typemaps.i>
%include "std_string.i"
%include "std_vector.i"


namespace std {
        %template(StringVector) vector<string>;
        %template(ByteVector) vector<char>;
}

%include "ELInterpretterWrapper.h"