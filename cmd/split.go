package cmd

import (
	"errors"
	"io"
	"log"
	"os"
	"time"

	"encoding/json"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(splitCmd)
}

type Doc struct {
	Entries []Entry `json:"entries"`
}

type Entry struct {
	CreationDate time.Time `json:"creationDate"`
	Text         string    `json:"text"`
}

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split exports",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		inputPath := args[0]
		outputDir := args[1]

		if _, err := os.Stat(inputPath); errors.Is(err, os.ErrNotExist) {
			log.Fatal("input file can't be found")
		}

		if _, err := os.Stat(outputDir); errors.Is(err, os.ErrNotExist) {
			log.Fatal("outputDir can't be found")
		}

		readJson(inputPath, outputDir)
	},
}

func readJson(inputPath string, outputDir string) {
	var doc Doc
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(data, &doc)

	writeAll(doc.Entries, outputDir)
}

func writeAll(entries []Entry, outputDir string) {
	for _, entry := range entries {
		write(entry, outputDir)
	}
}

func write(entry Entry, dir string) {
	isoDay := entry.CreationDate.Format("2006-01-02")
	path := dir + "/" + isoDay + ".md"

	// create file
	f, err := os.Create(path)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	data := []byte(entry.Text)
	_, err = f.Write(data)

	if err != nil {
		log.Fatal(err)
	}
	f.Sync()

}

func init() {
	// nothing to init
}
