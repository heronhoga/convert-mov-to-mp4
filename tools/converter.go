package converter

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Convert() {
	if err := os.MkdirAll("output", os.ModePerm); err != nil {
		fmt.Println("Error creating directory:", err)
		return
	}

	entries, err := os.ReadDir("input")
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	fmt.Println("Total video to convert: ", len(entries))
	fmt.Println("Processing..")
	for index, e := range entries {
		if !e.IsDir() {
			inputFile := filepath.Join("input", e.Name())
			outputName := filepath.Join("output", e.Name() + ".mp4")

			cmd := exec.Command("ffmpeg", "-i", inputFile, "-q:v", "0", outputName)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error running FFmpeg:", err)
				continue
			}

			fmt.Println("Converted video: ", index+1, " ", e.Name())
		}
	}
	fmt.Println("All videos have been converted :D")

}
