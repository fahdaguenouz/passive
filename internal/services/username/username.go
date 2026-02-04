package username

import (
	"strings"
	"time"

	"passive/internal/core"
	"passive/internal/detect"
)

func Run(query string) (core.Result, error) {
	q := strings.TrimSpace(query)
	if !detect.IsUsername(q) {
		return core.Result{}, core.NewUserError("invalid username format")
	}

	handle := strings.TrimPrefix(q, "@")

	// TODO: implement real passive checks (HTTP HEAD/GET) in next step
	results := make([]core.NetworkResult, 0, len(DefaultNetworks))
	for _, n := range DefaultNetworks {
		results = append(results, core.NetworkResult{
			Name:  n.Name,
			URL:   n.URL(handle),
			Found: false, // stub
		})
	}

	return core.Result{
		Kind:      core.KindUsername,
		Input:     q,
		Timestamp: time.Now(),
		Username: core.UsernameResult{
			Username: handle,
			Networks: results,
		},
	}, nil
}
