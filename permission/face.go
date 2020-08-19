package permission

import "github.com/gogf/gf/util/gconv"

const (
	faceIssuedUrl = "https://api2.hik-cloud.com/api/v1/estate/entranceGuard/permissions/actions/faceIssued"
	deleteFaceUrl = "https://api2.hik-cloud.com/api/v1/estate/entranceGuard/permissions/actions/deleteFaceIssued"
)

type FaceIssuedData struct {
	PersonId   string `c:"personId" json:"person_id"`            //需要下发的人员ID
	PersonType int    `c:"personType" json:"person_type"`        //人员类型 [0]物业 [1]住户
	DeviceId   string `c:"deviceId, omitempty" json:"device_id"` //设备ID
	FaceUrl    string `c:"faceUrl" json:"face_url"`              //人脸图片URL
}

func (p *Permission) FaceIssued(data *FaceIssuedData) ([]byte, error) {
	body, err := p.Http.Post(faceIssuedUrl, gconv.Map(data))
	return body, err
}
