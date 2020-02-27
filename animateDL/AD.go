package animateDL

import "time"

func (anime *Animate) SeeAD() {
	anime.StartAd()
	time.Sleep(8 * time.Second)
	anime.SkipAd()
}
func (anime *Animate) StartAd() {
	anime.request("GET", "https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn="+anime.Sn+"&s=194699").Body.Close()
}

func (anime *Animate) SkipAd() {
	anime.request("GET", "https://ani.gamer.com.tw/ajax/videoCastcishu.php?sn="+anime.Sn+"&s=194699&ad=end").Body.Close()
}
