package animateDL

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func (anime *Animate) GetDeviceID() {
	resp := anime.request("GET", "https://ani.gamer.com.tw/ajax/getdeviceid.php")
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("deviceID body read fail:" + err.Error())
	}
	err = json.Unmarshal(body, anime)
	if err != nil {
		log.Fatal("parse deviceID json fail:" + err.Error())
	}
	for _, cookie := range resp.Cookies() {
		if cookie.Name == "nologinuser" {
			anime.noLoginCookie = "nologinuser=" + cookie.Value
		}
	}
}
func (anime *Animate) GetAccess() {
	resp := anime.request("GET", "https://ani.gamer.com.tw/ajax/token.php?adID=0&sn="+anime.Sn+"&device="+anime.DeviceID+"&hash="+randomHash())
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("token body read fail:", err.Error())
	}
	err = json.Unmarshal(body, anime)
	if err != nil {
		log.Fatal("parse token json fail:" + err.Error())
	}

}
func randomHash() string {
	rand.Seed(time.Now().UnixNano())
	dictionary := "1234567890qwertyuiopasdfghjklzxcvbnm"
	hash := ""
	for i := 0; i < 12; i++ {
		hash = hash + string(dictionary[rand.Intn(len(dictionary))])
	}
	return hash
}
