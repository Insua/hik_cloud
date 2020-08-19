package context

import (
	"github.com/Insua/hik_cloud/config"
	"github.com/Insua/hik_cloud/credential"
)

type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
