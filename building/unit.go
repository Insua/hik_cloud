package building

import (
	"github.com/Insua/hik_cloud/context"

	"github.com/gogf/gf/util/gconv"
)

const (
	createUnitUrl       = "https://api2.hik-cloud.com/api/v1/estate/system/units"
	deleteUnitUrl       = "https://api2.hik-cloud.com/api/v1/estate/system/units/"
	listBuildingUnitUrl = "https://api2.hik-cloud.com/api/v1/estate/system/units/actions/buildingUnitList"
)

type CreateUnitData struct {
	UnionId    string `c:"unionId, omitempty" json:"union_id"`   //关联ID,保留字段
	BuildingId string `c:"buildingId" json:"building_id"`        //楼栋ID
	UnitName   string `c:"unitName, omitempty" json:"unit_name"` //单元名称
	UnitNumber string `c:"unitNumber" json:"unit_number"`        //单元号（1-30）
}

type ListUnitData struct {
	UnitId string `c:"unitId" json:"unit_id"` //单元ID
	context.ListData
}

func (b *Building) CreateUnit(data *CreateUnitData) ([]byte, error) {
	body, err := b.Http.Post(createUnitUrl, gconv.Map(data))
	return body, err
}

func (b *Building) DeleteUnit(unitId string) ([]byte, error) {
	body, err := b.Http.Delete(deleteUnitUrl+unitId, gconv.Map(nil))
	return body, err
}

func (b *Building) ListBuildingUnit(data *ListBuildingData) ([]byte, error) {
	body, err := b.Http.Post(listBuildingUnitUrl, gconv.MapDeep(data))
	return body, err
}
