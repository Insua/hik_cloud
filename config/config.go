package config

import "github.com/gogf/gf/database/gredis"

type Config struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Redis        *gredis.Redis
}
