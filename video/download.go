package video

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Download() {
	bot, err := tgbotapi.NewBotAPI("YOUR_TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if strings.HasPrefix(update.Message.Text, "/download") {
			go handleDownloadCommand(bot, update.Message)
		}
	}
}

func handleDownloadCommand(bot *tgbotapi.BotAPI, msg *tgbotapi.Message) {
	args := strings.Fields(msg.Text)

	if len(args) < 3 {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Usage: /download <YouTube URL> <format>\nExample: /download https://youtu.be/abc123 mp3"))
		return
	}

	url := args[1]
	format := strings.ToLower(args[2])
	filename := "output"

	// Step 1: Download video using yt-dlp
	cmd := exec.Command("yt-dlp", "-o", filename+".%(ext)s", url)
	if err := cmd.Run(); err != nil {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Error downloading video: "+err.Error()))
		return
	}

	// Find downloaded file
	files, err := os.ReadDir(".")
	if err != nil {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Error reading downloaded file"))
		return
	}

	var downloadedFile string
	for _, file := range files {
		if strings.HasPrefix(file.Name(), filename) && !strings.HasSuffix(file.Name(), format) {
			downloadedFile = file.Name()
			break
		}
	}

	// Step 2: Convert using ffmpeg
	converted := filename + "." + format
	cmd = exec.Command("ffmpeg", "-i", downloadedFile, converted)
	if err := cmd.Run(); err != nil {
		bot.Send(tgbotapi.NewMessage(msg.Chat.ID, "Error converting file: "+err.Error()))
		return
	}

	// Step 3: Send file back to user
	file := tgbotapi.NewDocument(msg.Chat.ID, tgbotapi.FilePath(converted))
	bot.Send(file)

	// Cleanup
	os.Remove(downloadedFile)
	os.Remove(converted)
}

// DownloadVideo downloads a video from the given URL using yt-dlp.
// It returns the absolute file path of the downloaded video or an error.
func DownloadVideo(url string) (string, error) {
	url = sanitizeFacebookURL(url)
	downloadDir := "/Users/andeladeveloper/Downloads"

	if _, err := os.Stat(downloadDir); os.IsNotExist(err) {
		if err := os.Mkdir(downloadDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create download directory: %w", err)
		}
	}

	outputTemplate := filepath.Join(downloadDir, "%(title).80s.%(ext)s")

	// yt-dlp works for YouTube, Instagram, Facebook, Twitter, etc.
	cmd := exec.Command("yt-dlp",
		"-f", "b",
		"--no-playlist", // download only the video, not entire playlists
		"--no-simulate",
		"-o", outputTemplate,
		"--print", "filename", // get output filename
		url,
	)

	// Capture stderr and stdout
	var stderr strings.Builder
	cmd.Stderr = &stderr

	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("yt-dlp failed: %w\nDetails: %s", err, stderr.String())
	}

	downloadedFile := strings.TrimSpace(string(output))
	absPath, err := filepath.Abs(downloadedFile)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return "", fmt.Errorf("downloaded file not found at: %s", absPath)
	}

	return absPath, nil
}


func sanitizeFacebookURL(raw string) string {
	// If it's a /watch link, try to resolve it using an HTTP client (optional)
	// Otherwise, notify user to use a direct video link
	if strings.Contains(raw, "/watch") {
		log.Println("⚠️ Facebook /watch URLs are not supported by yt-dlp. Please use a direct video link.")
	}
	return raw
}
