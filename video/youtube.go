package video

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

// sanitizeFileName replaces unsafe characters with underscores
func sanitizeFileName(name string) string {
	// Replace all non-alphanumeric with "_"
	re := regexp.MustCompile(`[^a-zA-Z0-9._-]+`)
	safe := re.ReplaceAllString(name, "_")
	// Trim leading/trailing underscores
	return strings.Trim(safe, "_")
}

func DownloadAndConvert(url string, outputDir string) error {
	tmpDir := os.TempDir()
	title := "yt2mp3"

	// Raw output file path (yt-dlp decides extension, e.g. .webm or .m4a)
	tmpPattern := filepath.Join(tmpDir, title+".%(ext)s")

	// Step 1: Try direct mp3 extraction
	cmd := exec.Command("yt-dlp",
		"-x", "--audio-format", "mp3",
		"-o", tmpPattern,
		url,
	)
	if out, err := cmd.CombinedOutput(); err == nil {
		// ✅ Success, yt-dlp produced mp3 directly
		srcPath := filepath.Join(tmpDir, title+".mp3")
		dstPath := filepath.Join(outputDir, sanitizeFileName(title)+".mp3")
		return os.Rename(srcPath, dstPath)
	} else {
		fmt.Println("yt-dlp did not produce mp3, will try manual conversion...")
		fmt.Println(string(out))
	}

	// Step 2: Fallback – download bestaudio (likely .webm or .m4a)
	cmd = exec.Command("yt-dlp",
		"-f", "bestaudio",
		"-o", tmpPattern,
		url,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("yt-dlp bestaudio failed: %v\n%s", err, string(out))
	}

	// Figure out what yt-dlp wrote
	var audioPath string
	files, _ := filepath.Glob(filepath.Join(tmpDir, title+".*"))
	if len(files) > 0 {
		audioPath = files[0]
	} else {
		return fmt.Errorf("yt-dlp did not produce any audio file")
	}

	// Step 3: Sanitize filename for ffmpeg
	ext := filepath.Ext(audioPath)
	safeBase := sanitizeFileName(title) + ext
	safeAudioPath := filepath.Join(tmpDir, safeBase)

	if audioPath != safeAudioPath {
		if err := os.Rename(audioPath, safeAudioPath); err != nil {
			return fmt.Errorf("failed to rename file for ffmpeg: %v", err)
		}
		audioPath = safeAudioPath
	}

	// Step 4: Convert manually with ffmpeg
	mp3Path := filepath.Join(outputDir, sanitizeFileName(title)+".mp3")
	cmd = exec.Command("/usr/local/bin/ffmpeg",
		"-y", "-i", audioPath,
		"-vn", "-ar", "44100", "-ac", "2", "-b:a", "192k",
		mp3Path,
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("manual ffmpeg conversion failed: %v\n%s", err, string(out))
	}

	fmt.Println("✅ Conversion finished:", mp3Path)
	return nil
}

// DownloadYoutubeVideo downloads a YouTube video to the specified output directory.
// It preserves the best video quality available.
// youtubeURL: the YouTube video URL
// outputDir: the directory where the video will be saved
func DownloadYoutubeVideo(youtubeURL, outputDir string) error {
	// Validate input
	if youtubeURL == "" {
		return fmt.Errorf("YouTube URL cannot be empty")
	}
	if outputDir == "" {
		return fmt.Errorf("output directory cannot be empty")
	}

	fmt.Printf("Downloading from: %s\n", youtubeURL)
	fmt.Printf("Output directory: %s\n", outputDir)

	// Ensure outputDir exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// yt-dlp output template
	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")

	// Try multiple format options in order of preference
	formats := []string{
		"bestvideo[ext=mp4]+bestaudio[ext=m4a]/best[ext=mp4]",
		"best[ext=mp4]/best",
		"best",
	}

	var lastErr error
	for _, format := range formats {
		fmt.Printf("Trying format: %s\n", format)
		cmd := exec.Command("yt-dlp",
			"--no-update",
			"--socket-timeout", "30",
			"-f", format,
			"-o", outputTemplate,
			youtubeURL,
		)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			lastErr = err
			fmt.Printf("Format %s failed, trying next...\n", format)
			continue
		}

		// Success
		fmt.Printf("✅ Video downloaded successfully to: %s\n", outputDir)
		return nil
	}

	return fmt.Errorf("yt-dlp failed to download video with all format options: %v\nPlease update yt-dlp with: pip install --upgrade yt-dlp", lastErr)
}

// DownloadYoutubeAudio downloads only the audio from a YouTube video to the specified output directory.
// youtubeURL: the YouTube video URL
// outputDir: the directory where the audio will be saved
func DownloadYoutubeAudio(youtubeURL, outputDir string) error {
	// Ensure outputDir exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// yt-dlp output template
	outputTemplate := filepath.Join(outputDir, "%(title)s.%(ext)s")

	// Download the best audio
	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-o", outputTemplate, youtubeURL)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("yt-dlp failed to download audio: %v", err)
	}

	fmt.Printf("✅ Audio downloaded successfully to: %s\n", outputDir)
	return nil
}

// // main for CLI usage
// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Usage: go run convertYoutube.go <youtube_url>")
// 		return
// 	}
// 	url := os.Args[1]
// 	if err := ConvertYoutubeToMP3(url); err != nil {
// 		fmt.Println("Error:", err)
// 	}
// }
//
// https://www.youtube.com/watch?v=kvnt0mt9Grg
