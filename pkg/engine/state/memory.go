package state

type Memory struct {
	Context  map[string]interface{} `json:"context"`
	Response interface{}            `json:"response"`
}
