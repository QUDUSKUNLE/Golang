package video

/* cut using a duration
ffmpeg -i input.mp4 -ss 00:05:20 -t 00:10:00 -c:v copy -c:a copy output1.mp4
*/

/* cut using a specific time
ffmpeg -i input.mp4 -ss 00:05:10 -to 00:15:30 -c:v copy -c:a copy output2.mp4
*/

/* ffprobe a video
ffprobe input.mp4
*/

/* Extract audio track
ffmpeg -i video.mkv -map 0:a -acodec copy audio.mp4
*/

/* Remove all audio tracks without re-encoding
ffmpeg -i input.mp4 -an -c:v copy output.mp4
*/

/* Remove an audio track
ffmpeg -i input.mp4 -map 0 -map -0:a:2 -c copy output.mp4
*/

/* Adding audio to a video
ffmpeg -i video.mp4 -i audio.mp3 -map 0:v -map 1:a -c:v copy -c:a copy -shortest output.mp4

with the shortest, it uses the shortest input as the output
*/

/* Fading out the audio
ffmpeg -i video.mp4 -i audio.mp3 -af "afade=out:st=10:d=2" -map 0:v -map 1:a -c:v copy -shortest output.mp4
*/

/* Get a file metadata
ffprobe -v quiet -print_format json -show_format -show_streams "lolwut.mp4" > "lolwut.mp4.json"
*/

/* Convert an mp4 file to mp3 and set to a bitrate of 64K
The higher the bitrate, the more the audio quality, the lower the bitrate the lower the audio quality.
ffmpeg -i inputFile.mp4 -b:a 64K output.mp3
*/

/* Get a file details
ffmpeg -i inputFile.mp4 -hide_banner
*/

/* Extract audio from a video
ffmpeg -i video.mp4 -vn audio.mp3
*/

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ProcessVideo() {
	// Define the folder containing the audio files
	folderPath := "/Users/andeladeveloper/Downloads/Hussory/030" // Replace with your folder path

	// Walk through the folder to find audio files
	err := filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		// Check if the file is an audio file (e.g., .mp3, .m4a)
		if !info.IsDir() && (strings.HasSuffix(info.Name(), ".mp3") || strings.HasSuffix(info.Name(), ".m4a")) {
			// Create a temporary output file in the same directory
			tempFile := path + ".tmp"

			// Run the ffmpeg command to omit images
			cmd := exec.Command("ffmpeg", "-i", path, "-vn", "-acodec", "copy", tempFile)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			fmt.Printf("Processing file: %s\n", path)
			if err := cmd.Run(); err != nil {
				fmt.Printf("Error processing file %s: %v\n", path, err)
				// Clean up the temporary file if it exists
				if removeErr := os.Remove(tempFile); removeErr != nil {
					fmt.Printf("Error removing temporary file %s: %v\n", tempFile, removeErr)
				}
				return fmt.Errorf("ffmpeg error for file %s: %v", path, err)
			}

			// Replace the original file with the processed file
			if err := os.Rename(tempFile, path); err != nil {
				fmt.Printf("Error replacing file %s: %v\n", path, err)
				// Clean up the temporary file if renaming fails
				if removeErr := os.Remove(tempFile); removeErr != nil {
					fmt.Printf("Error removing temporary file %s: %v\n", tempFile, removeErr)
				}
				return fmt.Errorf("rename error for file %s: %v", path, err)
			}

			fmt.Printf("Processed successfully: %s\n", path)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through folder: %v\n", err)
	}
}
