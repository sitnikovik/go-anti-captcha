package task

// Task is response from AntiCaptcha API to captcha resolving
type Task struct {
	ErrorID  int    `json:"errorId"`
	TaskID   int    `json:"taskId"`
	Status   string `json:"status"`
	Solution struct {
		Text string `json:"text"`
		Url  string `json:"url"`
	} `json:"solution"`
	Cost       string `json:"cost"`
	Ip         string `json:"ip"`
	CreateTime int    `json:"createTime"`
	EndTime    int    `json:"endTime"`
	SolveCount int    `json:"solveCount"`
}
