package bufferconsumer

import (
	"encoding/json"
	"ffmpegCompress/execit"
	"fmt"
	"strings"
	"sync"
)

// Read the output based on the command run
// If ffmpeg used parse json out
// If data is progress, logic to consume the progress
// If data output needed to be raw or is unknown at execution, simply put the output buffer
// If data gives no output let know when it's done

// For now print logic here will be soon changed to a separate file

type BufferPoll interface {
	Absorb(execit.IOBuff) error
}

func (probeStreams *FfprobeStreamsPoller) Absorb(iobuffs []execit.IOBuff) error {
	var wg sync.WaitGroup
	wg.Add(len(iobuffs))
	// For now no channel as the goroutine will print directly to console
	for _, iobuff := range iobuffs {
		go func(op execit.IOBuff) {
			data := op.Stdout
			outJson := probeStreams.ProbeStreams
			// fmt.Println(string(data))
			if err := json.Unmarshal(data, &outJson); err != nil {
				fmt.Println(err)
			}
			// fmt.Println(outJson)
			// Don't do only the first index !!
			// fmt.Println(outJson.Streams[0].BitRate)
			// fmt.Println(outJson.Streams[0].CodecName)
			// fmt.Println(outJson.Streams[0].CodecType)
			// fmt.Println(outJson.Streams[0].CodedHeight)
			// fmt.Println(outJson.Streams[0].CodedWidth)
			pretty, _ := json.MarshalIndent(outJson, "", "  ")
			fmt.Println(string(pretty))
			wg.Done()
		}(iobuff)

	}
	wg.Wait()
	return nil
}

func (stringStream *StringStreamsPoller) Absorb(iobuffs []execit.IOBuff) error {
	var wg sync.WaitGroup
	wg.Add(len(iobuffs))
	// For now no channel as the goroutine will print directly to console
	for _, iobuff := range iobuffs {
		go func(op execit.IOBuff) {
			fmt.Println("##############################")
			fmt.Println("Output:\n\t")
			fmt.Print(string(op.Stdout))
			if stderr := strings.TrimSpace(string(op.Stderr)); stderr != "" {
				fmt.Println("Error:")
				fmt.Println(stderr)
			}
			wg.Done()
		}(iobuff)
	}
	wg.Wait()
	return nil
}
