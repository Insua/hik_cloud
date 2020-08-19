package permission

import "github.com/gogf/gf/util/gconv"

const (
	authorityIssuedUrl = "https://api2.hik-cloud.com/api/v1/estate/entranceGuard/permissions/actions/authorityIssued"
)

type AuthorityIssuedData struct {
	CommunityId string `c:"communityId" json:"community_id"` //社区ID
	PersonId    string `c:"personId" json:"person_id"`       //人员ID
	PersonType  uint8  `c:"personType" json:"person_type"`   //人员类型 [0]物业 [1]住户
	DeviceId    string `c:"deviceId" json:"device_id"`       //设备ID
}

func (p *Permission) AuthorityIssued(data *AuthorityIssuedData) ([]byte, error) {
	body, err := p.Http.Post(authorityIssuedUrl, gconv.Map(data))
	return body, err
}
