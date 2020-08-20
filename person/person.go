package person

import (
	"github.com/Insua/hik_cloud/community"
	"github.com/Insua/hik_cloud/context"
	"github.com/Insua/hik_cloud/http"
	"github.com/gogf/gf/util/gconv"
)

const (
	createUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person"
	deleteUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person/"
	updateUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person/actions/updatePerson"
	addCommunityRelationUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person/actions/addCommunityRelation"
	deleteCommunityRelationUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person/actions/deleteCommunityRelation"
	addRoomRelationUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person/actions/addRoomRelation"
	deleteRoomRelationUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person/actions/deleteRoomRelation"
	listRoomUrl = "https://api2.hik-cloud.com/api/v1/estate/system/person/actions/roomList"
	listUrl   = "https://api2.hik-cloud.com/api/v1/estate/system/person/actions/personInfoList"
)

type Gender int8

type CredentialType uint8

type EducationalLevel uint8

type Religion uint8

type IdentityType uint8

type IsAudit uint8

const (
	None   Gender = iota - 1 //[-1]无
	Female //[0]女
	Male   //[1]男
)

const (
	IdCard CredentialType = iota + 1 //[1]身份证
	PassPort //[2护照
	CredentialTypeOthers //[3]其他
)

const (
	TechnicalSecondarySchool EducationalLevel = iota + 1 //[1]中专
	Diploma //[2]大专
	Bachelor //[3]本科
	Master //[4]硕士
	Phd //[5]博士
	EducationalLevelOthers //[6]其他
)

const (
	Christian Religion = iota + 1 //[1]基督教
	Islam //[2]伊斯兰教
	Buddhism //[3]佛教
)

const (
	Property IdentityType = iota + 1 //[1]业主
	Tenat //[2]租客
	FamilyMember //[3]家属
)

const (
	DoNotNeedAudit = iota
	NeedAudit
)

type CreateData struct {
	UnionId          string `c:"unionId, omitempty" json:"union_id"`                   //关联ID,保留字段
	PersonName       string `c:"personName" json:"person_name"`                        //姓名
	Gender           Gender `c:"gender" json:"gender"`                                 //性别
	CredentialType   CredentialType `c:"credentialType" json:"credential_type"`                //证件类型
	CredentialNumber string `c:"credentialNumber" json:"credential_number"`            //证件号码
	Mobile           string `c:"mobile" json:"mobile"`                                 //手机号
	Birthday         string `c:"birthday, omitempty" json:"birthday"`                  //出生日期
	PersonRemark     string `c:"personRemark, omitempty" json:"person_remark"`         //备注
	Nation           string `c:"nation, omitempty" json:"nation"`                      //民族
	EducationalLevel EducationalLevel `c:"educationalLevel, omitempty" json:"educational_level"` //文化程度
	WorkUnit         string `c:"workUnit, omitempty" json:"work_unit"`                 //工作单位
	Position         string `c:"position, omitempty" json:"position"`                  //职务
	Religion         Religion `c:"religion, omitempty" json:"religion"`                //宗教信仰
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
	IdentityType int    `c:"identityType" json:"identity_type"`             //身份类型
	CheckInDate  string `c:"checkInDate" json:"check_in_date"`              //入住时间
	CheckOutDate string `c:"checkOutDate, omitempty" json:"check_out_date"` //离开时间
}

type UpdateData struct {
	PersonId string `c:"personId" json:"person_id"` //人员Id
	PersonName string `c:"personName" json:"person_name"` //姓名
	Gender Gender `c:"gender" json:"gender"` //性别
	CredentialType CredentialType `c:"credentialType" json:"credential_type"` //证件类型
	CredentialNumber string `c:"credentialNumber" json:"credential_number"` //证件号码
	Mobile string `c:"mobile" json:"mobile"` //手机号
	Birthday string `c:"birthday,omitempty" json:"birthday"` //出生日期
	PersonRemark string `c:"personRemark, omitempty" json:"person_remark"` //备注
	Nation string `c:"nation, omitempty" json:"nation"` //民族
	EducationalLevel EducationalLevel `c:"educationalLevel, omitempty" json:"educational_level"` //文化程度
	WorkUnit string `c:"workUnit, omitempty" json:"work_unit"` //工作单位
	Position string `c:"position, omitempty" json:"position"` //职务
	Religion Religion `c:"religion, omitempty" json:"religion"` //宗教信仰
	EnglishName string `c:"englishName, omitempty" json:"english_name"` //英文名称
	Email string `c:"email, omitempty" json:"email"` //邮箱
	AddressDetail string `c:"addressDetail, omitempty" json:"address_detail"` //详细地址
	ProvinceCode     string `c:"provinceCode, omitempty" json:"province_code"`         //现户籍地/省编码
	CityCode         string `c:"cityCode, omitempty" json:"city_code"`                 //现户籍地/市编码
	CountyCode       string `c:"countyCode, omitempty" json:"county_code"`             //现户籍地/区编码
}

type CommunityRelationData struct {
	PersonId string `c:"personId" json:"person_id"` //人员ID
	CommunityId string `c:"communityId" json:"community_id"` //社区ID
}

type AddRoomRelationData struct {
	PersonId string `c:"personId" json:"person_id"` //人员ID
	RoomId string `c:"roomId" json:"room_id"` //户室ID
	IdentityType IdentityType `c:"identityType" json:"identity_type"` //身份类型
	CheckInDate string `c:"checkInDate" json:"check_in_date"` //入住时间
	CheckOutDate string `c:"checkOutDate, omitempty" json:"check_out_date"` //离开时间
	IsAudit IsAudit `c:"isAudit, omitempty" json:"is_audit"` //是否需要审核
}

type DeleteRoomRelationData struct {
	PersonId string `c:"personId" json:"person_id"` //人员ID
	RoomId string `c:"roomId" json:"room_id"` //户室ID
}

type ListRoomData struct {
	PersonId string `c:"personId" json:"person_id"` //人员ID
	CommunityId string `c:"communityId" json:"community_id"` //社区ID
	context.ListData
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

func (p *Person) Delete(personId string) ([]byte, error) {
	body, err := p.Http.Delete(deleteUrl+personId, gconv.Map(nil))
	return body, err
}

func (p *Person) Update(data *UpdateData) ([]byte, error)  {
	body, err := p.Http.Post(updateUrl, gconv.Map(data))
	return body, err
}

func (p *Person) AddCommunityRelation(data *CommunityRelationData) ([]byte, error) {
	body, err := p.Http.Post(addCommunityRelationUrl, gconv.Map(data))
	return body, err
}

func (p *Person) DeleteCommunityRelation(data *CommunityRelationData) ([]byte, error) {
	body, err := p.Http.Post(deleteCommunityRelationUrl, gconv.Map(data))
	return body, err
}

func (p *Person) AddRoomRelation(data *AddRoomRelationData) ([]byte, error) {
	body, err := p.Http.Post(addRoomRelationUrl, gconv.Map(data))
	return body, err
}

func (p *Person) DeleteRoomRelation(data *DeleteRoomRelationData) ([]byte, error) {
	body, err := p.Http.Post(deleteRoomRelationUrl, gconv.Map(data))
	return body, err
}

func (p *Person) ListRoom (data *ListRoomData) ([]byte, error)  {
	body, err := p.Http.Get(listRoomUrl, gconv.MapDeep(data))
	return body, err
}

func (p *Person) List (data *community.ListCommunityData) ([]byte, error)  {
	body, err := p.Http.Post(listUrl, gconv.MapDeep(data))
	return body, err
}
