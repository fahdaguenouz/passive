package ip

import (
	"strings"
	"time"

	"passive/internal/core"
	"passive/internal/detect"
)

func Run(query string) (core.Result, error) {
	q := strings.TrimSpace(query)
	if !detect.IsIPv4(q) {
		return core.Result{}, core.NewUserError("invalid IPv4 address")
	}

	// TODO: implement real lookup via provider(s) in providers.go
	isp, city, lat, lon, source := "N/A", "N/A", 0.0, 0.0, "N/A (not implemented yet)"

	return core.Result{
		Kind:      core.KindIP,
		Input:     q,
		Timestamp: time.Now(),
		IP: core.IPResult{
			IP:     q,
			ISP:    isp,
			City:   city,
			Lat:    lat,
			Lon:    lon,
			Source: source,
		},
	}, nil
}
