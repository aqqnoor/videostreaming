// Динамикалық параметрлерді оқу
package config

import (
	pb "videostreaming/proto"
)

type Settings struct {
	OutputDir       string
	OutputFormat    string
	SegmentDuration int32
	FilenamePrefix  string
	Cleanup         bool
}

func FromRequest(req *pb.VideoRequest) *Settings {
	return &Settings{
		OutputDir:       "output",
		OutputFormat:    defaultFormat(req.OutputFormat),
		SegmentDuration: req.SegmentDuration,
		FilenamePrefix:  req.FilenamePrefix,
		Cleanup:         req.Cleanup,
	}
}

func defaultFormat(f string) string {
	if f == "" {
		return "mp4"
	}
	return f
}
