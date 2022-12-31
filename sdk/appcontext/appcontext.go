package appcontext

import (
	"context"
	"time"
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

func SetRequestId(ctx context.Context, rid string) context.Context {
	return context.WithValue(ctx, requestId, rid)
}

func GetRequestId(ctx context.Context) string {
	rid, ok := ctx.Value(requestId).(string)
	if !ok {
		return ""
	}

	return rid
}

func SetServiceVersion(ctx context.Context, version string) context.Context {
	return context.WithValue(ctx, serviceVersion, version)
}

func GetServiceVersion(ctx context.Context) string {
	version, ok := ctx.Value(serviceVersion).(string)
	if !ok {
		return ""
	}

	return version
}

func SetUserAgent(ctx context.Context, ua string) context.Context {
	return context.WithValue(ctx, userAgent, ua)
}

func GetUserAgent(ctx context.Context) string {
	ua, ok := ctx.Value(userAgent).(string)
	if !ok {
		return ""
	}

	return ua
}

func SetRequestStartTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, requestStartTime, t)
}

func GetRequestStartTime(ctx context.Context) time.Time {
	t, ok := ctx.Value(requestStartTime).(time.Time)
	if !ok {
		return time.Time{}
	}

	return t
}

func SetDeviceType(ctx context.Context, platform string) context.Context {
	return context.WithValue(ctx, deviceType, platform)
}

func GetDeviceType(ctx context.Context) string {
	platform, ok := ctx.Value(deviceType).(string)
	if !ok {
		return "web"
	}

	return platform
}
