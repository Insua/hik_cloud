package consumer

import (
	"errors"
	"fmt"
	"github.com/Insua/hik_cloud/http"
	"github.com/gogf/gf/crypto/gsha1"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
)

const (
	CacheConstomerIdPrefix ="hik_cloud_customer_id_"
)

const (
	createUrl = "https://api2.hik-cloud.com/api/v1/mq/consumer/group1"
	consumeUrl = "https://api2.hik-cloud.com/api/v1/mq/consumer/messages"
)

type Consumer struct {
	Http *http.Http
}

type ConsumeData struct {
	ConsumerId string `c:"consumerId" json:"consumer_id"` //消费者ID
	AutoCommit bool `c:"autoCommit, omitempty" json:"auto_commit"` //是否自动提交偏移量，默认false
}

func NewConsumer(http *http.Http) *Consumer {
	c := new(Consumer)
	c.Http = http
	return c
}

func (c *Consumer) Create() ([]byte, error) {
	body, err := c.Http.PostFormUrl(createUrl, g.Map{
		"consumerName": "group1",
	})
	return body, err
}

func (c *Consumer) Consume(data *ConsumeData) ([]byte,error)  {
	body, err := c.Http.PostFormUrl(consumeUrl, gconv.Map(data))
	return body, err
}

func (c *Consumer) Messages() ([]byte, error)  {
	customerId := ""
	cacheValue, getRedisErr := c.Http.Redis.DoVar("GET", c.cacheKey())
	if getRedisErr == nil && len(gconv.String(cacheValue)) > 0 {
		customerId = gconv.String(cacheValue)
	} else {
		create, errCreate := c.Create()
		if errCreate != nil {
			return nil, errCreate
		}
		jsonCreate := gjson.New(create)
		if jsonId := jsonCreate.Get("data.consumerId"); jsonId != nil {
			customerId = gconv.String(jsonId)
			if len(customerId) == 32 {
				_, setErr := c.Http.Redis.DoVar("SET", c.cacheKey(), customerId)
				if setErr != nil {
					return nil, setErr
				}
			}
		}
	}
	if len(customerId) != 32 {
		return nil, errors.New("wrong customer id")
	}

	data := ConsumeData{
		ConsumerId: customerId,
		AutoCommit: true,
	}

	consume, consumeErr :=  c.Consume(&data)
	if consumeErr != nil {
		return nil, consumeErr
	}

	jsonConsume := gjson.New(consume)
	if codeConsume := jsonConsume.GetInt("code"); codeConsume != 200 {
		if _, removeErr := c.Http.Redis.DoVar("DEL", c.cacheKey()); removeErr != nil {
			return nil, removeErr
		}
	}
	return consume, consumeErr
}

func (c *Consumer) cacheKey() string {
	return fmt.Sprintf("%s%s", CacheConstomerIdPrefix, gsha1.Encrypt(c.Http.ClientId))
}
