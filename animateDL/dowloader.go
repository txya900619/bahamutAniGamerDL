package animateDL

type Animate struct {
	Sn            string
	M3u8          string `json:"src"`
	DeviceID      string `json:"deviceid"`
	noLoginCookie string
	Error         struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Status  string   `json:"status"`
		Details []string `json:"details"`
	} `json:"error"`
	Vip    bool `json:"vip"`
	SeenAD int  `json:"time"`
}
