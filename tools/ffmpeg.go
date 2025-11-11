package tools

import (
	"fmt"
	"os"
	"path/filepath"
)

func getFFmpegPath() (string, error) {
	execDir, err := os.Executable()
	if err != nil {
		return "", err
	}
	ffmpegPath := filepath.Join(filepath.Dir(execDir), "ffmpeg.exe")
	if _, err := os.Stat(ffmpegPath); os.IsNotExist(err) {
		return "", fmt.Errorf("ffmpeg.exe not found in app directory")
	}
	return ffmpegPath, nil
}