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
