package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func OrgVideos() {
	ImagensPath := filepath.Join(os.Getenv("USERPROFILE"), "Downloads", "Videos")

	err := os.MkdirAll(ImagensPath, 0777)
	if err != nil {
		panic(err)
	}

	downloads := filepath.Join(os.Getenv("USERPROFILE"), "Downloads")

	files, err := os.ReadDir(downloads)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		if strings.Contains(file.Name(), ".mp4") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		} else if strings.Contains(file.Name(), ".avi") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		} else if strings.Contains(file.Name(), ".mkv") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		} else if strings.Contains(file.Name(), ".mov") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		} else if strings.Contains(file.Name(), ".wmv") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		} else if strings.Contains(file.Name(), ".flv") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		} else if strings.Contains(file.Name(), ".webm") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		} else if strings.Contains(file.Name(), ".mpeg") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		} else if strings.Contains(file.Name(), ".3gp") {
			originPath := filepath.Join(downloads, file.Name())
			destinationPath := filepath.Join(ImagensPath, file.Name())

			err := os.Rename(originPath, destinationPath)
			if err != nil {
				data, err2 := os.ReadFile(originPath)
				if err2 != nil {
					continue
				}

				err2 = os.WriteFile(destinationPath, data, 0777)
				if err2 != nil {
					continue
				}
				os.Remove(originPath)
			}
		}
	}
}
