package main

import (
	"log"
	"os"

	"github.com/QUDUSKUNLE/Golang/tutorial/video"
)

// // main for CLI usage
func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: ./main <YouTube-URL>")
	}
	url := os.Args[1]
	outDir := "/Users/andeladeveloper/Downloads"

	if err := video.DownloadAndConvert(url, outDir); err != nil {
		log.Fatal(err)
	}
}
