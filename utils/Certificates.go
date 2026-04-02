package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func OrgCertificates() {
	ImagensPath := filepath.Join(os.Getenv("USERPROFILE"), "Downloads", "Certificates")

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

		if strings.Contains(file.Name(), ".pem") {
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
		} else if strings.Contains(file.Name(), ".key") {
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
		} else if strings.Contains(file.Name(), ".crt") {
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
		} else if strings.Contains(file.Name(), ".cer") {
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
		} else if strings.Contains(file.Name(), ".pfx") {
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
