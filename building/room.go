package building

import (
	"github.com/Insua/hik_cloud/community"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

const (
	createRoomUrl        = "https://api2.hik-cloud.com/api/v1/estate/system/rooms"
	deleteRoomUrl        = "https://api2.hik-cloud.com/api/v1/estate/system/rooms/"
	queryRoomByNumberUrl = "https://api2.hik-cloud.com/api/v1/estate/system/rooms/actions/info"
	queryRoomByIdUrl     = "https://api2.hik-cloud.com/api/v1/estate/system/rooms/actions/infoById"
	listCommunityRoomUrl = "https://api2.hik-cloud.com/api/v1/estate/system/rooms/actions/communityRoomList"
	listUnitRoomUrl      = "https://api2.hik-cloud.com/api/v1/estate/system/rooms/actions/unitRoomList"
)

type CreateRoomData struct {
	UnionId     string `c:"unionId, omitempty" json:"union_id"`   //关联ID,保留字段
	UnitId      string `c:"unitId" json:"unit_id"`                //单元ID
	FloorNumber string `c:"floorNumber" json:"floor_number"`      //所在楼层
	RoomNumber  string `c:"roomNumber" json:"room_number"`        //两位户室编号（不带楼层，如7层01室，传值为01，传值范围01-50）
	RoomName    string `c:"roomName, omitempty" json:"room_name"` //户室名称
}

type QueryRoomByNumberData struct {
	CommunityId    string `c:"communityId" json:"community_id"`       //社区id
	BuildingNumber string `c:"buildingNumber" json:"building_number"` //所属楼栋编号（1-999之间的整数）
	UnitNumber     string `c:"unitNumber" json:"unit_number"`         //所属单元编号（1-30之间的整数）
	RoomNumber     string `c:"roomNumber" json:"room_number"`         //户室编号
}

func (b *Building) CreateRoom(data *CreateRoomData) ([]byte, error) {
	body, err := b.Http.Post(createRoomUrl, gconv.Map(data))
	return body, err
}

func (b *Building) DeleteRoom(roomId string) ([]byte, error) {
	body, err := b.Http.Delete(deleteRoomUrl+roomId, gconv.Map(nil))
	return body, err
}

func (b *Building) QueryRoomByNumber(data *QueryRoomByNumberData) ([]byte, error) {
	body, err := b.Http.Get(queryRoomByNumberUrl, gconv.Map(data))
	return body, err
}

func (b *Building) QueryRoomById(roomId string) ([]byte, error) {
	body, err := b.Http.Get(queryRoomByIdUrl, g.Map{
		"roomId": roomId,
	})
	return body, err
}

func (b *Building) ListCommunityRoom(data *community.ListCommunityData) ([]byte, error) {
	body, err := b.Http.Post(listCommunityRoomUrl, gconv.MapDeep(data))
	return body, err
}

func (b *Building) ListUnitRoom(data *ListUnitData) ([]byte, error) {
	body, err := b.Http.Post(listUnitRoomUrl, gconv.MapDeep(data))
	return body, err
}
