package domain

type Buf struct {
	MsgType string `json:"msgType"` //regist,heartbeat,metric
	Msg     any    `json:"msg"`
}
