package animateDL

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/grafov/m3u8"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func (anime *Animate) GetM3u8Url() {
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
	anime.M3u8Url = "https:" + anime.M3u8Url
}
func (anime *Animate) GetPlaylist() {
	resp := anime.request("GET", anime.M3u8Url)
	defer resp.Body.Close()
	playList, listType, err := m3u8.DecodeFrom(bufio.NewReader(resp.Body), true)
	if err != nil {
		log.Fatal("m3u8 decode fail:" + err.Error())
	}
	if listType == m3u8.MASTER {
		masterpl := playList.(*m3u8.MasterPlaylist)
		for _, list := range masterpl.Variants {
			quality := strings.Split(list.Resolution, "x")[1]
			if quality == "720" {
				anime.ChunkListUrl = strings.Split(list.URI, "?")[0]
			}
		}
	}
}
func (anime Animate) ParseChunkList() {
	m3u8File, err := os.Open(anime.TempFolder + "/" + anime.ChunkListUrl)
	if err != nil {
		log.Fatal("Fail to read m3u8 file:", err.Error())
	}
	playList, listType, err := m3u8.DecodeFrom(bufio.NewReader(m3u8File), true)
	if err != nil {
		log.Fatal("chunkList decode fail:", err.Error())
	}
	if listType == m3u8.MEDIA {
		mediapl := playList.(*m3u8.MediaPlaylist)
		newPlayList, err := m3u8.NewMediaPlaylist(mediapl.WinSize(), mediapl.Count())
		if err != nil {
			log.Fatal("creat new media playlist fail:", err.Error())
		}
		chunkPrefix := strings.Split(anime.M3u8Url, "playlist.m3u8")[0]
		newPlayList.SetKey(mediapl.Key.Method, anime.DownloadM3u8Key(mediapl.Key.URI), "", "", "")
		for _, chunk := range mediapl.Segments {
			if chunk != nil {
				anime.ChunkList = append(anime.ChunkList, chunkPrefix+chunk.URI)
				newPlayList.Append(strings.Split(path.Base(chunk.URI), "?")[0], chunk.Duration, "")
			}
		}
		newPlayList.Close()
		ioutil.WriteFile(anime.TempFolder+"/"+anime.ChunkListUrl, newPlayList.Encode().Bytes(), 0755)
	}
}
