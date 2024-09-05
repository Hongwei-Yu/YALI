package global

type Engine struct {
	Status     int    `json:"status"`
	Node       string `json:"node"`
	Host       string `json:"host"`
	HostStatus string `json:"hostStatus"`
}

var GlobalEngine = Engine{
	Status:     0,
	Node:       "",
	Host:       "",
	HostStatus: "",
}
