package http

import (
	"github.com/Insua/hik_cloud/context"
	netUrl "net/url"
	"time"

	"github.com/gogf/gf/util/gconv"

	"github.com/gogf/gf/frame/g"
)

type Http struct {
	*context.Context
}

func NewHttp(context *context.Context) *Http {
	h := new(Http)
	h.Context = context
	return h
}

func (h *Http) Get(url string, data map[string]interface{}) ([]byte, error) {
	token, tokenErr := h.GetAccessToken()
	if tokenErr != nil {
		token, tokenErr = h.ForceUpdateAccessToken()
		if tokenErr != nil {
			return nil, tokenErr
		}
	}

	code, body, resErr := getData(token, url, data)
	if resErr != nil {
		return body, resErr
	}

	if code == 401 {
		token, tokenErr = h.ForceUpdateAccessToken()
		if tokenErr != nil {
			return nil, tokenErr
		}

		_, body, resErr = getData(token, url, data)
		return body, resErr
	}

	return body, resErr
}

func getData(token, url string, data map[string]interface{}) (int, []byte, error) {
	headers := make(map[string]string)
	headers["Authorization"] = "bearer " + token

	u, urlErr := netUrl.Parse(url)

	if urlErr != nil {
		return 0, nil, urlErr
	}

	q := u.Query()
	for k, v := range data {
		q.Set(k, gconv.String(v))
	}
	u.RawQuery = q.Encode()

	response, resErr := g.Client().
		Header(headers).Timeout(60 * time.Second).
		ContentJson().
		Get(u.String())

	if resErr != nil {
		return 0, nil, resErr
	}

	return response.StatusCode, response.ReadAll(), nil
}

func (h *Http) Post(url string, data map[string]interface{}) ([]byte, error) {
	token, tokenErr := h.GetAccessToken()
	if tokenErr != nil {
		token, tokenErr = h.ForceUpdateAccessToken()
		if tokenErr != nil {
			return nil, tokenErr
		}
	}

	code, body, resErr := postData(token, url, data)
	if resErr != nil {
		return body, resErr
	}

	if code == 401 {
		token, tokenErr = h.ForceUpdateAccessToken()
		if tokenErr != nil {
			return nil, tokenErr
		}

		_, body, resErr = postData(token, url, data)
		return body, resErr
	}

	return body, resErr
}

func postData(token, url string, data map[string]interface{}) (int, []byte, error) {
	headers := make(map[string]string)
	headers["Authorization"] = "bearer " + token

	response, resErr := g.Client().
		Header(headers).Timeout(60*time.Second).
		ContentJson().
		Post(url, data)

	if resErr != nil {
		return 0, nil, resErr
	}

	return response.StatusCode, response.ReadAll(), nil
}

func (h *Http) Delete(url string, data map[string]interface{}) ([]byte, error) {
	token, tokenErr := h.GetAccessToken()
	if tokenErr != nil {
		token, tokenErr = h.ForceUpdateAccessToken()
		if tokenErr != nil {
			return nil, tokenErr
		}
	}

	code, body, resErr := deleteData(token, url, data)
	if resErr != nil {
		return body, resErr
	}

	if code == 401 {
		token, tokenErr = h.ForceUpdateAccessToken()
		if tokenErr != nil {
			return nil, tokenErr
		}

		_, body, resErr = deleteData(token, url, data)
		return body, resErr
	}

	return body, resErr
}

func deleteData(token, url string, data map[string]interface{}) (int, []byte, error) {
	headers := make(map[string]string)
	headers["Authorization"] = "bearer " + token

	response, resErr := g.Client().
		Header(headers).Timeout(60*time.Second).
		ContentJson().
		Delete(url, data)

	if resErr != nil {
		return 0, nil, resErr
	}

	return response.StatusCode, response.ReadAll(), nil
}
