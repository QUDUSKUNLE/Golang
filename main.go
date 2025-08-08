package main

import (
	"fmt"

	"github.com/QUDUSKUNLE/Golang/tutorial/video"
)

func main() {

	urls := []string{
		"https://www.youtube.com/watch?v=rbkgThC_klQ",
		// "https://www.instagram.com/reel/C5KxiNlL0Z1/",
		// "https://fb.watch/nXtUy5xQ6e/",
	}

	for _, u := range urls {
		fmt.Printf("🎬 Downloading: %s\n", u)
		path, err := video.DownloadVideo(u)
		if err != nil {
			fmt.Printf("❌ Failed: %v\n", err)
			continue
		}
		fmt.Printf("✅ Saved to: %s\n", path)
	}
}
