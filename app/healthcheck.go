package app

type status string

const (
	OK = "ok"
	ERROR = "error"
)

type HealthCheck struct {
	Status status `json:"status"`
}
