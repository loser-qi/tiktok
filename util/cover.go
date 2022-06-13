package util

import (
	"os/exec"
	"path"
	"strings"
)

var DefaultCoverFilename = "cover.png"

// MakeCover 制作封面
func MakeCover(filename, dirname string) string {
	suffix := path.Ext(filename)
	if suffix != ".mp4" {
		return DefaultCoverFilename
	}
	coverFilename := strings.TrimSuffix(filename, suffix) + ".png"
	coverFilePath := dirname + coverFilename
	cmd := exec.Command("ffmpeg",
		"-ss", "0",
		"-t", "1",
		"-i", dirname+filename,
		"-vframes", "1",
		coverFilePath)

	if err := cmd.Run(); err != nil {
		return DefaultCoverFilename
	}
	return coverFilename
}
