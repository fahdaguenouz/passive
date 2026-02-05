package ip

import (
	"strings"

	"passive/internal/core"
	"passive/internal/detect"
)

func Run(query string) (core.Result, error) {
	q := strings.TrimSpace(query)

	if !detect.IsIPv4(q) {
		err := core.NewUserError("invalid IPv4 address")
		// return a unified Result AND an error (main already handles error printing)
		return core.Fail(core.KindIP, q, err), err
	}

	r := core.NewBaseResult(core.KindIP, q)

	// TODO: implement real lookup via provider(s)
	r.IP = core.IPResult{
		IP:   q,
		ISP:  "N/A",
		City: "N/A",
		Lat:  0.0,
		Lon:  0.0,
	}

	r.Warnings = append(r.Warnings, "IP lookup provider not implemented yet")
	r.Sources = append(r.Sources, "N/A")

	return r, nil
}
