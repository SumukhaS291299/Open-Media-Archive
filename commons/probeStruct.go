package commons

type FfprobeStreams struct {
	Streams []struct {
		AvgFrameRate       string `json:"avg_frame_rate,omitzero"`
		BitRate            string `json:"bit_rate,omitzero"`
		BitsPerSample      int    `json:"bits_per_sample,omitempty"`
		ChannelLayout      string `json:"channel_layout,omitempty,omitzero"`
		Channels           int    `json:"channels,omitempty,omitzero"`
		ChromaLocation     string `json:"chroma_location,omitempty,omitzero"`
		CodecLongName      string `json:"codec_long_name,omitzero"`
		CodecName          string `json:"codec_name,omitzero"`
		CodecTag           string `json:"codec_tag,omitzero"`
		CodecTagString     string `json:"codec_tag_string,omitzero"`
		CodecType          string `json:"codec_type,omitzero"`
		CodedHeight        int    `json:"coded_height,omitempty,omitzero"`
		CodedWidth         int    `json:"coded_width,omitempty,omitzero"`
		DisplayAspectRatio string `json:"display_aspect_ratio,omitempty,omitzero"`
		Disposition        struct {
			AttachedPic     int `json:"attached_pic"`
			Captions        int `json:"captions"`
			CleanEffects    int `json:"clean_effects"`
			Comment         int `json:"comment"`
			Default         int `json:"default,omitzero"`
			Dependent       int `json:"dependent"`
			Descriptions    int `json:"descriptions"`
			Dub             int `json:"dub"`
			Forced          int `json:"forced"`
			HearingImpaired int `json:"hearing_impaired"`
			Karaoke         int `json:"karaoke"`
			Lyrics          int `json:"lyrics"`
			Metadata        int `json:"metadata"`
			Multilayer      int `json:"multilayer"`
			NonDiegetic     int `json:"non_diegetic"`
			Original        int `json:"original"`
			StillImage      int `json:"still_image"`
			TimedThumbnails int `json:"timed_thumbnails"`
			VisualImpaired  int `json:"visual_impaired"`
		} `json:"disposition"`
		DivxPacked        string `json:"divx_packed,omitempty,omitzero"`
		Duration          string `json:"duration,omitzero"`
		DurationTs        int    `json:"duration_ts,omitzero"`
		ExtradataSize     int    `json:"extradata_size,omitempty,omitzero"`
		HasBFrames        int    `json:"has_b_frames,omitempty,omitzero"`
		Height            int    `json:"height,omitempty,omitzero"`
		ID                string `json:"id,omitzero"`
		Index             int    `json:"index"`
		InitialPadding    int    `json:"initial_padding,omitempty"`
		Level             int    `json:"level,omitempty,omitzero"`
		MimeCodecString   string `json:"mime_codec_string,omitzero"`
		NbFrames          string `json:"nb_frames,omitzero"`
		PixFmt            string `json:"pix_fmt,omitempty,omitzero"`
		Profile           string `json:"profile,omitempty,omitzero"`
		QuarterSample     string `json:"quarter_sample,omitempty,omitzero"`
		RFrameRate        string `json:"r_frame_rate,omitzero"`
		SampleAspectRatio string `json:"sample_aspect_ratio,omitempty,omitzero"`
		SampleFmt         string `json:"sample_fmt,omitempty,omitzero"`
		SampleRate        string `json:"sample_rate,omitempty,omitzero"`
		StartPts          int    `json:"start_pts"`
		StartTime         string `json:"start_time,omitzero"`
		Tags              struct {
			HandlerName string `json:"handler_name,omitzero"`
			Language    string `json:"language,omitzero"`
		} `json:"tags"`
		TimeBase string `json:"time_base,omitzero"`
		Width    int    `json:"width,omitempty,omitzero"`
	} `json:"streams"`
}
