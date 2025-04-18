package interfaces

import (
	"fmt"
	"github.com/duke-git/lancet/v2/convertor"
	"ruoyi-gin/src/ruoyi-common/core/reflection"
)

func JsonSerializableImpl(jsonSerializable interface{}) string {
	tmp, _ := reflection.ReflectionLowerCaseFieldNames(jsonSerializable)
	result, err := convertor.ToJson(tmp)

	if err != nil {
		fmt.Printf("%v", err)
	}
	return result
}
