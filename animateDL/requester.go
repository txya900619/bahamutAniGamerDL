package animateDL

import (
	"log"
	"net/http"
)

func (anime *Animate) request(method, url string) *http.Response {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatal("create request fail:" + err.Error())
	}
	if anime.noLoginCookie != "" {
		req.Header.Add("cookie", anime.noLoginCookie)
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36")
	req.Header.Add("referer", "https://ani.gamer.com.tw/animeVideo.php?sn="+anime.Sn)
	req.Header.Add("origin", "https://ani.gamer.com.tw")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("request run fail:" + err.Error())
	}
	return resp
}
func (anime *Animate) downloadRequest(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("create request fail:" + err.Error())
	}
	if anime.noLoginCookie != "" {
		req.Header.Add("cookie", anime.noLoginCookie)
	}
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.120 Safari/537.36")
	req.Header.Add("referer", "https://ani.gamer.com.tw/animeVideo.php?sn="+anime.Sn)
	req.Header.Add("origin", "https://ani.gamer.com.tw")
	resp, err := http.DefaultClient.Do(req)
	return resp, err
}
