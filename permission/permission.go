package permission

import "github.com/Insua/hik_cloud/http"

type Permission struct {
	Http *http.Http
}

func NewPermission(http *http.Http) *Permission {
	p := new(Permission)
	p.Http = http
	return p
}
