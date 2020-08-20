package permission

import "github.com/gogf/gf/util/gconv"

const (
	gateContolUrl = "https://api2.hik-cloud.com/api/v1/estate/entranceGuard/remoteControl/actions/gateControl"
)

type Command string

const (
	Open Command = "open"
	Close Command = "close"
	AlwaysOpen Command = "alwaysOpen"
	LlwaysClose Command = "alwaysClose"
)

type GateControlData struct {
	PersonId string `c:"personId" json:"person_id"` //人员ID
	DeviceId string `c:"deviceId" json:"device_id"` //设备ID
	Command Command `c:"command" json:"command"` //控制类型
}

func (p *Permission)GateControl(data *GateControlData) ([]byte,error) {
	body, err := p.Http.Post(gateContolUrl, gconv.Map(data))
	return body, err
}
