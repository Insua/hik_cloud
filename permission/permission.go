package permission

import (
	"github.com/Insua/hik_cloud/http"
	"github.com/gogf/gf/util/gconv"
)

const (
	personDeviceListUrl = "https://api2.hik-cloud.com/api/v1/estate/entranceGuard/remoteControl/actions/deviceList"
)

type PersonType uint8

const (
	Property PersonType = iota //[0]物业
	Resident //[1]住户
)

type Permission struct {
	Http *http.Http
}

func NewPermission(http *http.Http) *Permission {
	p := new(Permission)
	p.Http = http
	return p
}

type PersonDeviceData struct {
	PersonId string `c:"personId" json:"person_id"` //人员ID
	CommunityId string `c:"communityId" json:"community_id"` //社区ID
	PersonType PersonType `c:"personType" json:"person_type"` //人员类型
}

func (p *Permission) PersonDevice(data *PersonDeviceData) ([]byte, error)  {
	body, err := p.Http.Get(personDeviceListUrl, gconv.Map(data))
	return body, err
}
