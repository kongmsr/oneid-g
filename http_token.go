package oneidg

import (
	"encoding/json"
	"errors"
	"github.com/kongmsr/oneid-core/httpo"
	"github.com/kongmsr/oneid-core/modelo"
	"net/http"
)

func GetOneidToken(c *modelo.OneidConf, data ...string) *modelo.Response {
	if len(c.Subject) == 0 || len(c.AuthenticationUrl) == 0 ||
		len(c.AccessKeyId) == 0 || len(c.AccessKeySecret) == 0 {
		return FailWithError(errors.New("oneid conf is missing"))
	}

	req := &modelo.JwtReq{
		KeyID:            c.AccessKeyId,
		EncodedKeySecret: c.AccessKeySecret,
		Subject:          c.Subject,
	}
	if len(data) > 0 {
		req.Value = &data[0]
	}

	if bs, err := json.Marshal(req); err != nil {
		return FailWithError(errors.New("construct oneid token req error: " + err.Error()))
	} else {
		if resp, err := httpo.DoReq(http.MethodPost, c.AuthenticationUrl, bs); err != nil {
			return FailWithError(errors.New("call [" + c.AuthenticationUrl + "] api error: " + err.Error()))
		} else {
			return resp
		}
	}
}
