package dowloader

import (
	"log"
	"os/exec"
)

func DownloadInstagram(url string) error {
	cmd := exec.Command("instaloader", "--video", url)

	err := cmd.Run()
	if err != nil {
		log.Printf("Ошибка скачивания с Instagram: %v", err)
		return err
	}

	log.Println("Видео с Instagram успешно скачано!")
	return nil
}
