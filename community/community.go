package community

import (
	"github.com/Insua/hik_cloud/context"
	"github.com/Insua/hik_cloud/http"

	"github.com/gogf/gf/util/gconv"
)

const (
	createUrl = "https://api2.hik-cloud.com/api/v1/estate/system/communities"
	deleteUrl = "https://api2.hik-cloud.com/api/v1/estate/system/communities/"
	updateUrl = "https://api2.hik-cloud.com/api/v1/estate/system/communities/actions/updateCommunity"
	listUrl   = "https://api2.hik-cloud.com/api/v1/estate/system/communities/actions/list"
)

type Community struct {
	Http *http.Http
}

type CreateData struct {
	UnionId              string `c:"unionId, omitempty" json:"union_id"`                            //关联ID,保留字段
	CommunityName        string `c:"communityName" json:"community_name"`                           //社区名称
	ProvinceCode         string `c:"provinceCode" json:"province_code"`                             //省代号
	CityCode             string `c:"cityCode, omitempty" json:"city_code"`                          //市代号
	CountryCode          string `c:"countryCode, omitempty" json:"country_code"`                    //区代号
	AddressDetail        string `c:"addressDetail" json:"address_detail"`                           //街道详细地址
	CommunitySquareMeter string `c:"communitySquareMeter, omitempty" json:"community_square_meter"` //社区面积(万㎡) 最多8位整数, 2位小数
	Longitude            string `c:"longitude, omitempty" json:"longitude"`                         //经度坐标值
	Latitude             string `c:"latitude, omitempty" json:"latitude"`                           //维度坐标值
	ChargePersonId       string `c:"chargePersonId, omitempty" json:"charge_person_id"`             //负责人ID（该社区的物业负责人）
	PhoneNumber          string `c:"phoneNumber, omitempty" json:"phone_number"`                    //联系方式
	CommunityRemark      string `c:"communityRemark, omitempty" json:"community_remark"`            //备注
}

type UpdateData struct {
	CommunityId          string `c:"communityId" json:"community_id"`                               //社区ID
	CommunityName        string `c:"communityName" json:"community_name"`                           //社区名称
	ProvinceCode         string `c:"provinceCode" json:"province_code"`                             //省代号
	CityCode             string `c:"cityCode, omitempty" json:"city_code"`                          //市代号
	CountryCode          string `c:"countryCode, omitempty" json:"country_code"`                    //区代号
	AddressDetail        string `c:"addressDetail" json:"address_detail"`                           //街道详细地址
	CommunitySquareMeter string `c:"communitySquareMeter, omitempty" json:"community_square_meter"` //社区面积(万㎡) 最多8位整数, 2位小数
	Longitude            string `c:"longitude, omitempty" json:"longitude"`                         //经度坐标值
	Latitude             string `c:"latitude, omitempty" json:"latitude"`                           //维度坐标值
	ChargePersonId       string `c:"chargePersonId, omitempty" json:"charge_person_id"`             //负责人ID（该社区的物业负责人）
	PhoneNumber          string `c:"phoneNumber, omitempty" json:"phone_number"`                    //联系方式
	CommunityRemark      string `c:"communityRemark, omitempty" json:"community_remark"`            //备注
}

type ListCommunityData struct {
	CommunityId string `c:"communityId" json:"community_id"` //社区ID
	context.ListData
}

type ListData struct {
	context.ListData
}

func NewCommunity(http *http.Http) *Community {
	c := new(Community)
	c.Http = http
	return c
}

func (c *Community) Create(data *CreateData) ([]byte, error) {
	body, err := c.Http.Post(createUrl, gconv.Map(data))
	return body, err
}

func (c *Community) Delete(communityId string) ([]byte, error) {
	body, err := c.Http.Delete(deleteUrl+communityId, gconv.Map(nil))
	return body, err
}

func (c *Community) Update(data *UpdateData) ([]byte, error) {
	body, err := c.Http.Post(updateUrl, gconv.Map(data))
	return body, err
}

func (c *Community) List(data *ListData) ([]byte, error) {
	body, err := c.Http.Get(listUrl, gconv.MapDeep(data))
	return body, err
}
