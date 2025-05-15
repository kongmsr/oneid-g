package oneidg

import (
	"github.com/kongmsr/oneid-core/modelo"
)

func FailWithError(data error) *modelo.Response {
	return &modelo.Response{
		Code: modelo.SUCCESS,
		Data: data.Error(),
		Msg:  "Failed",
	}
}
