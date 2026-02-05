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

	// Unified metadata
	Sources  []string
	Warnings []string
	Error    string // empty = success

	// Payload (only one should be filled depending on Kind)
	FullName FullNameResult
	IP       IPResult
	Username UsernameResult
}

type FullNameResult struct {
	FirstName string
	LastName  string
	Address   string
	Phone     string
}

type IPResult struct {
	IP   string
	ISP  string
	City string
	Lat  float64
	Lon  float64
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

// ---- Constructors (recommended) ----

func NewBaseResult(kind Kind, input string) Result {
	return Result{
		Kind:      kind,
		Input:     input,
		Timestamp: time.Now(),
	}
}

func Fail(kind Kind, input string, err error) Result {
	r := NewBaseResult(kind, input)
	if err != nil {
		r.Error = err.Error()
	}
	return r
}
