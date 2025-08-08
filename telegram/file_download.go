package telegram

import (
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// StartFileDownloadBot listens for file URLs, downloads the file, and sends it back to the channel.
func StartFileDownloadBot(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil || update.Message.Chat == nil || !update.Message.Chat.IsChannel() {
			continue
		}
		text := strings.TrimSpace(update.Message.Text)
		if text == "" || !isValidURL(text) {
			continue
		}

		filePath, err := downloadFileToTemp(text)
		if err != nil {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Failed to download file: "+err.Error())
			bot.Send(msg)
			continue
		}

		doc := tgbotapi.NewDocument(update.Message.Chat.ID, tgbotapi.FilePath(filePath))
		bot.Send(doc)
		os.Remove(filePath)
	}
}

// isValidURL checks if the text is a valid http(s) URL.
func isValidURL(text string) bool {
	return strings.HasPrefix(text, "http://") || strings.HasPrefix(text, "https://")
}

// downloadFileToTemp downloads a file from the given URL to a temporary file and returns the file path.
func downloadFileToTemp(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Try to get filename from URL or Content-Disposition header
	filename := "downloaded_file"
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		if parts := strings.Split(cd, "filename="); len(parts) > 1 {
			filename = strings.Trim(parts[1], "\"")
		}
	} else if urlParts := strings.Split(url, "/"); len(urlParts) > 0 {
		filename = urlParts[len(urlParts)-1]
	}

	tmpFile, err := os.CreateTemp("", filename)
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, resp.Body)
	return tmpFile.Name(), err
}
