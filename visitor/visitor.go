package visitor

import (
	"github.com/Insua/hik_cloud/http"
	"github.com/gogf/gf/util/gconv"
)

const (
	reserveUrl = "https://api2.hik-cloud.com/api/v1/estate/visitors"
	deleteUrl = "https://api2.hik-cloud.com/api/v1/estate/visitors/"
)

type PersonType uint8

type Gender uint8

const (
	Property PersonType = iota //[0]物业
	Resident //[1]住户
)

const (
	Female Gender = iota //[0]女
	Male //[1]男
)

type Visitor struct {
	Http *http.Http
}

type ReserveData struct {
	PersonId string `c:"personId" json:"person_id"` //人员ID
	PersonType PersonType `c:"personType" json:"person_type"` //人员类型
	VisitorName string `c:"visitorName" json:"visitor_name"` //访客姓名
	Gender Gender `c:"gender" json:"gender"` //访客性别
	Phone string `c:"phone" json:"phone"` //访客手机号
	RoomId string `c:"roomId" json:"room_id"` //访问的户室ID
	VisitorEffectiveTime string `c:"visitorEffectiveTime" json:"visitor_effective_time"` //来访开始时间（UTC+08:00）
	VisitorLeaveTime string `c:"visitorLeaveTime" json:"visitor_leave_time"` //离开时间（UTC+08:00）
	VisitReason string `c:"visitReason, omitempty" json:"visit_reason"` //访问原由
	DeviceIds string `c:"deviceIds, omitempty" json:"device_ids"` //需要下发访客权限的设备列表
	OpenTimes uint8 `c:"openTimes, omitempty" json:"open_times"` //二维码开门次数
}

func NewVisitor(http *http.Http) *Visitor{
	v := new(Visitor)
	v.Http = http
	return v
}

func (v *Visitor) Reserve (data *ReserveData) ([]byte,error) {
	body, err := v.Http.Post(reserveUrl, gconv.Map(data))
	return body, err
}
