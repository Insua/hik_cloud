package property

import (
	"github.com/Insua/hik_cloud/http"
	"github.com/gogf/gf/util/gconv"
)

const (
	createUrl = "https://api2.hik-cloud.com/api/v1/estate/system/property"
	updateUrl = "https://api2.hik-cloud.com/api/v1/estate/system/property/actions/updateProperty"
	deleteUrl = "https://api2.hik-cloud.com/api/v1/estate/system/property/"
)

type Property struct {
	Http *http.Http
}

func NewProperty(http *http.Http) *Property {
	p := new(Property)
	p.Http = http
	return p
}

type CreateData struct {
	UserName string `c:"userName, omitempty" json:"user_name"` //账号
	PhoneNumber string `c:"phoneNumber" json:"phone_number"` //联系电话
	CommunityIds string `c:"communityIds" json:"community_ids"` //管辖社区
	RoleKeys string `c:"roleKeys" json:"role_keys"` //角色
}

func (p *Property) Create(data *CreateData) ([]byte,error) {
	body, err := p.Http.Post(createUrl, gconv.Map(data))
	return body, err
}
