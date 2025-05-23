package service

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"videostreaming/internal/config"
	"videostreaming/internal/downloader"
	"videostreaming/internal/ffmpeg"
	"videostreaming/internal/minio"

	pb "videostreaming/proto"
)

type Server struct {
	pb.UnimplementedVideoStreamingServer
}

func (s *Server) ProcessVideo(ctx context.Context, req *pb.VideoRequest) (*pb.VideoPartsResponse, error) {
	if filepath.Ext(req.VideoUrl) != ".mp4" {
		return nil, fmt.Errorf("only .mp4 files are currently supported")
	}

	videoPath := "temp/video.mp4"
	if err := downloader.DownloadVideo(req.VideoUrl, videoPath); err != nil {
		return nil, fmt.Errorf("failed to download video: %v", err)
	}

	cfg := config.FromRequest(req)
	if err := os.MkdirAll(cfg.OutputDir, os.ModePerm); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %v", err)
	}

	if err := ffmpeg.SplitVideo(videoPath, cfg.OutputDir, cfg.SegmentDuration, cfg.FilenamePrefix); err != nil {
		return nil, fmt.Errorf("failed to split video: %v", err)
	}

	urls, err := minio.UploadParts(cfg.OutputDir, cfg.OutputFormat)
	if err != nil {
		return nil, fmt.Errorf("upload failed: %v", err)
	}

	if cfg.Cleanup {
		if err := os.Remove(videoPath); err != nil {
			return nil, fmt.Errorf("failed to clean up temp file: %v", err)
		}
	}

	return &pb.VideoPartsResponse{Parts: urls}, nil
}
