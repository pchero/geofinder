package request

// BodyFindsPOST is rquest body define for POST /finds
type BodyFindsPOST struct {
	Address   string   `json:"address"`
	Countries []string `json:"countries"`
}
