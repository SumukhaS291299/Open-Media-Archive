package main

import (
	bufferconsumer "ffmpegCompress/bufferConsumer"
	"ffmpegCompress/execit"
	"fmt"
)

func main() {
	stringout, err := execit.RunCmd(5, []string{
		"ffmpeg -version",
		"ffprobe -version",
		"Start-Sleep -Seconds 10; Write-Output 'Finished 10 seconds sleep'",
		"ls",
		"pwd",
	})
	if err != nil {
		fmt.Println(err)
	}
	stringbuffout := &bufferconsumer.StringStreamsPoller{}
	stringbuffout.Absorb(stringout)

	out, err := execit.RunCmd(3, []string{
		"ffprobe -v error -print_format json -show_streams 'path\\to\\media1'",
		"ffprobe -v error -print_format json -show_streams 'path\\to\\media2'",
		"ffprobe -v error -print_format json -show_streams 'path\\to\\media3'",
		// "curl -L -o test.mp4 https://download.blender.org/peach/bigbuckbunny_movies/big_buck_bunny_1080p_h264.mov && ffprobe -v error -print_format json -show_streams test.mp4",
	})
	if err != nil {
		fmt.Println(err)
	}

	probe := &bufferconsumer.FfprobeStreamsPoller{}
	probe.Absorb(out)

}
