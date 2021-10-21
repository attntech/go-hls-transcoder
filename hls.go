// Package hls provides a few functionalities to generate HLS
// files using ffmpeg.
package hls

import (
	"os"
	"os/exec"
)

// GenerateHLS will generate HLS file based on resolution presets.
// The available resolutions are: 360p, 480p, 720p and 1080p.
func GenerateHLS(ffmpegPath, srcPath, targetPath, resolution string, isSync bool) error {
	options, err := getOptions(srcPath, targetPath, resolution)
	if err != nil {
		return err
	}

	return GenerateHLSCustom(ffmpegPath, options, isSync)
}

// GenerateHLSCustom will generate HLS using the flexible options params.s
// options is array of string that accepted by ffmpeg command
func GenerateHLSCustom(ffmpegPath string, options []string, isSync bool) error {
	cmd := exec.Command(ffmpegPath, options...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Start()
	if err != nil {
		return err
	}

	if isSync {
		err = cmd.Wait()
	}
	return err
}
