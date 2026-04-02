package utils

import (
	"os"
	"path/filepath"
	"strings"
)

func OrgProgramming() {
	ImagensPath := filepath.Join(os.Getenv("USERPROFILE"), "Downloads", "Programming")

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

		if strings.Contains(file.Name(), ".go") {
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
		} else if strings.Contains(file.Name(), ".py") {
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
		} else if strings.Contains(file.Name(), ".ts") {
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
		} else if strings.Contains(file.Name(), ".java") {
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
		} else if strings.Contains(file.Name(), ".c") {
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
		} else if strings.Contains(file.Name(), ".cpp") {
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
		} else if strings.Contains(file.Name(), ".cs") {
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
		} else if strings.Contains(file.Name(), ".php") {
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
		} else if strings.Contains(file.Name(), ".rb") {
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
		} else if strings.Contains(file.Name(), ".swift") {
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
		} else if strings.Contains(file.Name(), ".kt") {
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
