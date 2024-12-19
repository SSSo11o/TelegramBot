package dowloader

import (
	"bytes"
	"log"
	"os"
	"os/exec"
	"strings"
)

func DownloadVideo(url string) (string, error) {

	cmd := exec.Command("yt-dlp", "-g", url)

	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Printf("Ошибка скачивания видео: %v", err)
		return "", err
	}

	link := strings.TrimSpace(out.String())
	if link == "" {
		return "", err
	}

	return link, nil
}
