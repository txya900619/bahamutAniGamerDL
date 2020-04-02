package animateDL

import (
	"fmt"
	"github.com/cheggaaa/pb/v3"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

type Animate struct {
	Sn            string
	M3u8Url       string `json:"src"`
	DeviceID      string `json:"deviceid"`
	noLoginCookie string
	Error         struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Status  string   `json:"status"`
		Details []string `json:"details"`
	} `json:"error"`
	Vip          bool `json:"vip"`
	SeenAD       int  `json:"time"`
	ChunkListUrl string
	ChunkList    []string
	TempFolder   string
	ProgressBar  *pb.ProgressBar
}

func (anime *Animate) DownloadM3u8() {
	anime.TempFolder = ".tmp" + anime.Sn
	os.Mkdir(anime.TempFolder, 0755)
	out, err := os.Create(anime.TempFolder + "/" + anime.ChunkListUrl)
	if err != nil {
		log.Fatal("create m3u8 file fail:" + err.Error())
	}
	defer out.Close()
	resp := anime.request("GET", strings.Replace(anime.M3u8Url, "playlist.m3u8", anime.ChunkListUrl, -1))
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal(" m3u8 file save fail:" + err.Error())
	}
}
func (anime *Animate) DownloadM3u8Key(url string) string {
	fileName := strings.Split(path.Base(url), "?")[0]
	out, err := os.Create(anime.TempFolder + "/" + fileName)
	if err != nil {
		log.Fatal("Fail to creat m3u8 key file:", err.Error())
	}
	defer out.Close()
	resp := anime.request("GET", url)
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatal("Fail to save m3u8 key file:", err.Error())
	}
	return fileName
}
func (anime *Animate) downloadChunk(url string) bool {
	filename := strings.Split(path.Base(url), "?")[0]

	fileStat, err := os.Stat(anime.TempFolder + "/" + filename)
	if err == nil && fileStat.Size() != 0 {
		return true
	}

	out, err := os.Create(anime.TempFolder + "/" + filename)
	if err != nil {
		log.Fatal("Fail to create chunk file:" + err.Error())
	}
	resp, err := anime.downloadRequest(url)
	if err != nil {
		fmt.Println("Download " + filename + " fail")
		fmt.Println("Retry")
		out.Close()
		os.Remove(anime.TempFolder + "/" + filename)
		time.Sleep(500 * time.Millisecond)
		return false
	}
	defer resp.Body.Close()
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		fmt.Println(filename + " save failed ")
		fmt.Println("Retry")
		out.Close()
		os.Remove(anime.TempFolder + "/" + filename)
		time.Sleep(500 * time.Millisecond)
		return false
	}
	out.Close()
	return true
}
func (anime *Animate) DownloadAnimate() {
	anime.ProgressBar = pb.StartNew(len(anime.ChunkList))
	xthreads := 64
	var ch = make(chan string)
	var wg sync.WaitGroup
	wg.Add(xthreads)
	for i := 0; i < xthreads; i++ {
		go func() {
			for {
				url, ok := <-ch
				if !ok {
					wg.Done()
					return
				}
				for {
					if anime.downloadChunk(url) {
						anime.ProgressBar.Increment()
						break
					}
				}
			}
		}()
	}
	for _, url := range anime.ChunkList {
		ch <- url
	}
	close(ch)
	wg.Wait()
	anime.ProgressBar.Finish()
}
