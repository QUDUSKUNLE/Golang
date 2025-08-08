package video

import (
	"fmt"
	// "io/ioutil" // deprecated, use os.ReadDir
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"golang.org/x/text/unicode/norm"
)

// sanitizeFileName replaces or removes problematic characters for file names
func sanitizeFileName(name string) string {
	replacer := strings.NewReplacer(
		" ", "_", "｜", "_", "|", "_", "'", "", "\"", "", ":", "_", "/", "_", "\\", "_",
		"<", "_", ">", "_", "?", "_", "*", "_",
	)
	return replacer.Replace(name)
}

// ConvertYoutubeToMP3 downloads a YouTube video from the given URL and converts it to mp3.
// The mp3 file will be saved in the specified outputDir with the video title as the filename.
func ConvertYoutubeToMP3(youtubeURL, outputDir string) error {
	tmpDir := os.TempDir()
	outputTemplate := filepath.Join(tmpDir, "yt2mp3_%(title)s.%(ext)s")

	cmd := exec.Command("yt-dlp", "-f", "bestaudio", "-o", outputTemplate, "--print", "%(filename)s", youtubeURL)
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("yt-dlp error: %v\n%s", err, string(out))
	}
	// Get the last non-empty line as the filename
	lines := strings.Split(string(out), "\n")
	audioPath := ""
	for i := len(lines) - 1; i >= 0; i-- {
		line := strings.TrimSpace(lines[i])
		if line != "" {
			audioPath = line
			break
		}
	}
	// Normalize the filename to NFC to match macOS filesystem
	audioPath = norm.NFC.String(audioPath)
	fmt.Printf("Normalized yt-dlp output filename: %q\n", audioPath)
	fmt.Printf("Normalized yt-dlp output filename bytes: %v\n", []byte(audioPath))

	// List all files in the temp directory for debugging
	files, _ := os.ReadDir(tmpDir)
	fmt.Println("Files in temp dir:")
	var matchedPath string
	for _, f := range files {
		diskName := filepath.Join(tmpDir, f.Name())
		diskNameNFC := norm.NFC.String(diskName)
		fmt.Printf("  %q\n", diskNameNFC)
		if diskNameNFC == audioPath {
			matchedPath = diskName
		}
	}
	if matchedPath != "" {
		audioPath = matchedPath
	}

	// Try to find the file by base name if direct stat fails
	if _, err := os.Stat(audioPath); err != nil {
		fmt.Println("Direct stat failed, trying to match by base name (Unicode-insensitive)...")
		baseName := filepath.Base(audioPath)
		files, _ := os.ReadDir(tmpDir)
		for _, f := range files {
			if utf8.ValidString(f.Name()) && (f.Name() == baseName || norm.NFC.String(f.Name()) == norm.NFC.String(baseName) || norm.NFD.String(f.Name()) == norm.NFD.String(baseName)) {
				audioPath = filepath.Join(tmpDir, f.Name())
				fmt.Printf("Matched file: %q\n", audioPath)
				break
			}
		}
	}

	if _, err := os.Stat(audioPath); err != nil {
		return fmt.Errorf("audio file not found: %v\nyt-dlp output:\n%s", err, string(out))
	}

	// Use the base name (without extension) as the mp3 file name, sanitized
	base := filepath.Base(audioPath)
	title := strings.TrimSuffix(base, filepath.Ext(base))
	title = sanitizeFileName(title)
	mp3Path := filepath.Join(outputDir, title+".mp3")

	// Ensure outputDir exists
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	// Sanitize the downloaded file name for ffmpeg (rename if needed)
	safeAudioPath := filepath.Join(tmpDir, title+filepath.Ext(audioPath))
	if safeAudioPath != audioPath {
		if err := os.Rename(audioPath, safeAudioPath); err != nil {
			return fmt.Errorf("failed to rename file: %v", err)
		}
		audioPath = safeAudioPath
	}

	// Convert to mp3
	cmd = exec.Command("ffmpeg", "-y", "-i", audioPath, "-vn", "-ar", "44100", "-ac", "2", "-b:a", "192k", mp3Path)
	if out, err := cmd.CombinedOutput(); err != nil {
		os.Remove(audioPath)
		return fmt.Errorf("ffmpeg error: %v\n%s", err, string(out))
	}

	os.Remove(audioPath)
	fmt.Printf("MP3 saved as: %s\n", mp3Path)
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
