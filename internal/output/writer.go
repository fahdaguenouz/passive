package output

import (
	"fmt"
	"os"
	"strings"
	"time"

	"passive/internal/core"
)

func WriteResult(filename string, r core.Result) error {
	body := formatForFile(r)
	return os.WriteFile(filename, []byte(body), 0o644)
}

func formatForFile(r core.Result) string {
	var b strings.Builder
	b.WriteString("passive v1.0.0\n")
	b.WriteString(fmt.Sprintf("Timestamp: %s\n", r.Timestamp.Format(time.RFC3339)))
	b.WriteString(fmt.Sprintf("Type: %s\n", r.Kind))
	b.WriteString(fmt.Sprintf("Input: %s\n\n", r.Input))

	switch r.Kind {
	case core.KindFullName:
		b.WriteString(fmt.Sprintf("First name: %s\n", r.FullName.FirstName))
		b.WriteString(fmt.Sprintf("Last name: %s\n", r.FullName.LastName))
		if r.FullName.Address != "" {
			b.WriteString(fmt.Sprintf("Address: %s\n", r.FullName.Address))
		}
		if r.FullName.Phone != "" {
			b.WriteString(fmt.Sprintf("Number: %s\n", r.FullName.Phone))
		}
		if r.FullName.Source != "" {
			b.WriteString(fmt.Sprintf("Source: %s\n", r.FullName.Source))
		}

	case core.KindIP:
		if r.IP.ISP != "" {
			b.WriteString(fmt.Sprintf("ISP: %s\n", r.IP.ISP))
		}
		if r.IP.City != "" {
			b.WriteString(fmt.Sprintf("City: %s\n", r.IP.City))
		}
		if r.IP.Lat != 0 || r.IP.Lon != 0 {
			b.WriteString(fmt.Sprintf("Lat/Lon: %.4f / %.4f\n", r.IP.Lat, r.IP.Lon))
		}
		if r.IP.Source != "" {
			b.WriteString(fmt.Sprintf("Source: %s\n", r.IP.Source))
		}

	case core.KindUsername:
		for _, n := range r.Username.Networks {
			val := "no"
			if n.Found {
				val = "yes"
			}
			b.WriteString(fmt.Sprintf("%s : %s\n", n.Name, val))
			if n.URL != "" {
				b.WriteString(fmt.Sprintf("  URL: %s\n", n.URL))
			}
		}
	}

	return b.String()
}
