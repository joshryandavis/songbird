package starling

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"

	"golang.org/x/time/rate"
)

type Client struct {
	Token string
}

func New(token string) *Client {
	ret := new(Client)
	ret.Token = token
	return ret
}

// RateLimit because PAT has a rate limit of 5 requests per second
const RateLimit = 5

var httpClient = &http.Client{}
var limiter = rate.NewLimiter(rate.Limit(RateLimit), 1)

func (c *Client) Request(method string, url string, body string) (val []byte, err error) {
	var ret []byte
	log.WithFields(log.Fields{
		"method": method,
		"url":    url,
		"body":   body,
	}).Info("request")
	if err := limiter.Wait(context.Background()); err != nil {
		log.Panic("rate limit error:", err)
		os.Exit(1)
		return ret, err
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Panic("request error:", err)
		return ret, err
	}
	setRequestHeaders(req)
	setAccessToken(req, c.Token)
	if body != "" {
		setRequestBody(req, body)
	}
	res, err := httpClient.Do(req)
	if err != nil {
		log.Panic("http starling error:", err)
		return ret, err
	}
	err = verifyResponse(res)
	if err != nil {
		log.Panic("http response err", err)
		return ret, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Panic("error closing body:", err)
			return
		}
	}(res.Body)
	buf, err := io.ReadAll(res.Body)
	if err != nil {
		log.Panic("error reading body:", err)
		return ret, err
	}
	ret = append(ret, buf...)
	return ret, nil
}

func verifyResponse(res *http.Response) error {
	status := res.StatusCode
	log.WithFields(log.Fields{
		"status": status,
	}).Info("response status")
	if status < 200 || status > 299 {
		// output the response body for context
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				log.Panic("error closing body:", err)
			}
		}(res.Body)
		buf, err := io.ReadAll(res.Body)
		if err != nil {
			log.Panic("error reading body:", err)
			return err
		}
		errResponse := ErrorResponse{}
		err = json.Unmarshal(buf, &errResponse)
		if err != nil {
			log.Panic("error parsing response body:", err)
			return err
		}
		log.WithFields(log.Fields{
			"error": errResponse,
		}).Error("error response")
		return err
	}
	return nil
}

func setRequestBody(req *http.Request, body string) {
	log.Println("request body:", body)
	req.Body = io.NopCloser(strings.NewReader(body))
}

func setRequestHeaders(req *http.Request) {
	req.Header.Set("Accept", "application/json")
	if req.Method == "POST" || req.Method == "PUT" {
		req.Header.Set("Content-Type", "application/json")
	}
}

func setAccessToken(req *http.Request, accessToken string) {
	req.Header.Set("Authorization", "Bearer "+accessToken)
}
