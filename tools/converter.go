package tools

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func Convert(input string, output string) {
	entries, err := os.ReadDir(input)
	if err != nil {
		fmt.Println("Error reading input directory:", err)
		return
	}

	// Ensure output directory exists
	if err := os.MkdirAll(output, os.ModePerm); err != nil {
		fmt.Println("Error creating output directory:", err)
		return
	}

	fmt.Printf("Total videos to convert: %d\n", len(entries))
	fmt.Println("Processing...")

	for index, e := range entries {
		if e.IsDir() {
			continue
		}

		inputFile := filepath.Join(input, e.Name())
		outputName := filepath.Join(output, e.Name()+"-output.mp4")

		cmd := exec.Command("ffmpeg", "-i", inputFile, "-q:v", "0", outputName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Printf("[%d] Error converting %s: %v\n", index+1, e.Name(), err)
			continue
		}

		fmt.Printf("[%d] Converted: %s → %s\n", index+1, e.Name(), outputName)
	}

	fmt.Println("All videos converted successfully.")
}
