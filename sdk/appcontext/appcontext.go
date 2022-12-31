package appcontext

import (
	"time"

	"github.com/gin-gonic/gin"
)

type contextKey string

const (
	// Context keys
	requestId        contextKey = "RequestId"
	serviceVersion   contextKey = "ServiceVersion"
	userAgent        contextKey = "UserAgent"
	requestStartTime contextKey = "RequestStartTime"
	deviceType       contextKey = "DeviceType"

	// Header keys
	HeaderRequestId    = "x-request-id"
	HeaderCacheControl = "cache-control"
	HeaderUserAgent    = "user-agent"
	HeaderDeviceType   = "x-device-type"
)

func SetRequestId(ctx *gin.Context, rid string) {
	ctx.Set(string(requestId), rid)
}

func GetRequestId(ctx *gin.Context) string {
	rid, ok := ctx.Get(string(requestId))
	if !ok {
		return ""
	}
	return rid.(string)
}

func SetServiceVersion(ctx *gin.Context, version string) {
	ctx.Set(string(serviceVersion), version)
}

func GetServiceVersion(ctx *gin.Context) string {
	version, ok := ctx.Get(string(serviceVersion))
	if !ok {
		return ""
	}
	return version.(string)
}

func SetUserAgent(ctx *gin.Context, ua string) {
	ctx.Set(string(userAgent), ua)
}

func GetUserAgent(ctx *gin.Context) string {
	ua, ok := ctx.Get(string(userAgent))
	if !ok {
		return ""
	}
	return ua.(string)
}

func SetRequestStartTime(ctx *gin.Context, t time.Time) {
	ctx.Set(string(requestStartTime), t)
}

func GetRequestStartTime(ctx *gin.Context) time.Time {
	t, ok := ctx.Get(string(requestStartTime))
	if !ok {
		return time.Time{}
	}
	return t.(time.Time)
}

func SetDeviceType(ctx *gin.Context, platform string) {
	ctx.Set(string(deviceType), platform)
}

func GetDeviceType(ctx *gin.Context) string {
	platform, ok := ctx.Get(string(deviceType))
	if !ok {
		return "web"
	}
	return platform.(string)
}
