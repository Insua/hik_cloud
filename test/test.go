package main

import (
	"fmt"
	"github.com/Insua/hik_cloud"
	"github.com/Insua/hik_cloud/config"
	"github.com/gogf/gf/frame/g"
)

func Client() *hik_cloud.HikCloud{
	return hik_cloud.NewHikCloud(&config.Config{
		ClientId:     g.Cfg().GetString("hikCloud.clientId"),
		ClientSecret: g.Cfg().GetString("hikCloud.clientSecret"),
		Redis:        g.Redis(),
	})
}

func main() {
	client := Client()
	fmt.Println(client)
}
