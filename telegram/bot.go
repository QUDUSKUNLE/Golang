package telegram

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// StartYouTubeToMP3BotDynamic listens for YouTube links in any channel the bot is invited to.
// It dynamically picks up the channel username or ID and responds to YouTube links.
func StartYouTubeToMP3BotDynamic(token string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	youtubeRegex := regexp.MustCompile(`https?://(www\.)?(youtube\.com|youtu\.be)/[^\s]+`)

	for update := range updates {
		if update.Message != nil && update.Message.Chat != nil && update.Message.Chat.IsChannel() {
			// Get channel username or ID
			chatUser := update.Message.Chat.UserName // may be empty for private channels
			chatID := update.Message.Chat.ID

			text := update.Message.Text
			matches := youtubeRegex.FindAllString(text, -1)
			if len(matches) == 0 {
				continue
			}

			for _, ytURL := range matches {
				title, mp3Path, err := downloadAndConvertToMP3(ytURL)
				if err != nil {
					msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Failed to process: %v", err))
					bot.Send(msg)
					continue
				}

				// Compose caption with channel info
				channelInfo := chatUser
				if channelInfo == "" {
					channelInfo = fmt.Sprintf("Channel ID: %d", chatID)
				}
				caption := fmt.Sprintf("🎵 *%s*\nHere is the mp3 version of your YouTube link.\n_Processed for %s_", title, channelInfo)
				audioMsg := tgbotapi.NewAudio(chatID, tgbotapi.FilePath(mp3Path))
				audioMsg.Caption = caption
				audioMsg.ParseMode = "Markdown"
				bot.Send(audioMsg)

				os.Remove(mp3Path)
			}
		}
	}
}

// downloadAndConvertToMP3 downloads a YouTube video and converts it to mp3, returning the title and mp3 path.
func downloadAndConvertToMP3(ytURL string) (string, string, error) {
	tmpDir := os.TempDir()
	outputTemplate := filepath.Join(tmpDir, "yt2mp3_%(title)s.%(ext)s")

	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-o", outputTemplate, "--print", "%(title)s", ytURL)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", "", fmt.Errorf("yt-dlp error: %v\n%s", err, string(out))
	}
	title := strings.TrimSpace(string(out))
	if title == "" {
		title = "audio"
	}

	var audioPath string
	files, _ := filepath.Glob(filepath.Join(tmpDir, "yt2mp3_"+title+".*"))
	for _, f := range files {
		if !strings.HasSuffix(f, ".mp3") {
			audioPath = f
			break
		}
	}
	if audioPath == "" {
		return title, "", fmt.Errorf("audio file not found")
	}

	mp3Path := filepath.Join(tmpDir, "yt2mp3_"+title+".mp3")
	cmd = exec.Command("ffmpeg", "-y", "-i", audioPath, "-vn", "-ar", "44100", "-ac", "2", "-b:a", "192k", mp3Path)
	if out, err := cmd.CombinedOutput(); err != nil {
		os.Remove(audioPath)
		return title, "", fmt.Errorf("ffmpeg error: %v\n%s", err, string(out))
	}

	os.Remove(audioPath)
	return title, mp3Path, nil
}
