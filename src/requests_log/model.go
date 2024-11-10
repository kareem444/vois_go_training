package requests_log

import "time"

type RequestLog struct {
	Method     string    `json:"method" bson:"method"`
	URL        string    `json:"url" bson:"url"`
	StatusCode int       `json:"status_code" bson:"status_code"`
	ClientIP   string    `json:"client_ip" bson:"client_ip"`
	UserAgent  string    `json:"user_agent" bson:"user_agent"`
	Timestamp  time.Time `json:"timestamp" bson:"timestamp"`
}
