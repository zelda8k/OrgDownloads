package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func OrgDocuments() {
	ImagensPath := filepath.Join(os.Getenv("USERPROFILE"), "Downloads", "Documents")

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

		if strings.Contains(file.Name(), ".txt") {
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
		} else if strings.Contains(file.Name(), ".pdf") {
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
		} else if strings.Contains(file.Name(), ".doc") {
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
		} else if strings.Contains(file.Name(), ".docx") {
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
		} else if strings.Contains(file.Name(), ".odt") {
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
		} else if strings.Contains(file.Name(), ".rtf") {
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
		} else if strings.Contains(file.Name(), ".md") {
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
