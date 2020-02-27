package animateDL

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func (anime *Animate) GetM3u8() {
	resp := anime.request("GET", "https://ani.gamer.com.tw/ajax/m3u8.php?sn="+anime.Sn+"&device="+anime.DeviceID)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal("m3u8 body read fail:" + err.Error())
	}
	err = json.Unmarshal(body, anime)
	if err != nil {
		log.Fatal("parse m3u8 json fail:" + err.Error())
	}
	if anime.Error.Message != "" {
		fmt.Printf("%s\n", anime.Error.Message)
		fmt.Println(anime.Error.Code)
	}
}
func (anime *Animate) GetPlaylist() {
	resp := anime.request("GET", "https:"+anime.M3u8)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("playlist body read fail:" + err.Error())
	}
	fmt.Println(string(body))
}
