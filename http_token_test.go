package oneidg

import (
	"fmt"
	"github.com/kongmsr/oneid-core/modelo"
	"testing"
)

var c *modelo.OneidConf

func init() {
	c = &modelo.OneidConf{
		AuthenticationUrl: "http://localhost:8000/oneid/token",
		AccessKeyId:       "058d1d5f-a23f-xxxx-8e9e-d31f412786ce",
		AccessKeySecret:   "972eff62-xxxx-4757-b332-b32b094a7aa5",
		Subject:           "wpp-admin",
	}
}

func TestGetOneidToken(t *testing.T) {

	resp := GetOneidToken(&modelo.OneidConf{})
	fmt.Println(resp)

	resp2 := GetOneidToken(&modelo.OneidConf{}, "zack")
	fmt.Println(resp2)
}
