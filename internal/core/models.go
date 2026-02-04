package core

import "time"

type Kind string

const (
	KindFullName Kind = "full-name"
	KindIP       Kind = "ip"
	KindUsername Kind = "username"
)

type Result struct {
	Kind      Kind
	Input     string
	Timestamp time.Time

	FullName  FullNameResult
	IP        IPResult
	Username  UsernameResult
}

type FullNameResult struct {
	FirstName string
	LastName  string
	Address   string
	Phone     string
	Source    string
}

type IPResult struct {
	IP     string
	ISP    string
	City   string
	Lat    float64
	Lon    float64
	Source string
}

type UsernameResult struct {
	Username string
	Networks []NetworkResult
}

type NetworkResult struct {
	Name  string
	URL   string
	Found bool
}
