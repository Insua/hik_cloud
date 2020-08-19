package device

import (
	"github.com/Insua/hik_cloud/http"

	"github.com/gogf/gf/util/gconv"
)

const (
	createUrl = "https://api2.hik-cloud.com/api/v1/estate/devices"
)

type Device struct {
	Http *http.Http
}

func NewDevice(http *http.Http) *Device {
	d := new(Device)
	d.Http = http
	return d
}

type CreateData struct {
	UnionId      string `c:"unionId, omitempty" json:"union_id"`       //关联ID,保留字段
	DeviceSerial string `c:"deviceSerial " json:"device_serial"`       //设备序列号
	ValidateCode string `c:"validateCode" json:"validate_code"`        //设备验证码
	DeviceName   string `c:"deviceName" json:"device_name"`            //设备名称
	CommunityId  string `c:"communityId" json:"community_id"`          //社区ID
	BuildingId   string `c:"buildingId, omitempty" json:"building_id"` //楼栋ID
	UnitId       string `c:"unitId, omitempty" json:"unit_id"`         //单元ID
}

func (d *Device) Create(data *CreateData) ([]byte, error) {
	body, err := d.Http.Post(createUrl, gconv.Map(data))
	return body, err
}
