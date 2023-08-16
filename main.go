package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Config struct {
	Marker    string `json:"marker"`
	MakeCopy  bool   `json:"makeCopy"`
	OutputDir string `json:"outputDir"`
}

func getConfigFilePath() string {
	exePath, _ := os.Executable()
	exeDir := filepath.Dir(exePath)
	return filepath.Join(exeDir, "renconfig.json")
}

func readConfig() (Config, error) {
	config := Config{Marker: "SO ", MakeCopy: true, OutputDir: "RenamedFiles"}

	configPath := getConfigFilePath()

	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		// If the file doesn't exist, create it with default value
		if os.IsNotExist(err) {
			err := writeConfig(config)
			if err != nil {
				return Config{}, err
			}
			return config, nil
		}
		return config, err
	}

	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func writeConfig(config Config) error {
	configBytes, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	configPath := getConfigFilePath()

	err = ioutil.WriteFile(configPath, configBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func renameAndCopyFiles(arrayOfFilePaths []string, marker string, copy bool, outDir string) {
	if len(arrayOfFilePaths) == 0 {
		fmt.Println("No files to process.")
		return
	}

	// Get the directory of the first file
	firstFilePath := arrayOfFilePaths[0]
	outputDir := filepath.Dir(firstFilePath)

	// Only create the new directory if copy is true
	var newDir string
	if copy {
		newDir = filepath.Join(outputDir, outDir)
		_, err := os.Stat(newDir)
		if os.IsNotExist(err) {
			err := os.Mkdir(newDir, 0755)
			if err != nil {
				fmt.Println("Error creating directory:", err)
				return
			}
		}
	}

	for _, filePath := range arrayOfFilePaths {
		fileName := filepath.Base(filePath)
		modifiedFileName := modifyFileName(fileName, marker)
		newFilePath := filepath.Join(newDir, modifiedFileName)
		modifiedFilePath := filepath.Join(outputDir, modifiedFileName)

		//if copy is false, rename the file, else copy the file
		if !copy {
			err := renameFile(filePath, modifiedFilePath)
			if err != nil {
				return
			}
			fmt.Printf("Renamed: %s -> %s\n", fileName, modifiedFileName)
		} else {
			err := copyFile(filePath, newFilePath)
			if err != nil {
				return
			}
			fmt.Printf("Copied and renamed: %s -> %s\n", fileName, modifiedFileName)
		}
	}
}

func modifyFileName(fileName, marker string) string {
	fmt.Println(fileName)
	if !strings.Contains(fileName, marker) {
		return fmt.Sprintf("_ERR_%s", fileName)
	}

	fmt.Println(fileName)

	markerIndex := strings.Index(fileName, marker)
	modifiedString := fileName[markerIndex:]

	splitArray := strings.Split(modifiedString, " - ")
	splitArray[0] = strings.ReplaceAll(splitArray[0], "-", ".")
	splitArray[0] = splitArray[0] + "_" + splitArray[1]
	fmt.Println(strings.Join(splitArray, " - "))
	return strings.Join(splitArray, " - ")
}

func renameFile(src, dst string) error {
	return os.Rename(src, dst)
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func main() {
	config, err := readConfig()
	if err != nil {
		fmt.Println("Error reading config:", err)
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}
	// Provide the paths of the files you want to process
	arrayOfFilePaths := os.Args[1:]

	if len(arrayOfFilePaths) == 0 {
		fmt.Println("No files to process.")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	marker := config.Marker
	makeCopy := config.MakeCopy
	outDir := strings.ReplaceAll(config.OutputDir, " ", "_")

	if outDir == "" {
		outDir = "RenamedFiles"
	} else if strings.ContainsAny(outDir, `\/:*?"<>|`) {
		fmt.Println("Output directory name contains illegal characters for windows directory name.")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	renameAndCopyFiles(arrayOfFilePaths, marker, makeCopy, outDir)
}
