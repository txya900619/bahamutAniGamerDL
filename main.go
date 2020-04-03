package main

import (
	"flag"
	"github.com/txya900619/bahamutAniGamerDL/animateDL"
)

func main() {
	flag.Parse()
	animate := animateDL.Animate{Sn: flag.Arg(0)}
	animate.GetDeviceID()
	animate.GetAccess()
	if !animate.Vip {
		animate.SeeAD()
		animate.GetAccess()
		if animate.SeenAD == 0 {
			animate.SeeAD()
		}
	}
	animate.GetM3u8Url()
	animate.GetPlaylist()
	animate.DownloadM3u8()
	animate.ParseChunkList()
	animate.DownloadAnimate()
}
