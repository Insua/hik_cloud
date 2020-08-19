package credential

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/encoding/gjson"

	"github.com/gogf/gf/frame/g"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/crypto/gsha1"

	"github.com/gogf/gf/database/gredis"
)

const (
	accessTokenURL = "https://api2.hik-cloud.com/oauth/token"
	CachePrefix    = "hik_cloud_access_token_"
)

type AccessTokenHandle interface {
	GetAccessToken() (accessToken string, err error)
	ForceUpdateAccessToken() (accessToken string, err error)
}

type AccessToken struct {
	clientId        string
	clientSecret    string
	grantType       string
	scope           string
	accessTokenLock *sync.Mutex
	redis           *gredis.Redis
}

func NewAccessToken(clientId, clientSecret string, redis *gredis.Redis) AccessTokenHandle {
	return &AccessToken{
		clientId:        clientId,
		clientSecret:    clientSecret,
		accessTokenLock: new(sync.Mutex),
		redis:           redis,
	}
}

func (ak *AccessToken) GetAccessToken() (accessToken string, err error) {

	cacheValue, getRedisErr := ak.redis.DoVar("GET", ak.cacheKey())
	if getRedisErr == nil && len(gconv.String(cacheValue)) > 0 {
		accessToken = gconv.String(cacheValue)
		return
	}

	return ak.ForceUpdateAccessToken()
}

func (ak *AccessToken) ForceUpdateAccessToken() (accessToken string, err error) {
	ak.accessTokenLock.Lock()
	defer ak.accessTokenLock.Unlock()

	accessTokenCacheKey := ak.cacheKey()

	response, resErr := g.Client().Timeout(60*time.Second).ContentType("application/x-www-form-urlencoded").
		Post(accessTokenURL, g.Map{
			"client_id":     ak.clientId,
			"client_secret": ak.clientSecret,
			"grant_type":    "client_credentials",
			"scope":         "app",
		})
	if resErr != nil {
		err = resErr
		return
	}
	defer response.Close()
	if response == nil {
		err = errors.New("empty access token return")
		return
	}

	json := gjson.New(response.ReadAll())

	jsonAccessToken := json.GetString("access_token")
	jsonExpiresIn := json.GetInt("expires_in")
	if len(jsonAccessToken) == 0 || jsonExpiresIn == 0 {
		err = errors.New("wrong access token or expires in get in response")
		return
	}

	accessToken = jsonAccessToken
	_, setErr := ak.redis.DoVar("SET", accessTokenCacheKey, accessToken)
	_, expireErr := ak.redis.DoVar("EXPIRE", accessTokenCacheKey, jsonExpiresIn)
	if setErr != nil || expireErr != nil {
		err = errors.New("cache access token error")
	}
	return
}

func (ak *AccessToken) cacheKey() string {
	return fmt.Sprintf("%s%s", CachePrefix, gsha1.Encrypt(ak.clientId))
}
