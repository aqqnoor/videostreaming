package ffmpeg

import (
	"fmt"
	"os/exec"
)

func SplitVideo(inputPath, outputDir string, duration int32, prefix string) error {
	if prefix == "" {
		prefix = "part_"
	}
	outputTemplate := fmt.Sprintf("%s/%s%%03d.mp4", outputDir, prefix)
	cmd := exec.Command("ffmpeg",
		"-i", inputPath,
		"-c", "copy",
		"-map", "0",
		"-f", "segment",
		"-segment_time", fmt.Sprintf("%d", duration),
		outputTemplate,
	)
	return cmd.Run()
}
