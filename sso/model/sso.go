package model

type ServiceTicket struct {
	Service string
	Tgt     string
	St      string
}

type TGT struct {
	Tgt      string        `json:"tgt"`
	Username string        `json:"username"`
	St       []string        `json:"st"`
}
