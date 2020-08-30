package door

import "github.com/gogf/gf/frame/g"

func EventCode() []g.Map {
	data := make([]g.Map, 10)
	data[0] = g.Map{"code": 10101, "type": "刷卡开锁"}
	data[1] = g.Map{"code": 10104, "type": "指纹开锁"}
	data[2] = g.Map{"code": 10114, "type": "人脸开锁"}
	data[3] = g.Map{"code": 10118, "type": "人证开锁"}
	data[4] = g.Map{"code": 10119, "type": "蓝牙开锁"}
	data[5] = g.Map{"code": 10120, "type": "密码开锁"}
	data[6] = g.Map{"code": 10122, "type": "二维码开锁"}
	data[7] = g.Map{"code": 10124, "type": "远程开门"}
	data[8] = g.Map{"code": 10125, "type": "密码开锁"}
	data[9] = g.Map{"code": 10126, "type": "人脸开锁失败"}
	return data
}
