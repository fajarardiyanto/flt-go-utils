package grpc

import (
	"net/http"
	"strings"
)

// https://github.com/grpc/grpc/blob/master/doc/PROTOCOL-WEB.md#protocol-differences-vs-grpc-over-http2

const GrpcContentType = "application/grpc"
const GrpcWebContentType = "application/grpc-web"
const GrpcWebTextContentType = "application/grpc-web-text"

func IsGrpcRequest(req *http.Request) bool {
	if req.Method == http.MethodPost && strings.HasPrefix(req.Header.Get("content-type"), GrpcContentType) {
		return true
	}
	if req.Method == http.MethodPost && strings.HasPrefix(req.Header.Get("content-type"), GrpcWebContentType) {
		return true
	}
	if req.Method == http.MethodPost && strings.HasPrefix(req.Header.Get("content-type"), GrpcWebTextContentType) {
		return true
	}
	return false
}

func IsGrpcWebRequest(req *http.Request) bool {
	return req.Method == http.MethodPost && strings.HasPrefix(req.Header.Get("content-type"), GrpcWebContentType)
}

func IsGrpcNativeRequest(req *http.Request) bool {
	if req.Method == http.MethodPost {
		if req.Header.Get("content-type") == GrpcContentType {
			return true
		}
	}

	return false
}

func IsGrpcRequestReflection(req *http.Request) bool {
	pp := strings.ToLower(req.URL.Path)
	return pp == "/grpc.reflection.v1alpha.serverreflection/serverreflectioninfo"
}
