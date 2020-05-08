package hls

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestCmdExecuteFfmpeg(t *testing.T) {
	base, _ := os.Getwd()

	targetPath := path.Join(base, "assets", "hls")
	srcPath := path.Join(base, "assets", "raw", "sample.mov")
	ffmpegPath := "/usr/local/bin/ffmpeg"

	fmt.Println("targetPath", targetPath)
	fmt.Println("srcPath", srcPath)

	GenerateHLS(ffmpegPath, srcPath, targetPath, "480p")
}
