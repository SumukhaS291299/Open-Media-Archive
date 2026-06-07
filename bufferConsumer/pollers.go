package bufferconsumer

import "github.com/SumukhaS291299/Open-Media-Archive/commons"

type FfprobeStreamsPoller struct {
	ProbeStreams *commons.FfprobeStreams
}

type StringStreamsPoller struct {
	PlainTextOut *string
}

// Will soon implement multiple types down which will have custom json
// type FFmpegPoller struct {
// 	Progress FFmpegProgress
// 	Result   any
// }
