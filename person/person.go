package person

import (
	"github.com/Insua/hik_cloud/http"

	"github.com/gogf/gf/util/gconv"
)

const (
	createUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person"
)

type CreateData struct {
	UnionId          string `c:"unionId, omitempty" json:"union_id"`                   //关联ID,保留字段
	PersonName       string `c:"personName" json:"person_name"`                        //姓名
	Gender           int    `c:"gender, omitempty" json:"gender"`                      //性别 [-1]无 [0]女 [1]男
	CredentialType   int    `c:"credentialType" json:"credential_type"`                //证件类型 [1]身份证 [2护照 [3]其他
	CredentialNumber string `c:"credentialNumber" json:"credential_number"`            //证件号码
	Mobile           string `c:"mobile" json:"mobile"`                                 //手机号
	Birthday         string `c:"birthday, omitempty" json:"birthday"`                  //出生日期
	PersonRemark     string `c:"personRemark, omitempty" json:"person_remark"`         //备注
	Nation           string `c:"nation, omitempty" json:"nation"`                      //民族
	EducationalLevel int    `c:"educationalLevel, omitempty" json:"educational_level"` //文化程度 [1]中专 [2]大专 [3]本科 [4]硕士 [5]博士 [6]其他
	WorkUnit         string `c:"workUnit, omitempty" json:"work_unit"`                 //工作单位
	Position         string `c:"position, omitempty" json:"position"`                  //职务
	Religion         int    `c:"religion, omitempty" json:"religion"`                  //宗教信仰 [1]基督教 [2]伊斯兰教 [3]佛教
	EnglishName      string `c:"englishName, omitempty" json:"english_name"`           //英文名称
	Email            string `c:"email, omitempty" json:"email"`                        //邮箱
	AddressDetail    string `c:"addressDetail, omitempty" json:"address_detail"`       //详细地址
	ProvinceCode     string `c:"provinceCode, omitempty" json:"province_code"`         //现户籍地/省编码
	CityCode         string `c:"cityCode, omitempty" json:"city_code"`                 //现户籍地/市编码
	CountyCode       string `c:"countyCode, omitempty" json:"county_code"`             //现户籍地/区编码
}

type CommunityRel struct {
	CommunityId  string `c:"communityId" json:"community_id"`               //社区ID
	BuildingId   string `c:"buildingId" json:"building_id"`                 //楼栋ID
	UnitId       string `c:"unitId" json:"unit_id"`                         //单元ID
	RoomId       string `c:"roomId" json:"room_id"`                         //户室ID
	IdentityType int    `c:"identityType" json:"identity_type"`             //身份类型 [1]业主 [2]租客 [3]家属
	CheckInDate  string `c:"checkInDate" json:"check_in_date"`              //入住时间
	CheckOutDate string `c:"checkOutDate, omitempty" json:"check_out_date"` //离开时间
}

type Person struct {
	Http *http.Http
}

func NewPerson(http *http.Http) *Person {
	p := new(Person)
	p.Http = http
	return p
}

func (p *Person) Create(data *CreateData, rels []CommunityRel) ([]byte, error) {
	mapData := gconv.Map(data)
	if len(rels) > 0 {
		mapData["personCommunityRels"] = gconv.SliceMap(rels)
	}

	body, err := p.Http.Post(createUrl, mapData)
	return body, err
}
