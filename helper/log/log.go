package loggingHelper

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	conversionHelper "belajar/helper/conversion"
	jwtAuth "belajar/package/auth/jwt"
	"belajar/package/config"

	log "github.com/sirupsen/logrus"
)

// Return JSON Logging
func Addlog(r *http.Request, level string, data interface{}) {
	getParameters, _ := url.ParseQuery(r.URL.RawQuery)
	param := make(map[string]interface{}, len(getParameters))
	for k, v := range getParameters {
		param[k] = v[0]
	}
	if len(param) == 0 {
		param = nil
	}

	ctx := r.Context()
	checkBody := ctx.Value(config.ContextKey("body"))
	body := make(map[string]interface{})
	if checkBody != nil {
		getBody := ctx.Value(config.ContextKey("body")).([]byte)
		_ = json.Unmarshal(getBody, &body)
		if len(body) == 0 {
			body = nil
		}
	}

	headers := make(map[string]interface{})
	exceptionHeaders := map[string]bool{
		"Cache-Control":   true,
		"Postman-Token":   true,
		"Content-Length":  true,
		"Host":            true,
		"User-Agent":      true,
		"Accept-Encoding": true,
		"Accept":          true,
		"Connection":      true,
	}
	for k, v := range r.Header {
		if !exceptionHeaders[k] {
			headers[k] = v[0]
		}
	}
	if len(headers) == 0 {
		headers = nil
	}

	userInfo, ok := ctx.Value(config.ContextKey("userInfo")).(*jwtAuth.JwtData)

	var UserID int64
	if ok {
		UserID, _ = conversionHelper.StrToInt64(userInfo.UserID)
	}

	checkStartTime := ctx.Value(config.ContextKey("startTime"))
	responseTime := time.Duration(0)
	if checkStartTime != nil {
		getStartTime := ctx.Value(config.ContextKey("startTime")).(time.Time)
		responseTime = time.Duration(time.Since(getStartTime).Milliseconds())
	}

	req := &ResponseLog{
		HostName:      r.Host,
		Path:          r.URL.Path,
		RequestMethod: r.Method,
		Params:        param,
		Body:          body,
		Headers:       headers,
		UserAgent:     r.UserAgent(),
		ResponseTime:  responseTime,
		Response:      data,
		UserID:        UserID,
	}

	log.SetFormatter(&log.JSONFormatter{})
	if level == "ERROR" {
		log.WithFields(log.Fields{
			"attributes": &req,
		}).Error(data)
	} else {
		log.WithFields(log.Fields{
			"attributes": req,
		}).Info("Success")
	}
}
