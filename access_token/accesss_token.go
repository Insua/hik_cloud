package access_token

import "github.com/Insua/hik_cloud/context"

type AccessToken struct {
	*context.Context
}

func NewAccessToken(context *context.Context) *AccessToken {
	at := new(AccessToken)
	at.Context = context
	return at
}

func (ak *AccessToken) AccessToken() (token string, err error) {
	token, err = ak.GetAccessToken()
	return
}
