package system

import (
	systemDTO "github.com/MuserQuantity/gin-project-example-without-users/dal/dto/system"
	systemModel "github.com/MuserQuantity/gin-project-example-without-users/dal/model/system"
)

func TestDTO(testStruct *systemModel.TestStruct) *systemDTO.TestDTO {
	return &systemDTO.TestDTO{
		Test: testStruct.Test,
	}
}
