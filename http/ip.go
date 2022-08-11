package http

import (
	"fmt"
	"strings"
)

type CopyRequestValue struct {
	ContentType   string
	Origin        string
	XForwardedFor string
	FrontendIP    string
}

func GetIPAddress(r CopyRequestValue) (ip string, err error) {
	xip := r.XForwardedFor
	xpp := strings.Split(xip, ",")
	if len(xpp) != 0 {
		if len(xpp[0]) != 0 {
			return xpp[0], nil
		}
	}

	//second option
	xip = r.FrontendIP
	if len(xip) != 0 {
		return xip, nil
	}

	return ip, fmt.Errorf("%s", "Failed to get client ip address")
}
