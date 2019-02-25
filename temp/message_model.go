package temp

type Command struct {
	Action  string                 `json:"action"`
	Content map[string]interface{} `json:"content"`
}

type PerfInfo struct {
	TS     int64   `json:"ts"`
	CPU    float32 `json:"cpu"`
	Memory int     `json:"mem"`
	NetIn  int     `json:"net_in"`
	NetOut int     `json:"net_out"`
}
