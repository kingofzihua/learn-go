package main

import (
	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/storage"
	"html"
)

const magnet string = "magnet:?xt=urn:btih:fc7cf482abfeda228a392cb182f4badef22c9f97&amp;dn=%E5%9B%BD%E4%BA%A7%E5%86%99%E7%9C%9F.JUQ-819+%E8%BF%85%E9%80%9F%E8%84%B1%E6%8E%89%E5%A5%B3%E5%AE%B6%E6%94%BF%E7%9A%84%E8%A3%A4%E5%AD%90%E6%8F%92%E5%85%A5%EF%BC%8C%E9%9C%B2%E5%87%BA%E5%8F%88%E5%A4%A7%E5%8F%88%E5%9C%86%E7%9A%84%E5%B1%81%E8%82%A1%E5%92%8C%E8%82%9B%E9%97%A8%E5%B9%B6%E7%96%AF%E7%8B%82%E5%B9%B2%E5%A5%B9_%E6%98%8E%E9%87%8C%E7%B4%AC.mp4&amp;xl=1413342412&amp;tr=udp://tracker.opentrackr.org:1337/announce&amp;tr=udp://9.rarbg.com:2810/announce&amp;tr=udp://tracker.openbittorrent.com:6969/announce&amp;tr=udp://opentracker.i2p.rocks:6969/announce&amp;tr=udp://tracker.dler.org:6969/announce&amp;tr=udp://zecircle.xyz:6969/announce&amp;tr=udp://www.peckservers.com:9000/announce&amp;tr=udp://wepzone.net:6969/announce&amp;tr=udp://vibeudp://v2.iperson.xyz:6969/announce&amp;tr=udp://v2.iperson.xyz:6969/announce&amp;tr=udp://tracker.torrent.eu.org:451/announce&amp;tr=udp://tracker.tiny-vps.com:6969/announce&amp;tr=udp://tracker.moeking.me:6969/"

func main() {
	config := torrent.NewDefaultClientConfig()
	config.DefaultStorage = storage.NewFileByInfoHash("./torrent-data")
	config.ListenPort = 10010

	client, err := torrent.NewClient(config)
	if err != nil {
		panic(err)
	}
	t, err := client.AddMagnet(html.UnescapeString(magnet))

	<-t.GotInfo()

	t.Stats()

	t.DownloadAll()

	client.WaitAll()
}
