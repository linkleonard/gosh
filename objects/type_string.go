// Code generated by "stringer -type Type"; DO NOT EDIT.

package objects

import "strconv"

const _Type_name = "IntegerTypeFloatTypeBooleanTypeStringTypeFunctionTypeGoFunctionTypeContinueType"

var _Type_index = [...]uint8{0, 11, 20, 31, 41, 53, 67, 79}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
