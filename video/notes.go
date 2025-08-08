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
	"bytes"
	"encoding/json"
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

func ConvertAllToMP3(dirPath string) {
	// Ensure the directory exists
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		fmt.Printf("Directory does not exist: %s\n", dirPath)
		return
	}

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Skip files that are already .mp3
		if strings.HasSuffix(strings.ToLower(info.Name()), ".mp3") {
			return nil
		}

		// Output file path with .mp3 extension
		outputFile := strings.TrimSuffix(path, filepath.Ext(path)) + ".mp3"

		// ffmpeg command to convert to mp3
		cmd := exec.Command("ffmpeg", "-y", "-i", path, "-vn", "-ar", "44100", "-ac", "2", "-b:a", "192k", outputFile)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		fmt.Printf("Converting %s to %s\n", path, outputFile)
		if err := cmd.Run(); err != nil {
			fmt.Printf("Error converting file %s: %v\n", path, err)
		} else {
			fmt.Printf("Successfully converted: %s\n", outputFile)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking through directory: %v\n", err)
	}
}

// RemoveBackgroundNoise uses ffmpeg's afftdn filter to reduce background noise in an audio file.
// inputPath: path to the original audio file
// outputPath: path to save the cleaned audio file
func RemoveBackgroundNoise(inputPath, outputPath string) error {
	cmd := exec.Command("ffmpeg", "-y", "-i", inputPath, "-af", "afftdn", outputPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

type AudioMetadata struct {
	Streams []struct {
		CodecName  string `json:"codec_name"`
		Channels   int    `json:"channels"`
		SampleRate string `json:"sample_rate"`
		BitRate    string `json:"bit_rate"`
		Duration   string `json:"duration"`
	} `json:"streams"`
}

// GetAudioMetadata uses ffprobe to extract audio metadata from a file
func GetAudioMetadata(filePath string) (*AudioMetadata, error) {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "a:0", "-show_entries", "stream=codec_name,channels,sample_rate,bit_rate,duration", "-of", "json", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return nil, err
	}
	var meta AudioMetadata
	if err := json.Unmarshal(out.Bytes(), &meta); err != nil {
		return nil, err
	}
	return &meta, nil
}

// SaveAudioMetadata uses ffprobe to extract audio metadata from a file and saves it as JSON in the specified directory.
func SaveAudioMetadata(filePath, saveDir string) error {
	cmd := exec.Command("ffprobe", "-v", "error", "-select_streams", "a:0", "-show_entries", "stream=codec_name,channels,sample_rate,bit_rate,duration", "-of", "json", filePath)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &out
	if err := cmd.Run(); err != nil {
		return err
	}

	// Ensure saveDir exists
	if err := os.MkdirAll(saveDir, 0755); err != nil {
		return err
	}

	// Create output file path
	base := filepath.Base(filePath)
	jsonFile := filepath.Join(saveDir, base+".metadata.json")

	// Write JSON output to file
	return os.WriteFile(jsonFile, out.Bytes(), 0644)
}

// CutOutAudioSegment removes a segment from an audio file and returns the path to the new file.
// inputPath: original audio file
// start: start time of the segment to remove (e.g., "00:01:00")
// end: end time of the segment to remove (e.g., "00:02:00")
// outputPath: path to save the new audio file
func CutOutAudioSegment(inputPath, start, end, outputPath string) error {
    // Temporary files for the two segments
    before := inputPath + ".before.mp3"
    after := inputPath + ".after.mp3"

    // 1. Extract before segment
    cmd1 := exec.Command("ffmpeg", "-y", "-i", inputPath, "-ss", "0", "-to", start, "-c", "copy", before)
    cmd1.Stdout = os.Stdout
    cmd1.Stderr = os.Stderr
    if err := cmd1.Run(); err != nil {
        return err
    }

    // 2. Extract after segment
    cmd2 := exec.Command("ffmpeg", "-y", "-i", inputPath, "-ss", end, "-c", "copy", after)
    cmd2.Stdout = os.Stdout
    cmd2.Stderr = os.Stderr
    if err := cmd2.Run(); err != nil {
        os.Remove(before)
        return err
    }

    // 3. Concatenate both segments
    listFile := inputPath + ".concat.txt"
    if err := os.WriteFile(listFile, []byte("file '"+before+"'\nfile '"+after+"'\n"), 0644); err != nil {
        os.Remove(before)
        os.Remove(after)
        return err
    }
    cmd3 := exec.Command("ffmpeg", "-y", "-f", "concat", "-safe", "0", "-i", listFile, "-c", "copy", outputPath)
    cmd3.Stdout = os.Stdout
    cmd3.Stderr = os.Stderr
    err := cmd3.Run()

    // Clean up
    os.Remove(before)
    os.Remove(after)
    os.Remove(listFile)
    return err
}
