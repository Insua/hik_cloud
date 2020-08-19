package building

import (
	"github.com/Insua/hik_cloud/community"
	"github.com/Insua/hik_cloud/context"
	"github.com/Insua/hik_cloud/http"

	"github.com/gogf/gf/util/gconv"
)

const (
	createBuildingUrl        = "https://api2.hik-cloud.com/api/v1/estate/system/buildings"
	deleteBuildingUrl        = "https://api2.hik-cloud.com/api/v1/estate/system/buildings/"
	listCommunityBuildingUrl = "https://api2.hik-cloud.com/api/v1/estate/system/buildings/actions/communityBuildingList"
)

type Building struct {
	Http *http.Http
}

type CreateBuildingData struct {
	UnionId          string `c:"unionId, omitempty" json:"union_id"`               //关联ID,保留字段
	CommunityId      string `c:"communityId" json:"community_id"`                  //社区ID
	BuildingName     string `c:"buildingName" json:"building_name"`                //楼栋名称
	BuildingNumber   string `c:"buildingNumber" json:"building_number"`            //楼栋编号（1-999之间的整数）
	FloorUpCount     string `c:"floorUpCount" json:"floor_up_count"`               //地上楼层数（最多支持100层）
	FloorDownCount   string `c:"floorDownCount omitempty" json:"floor_down_count"` //地下楼层数（最多支持3层）
	FloorFamilyCount string `c:"floorFamilyCount" json:"floor_family_count"`       //每层户数（每层最多支持50户）
	BuildingUnitSize string `c:"buildingUnitSize" json:"building_unit_size"`       //楼栋单元数量（每栋最多支持30单元）
	BuildingRemark   string `c:"buildingRemark omitempty" json:"building_remark"`  //备注
}

type ListBuildingData struct {
	BuildingId string `c:"buildingId" json:"building_id"` //楼栋ID
	context.ListData
}

func NewBuilding(http *http.Http) *Building {
	b := new(Building)
	b.Http = http
	return b
}

func (b *Building) CreateBuilding(data *CreateBuildingData) ([]byte, error) {
	body, err := b.Http.Post(createBuildingUrl, gconv.Map(data))
	return body, err
}

func (b *Building) DeleteBuilding(buildingId string) ([]byte, error) {
	body, err := b.Http.Delete(deleteBuildingUrl+buildingId, gconv.Map(nil))
	return body, err
}

func (b *Building) ListCommunityBuilding(data *community.ListCommunityData) ([]byte, error) {
	body, err := b.Http.Post(listCommunityBuildingUrl, gconv.MapDeep(data))
	return body, err
}
