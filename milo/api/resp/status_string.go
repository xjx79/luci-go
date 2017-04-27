// Code generated by "stringer -type=Status,ComponentType,Verbosity"; DO NOT EDIT.

package resp

import "fmt"

const _Status_name = "NotRunRunningSuccessFailureWarningInfraFailureExceptionExpiredDependencyFailureWaitingDependency"

var _Status_index = [...]uint8{0, 6, 13, 20, 27, 34, 46, 55, 62, 79, 96}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return fmt.Sprintf("Status(%d)", i)
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}

const _ComponentType_name = "RecipeStepSummary"

var _ComponentType_index = [...]uint8{0, 6, 10, 17}

func (i ComponentType) String() string {
	if i < 0 || i >= ComponentType(len(_ComponentType_index)-1) {
		return fmt.Sprintf("ComponentType(%d)", i)
	}
	return _ComponentType_name[_ComponentType_index[i]:_ComponentType_index[i+1]]
}

const _Verbosity_name = "NormalHiddenInteresting"

var _Verbosity_index = [...]uint8{0, 6, 12, 23}

func (i Verbosity) String() string {
	if i < 0 || i >= Verbosity(len(_Verbosity_index)-1) {
		return fmt.Sprintf("Verbosity(%d)", i)
	}
	return _Verbosity_name[_Verbosity_index[i]:_Verbosity_index[i+1]]
}
