package getui_model

type Single struct {
	Message      MessageItem      `json:"message"`
	Notification NotificationItem `json:"notification"`
	Cid          string           `json:"cid"`
	Requestid    string           `json:"requestid"`
}

type MessageItem struct {
	Appkey            string `json:"appkey"`
	IsOffline         bool   `json:"is_offline"`
	OfflineExpireTime int    `json:"offline_expire_time"`
	Msgtype           string `json:"msgtype"`
}
type NotificationItem struct {
	Style               StyleItem `json:"style"`
	TransmissionType    bool      `json:"transmission_type"`
	TransmissionContent string    `json:"transmission_content"`
}

type StyleItem struct {
	Type  int    `json:"type"`
	Text  string `json:"text"`
	Title string `json:"title"`
}

type RspBody struct {
	Result string `json:"result"`
	TaskID string `json:"taskid"`

	Status string `json:"status"`
}

// Generated by curl-to-Go: https://mholt.github.io/curl-to-go
type Auth struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	Appkey    string `json:"appkey"`
}
type Result struct {
	Result    string `json:"result"`
	AuthToken string `json:"auth_token"`
}

const (
	APPKEY       = "al0zZ6nvSO9tvxPPrTVHD9"
	APPID        = "g5heXUmtiv9UcdMBvEGJI1"
	MASTERSECRET = "UpL4wlvdSh7xWMdRyyHP21"
)
