package services

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/kpango/glg"
)

const BaseDir string = "_data"
const certDir string = "cert"
const MakeRootFileName string = "makeRoot.sh"
const MakeCertFileName string = "makeCertificate.sh"

func getFilepath(filename string, isCert bool) string {
	if isCert {
		folder := strings.Split(filename, ".")[0]
		return fmt.Sprintf("%s/%s/%s/%s", BaseDir, certDir, folder, filename)	
	}

	return fmt.Sprintf("%s/%s", BaseDir, filename)
}

func createDataFolder() error {
	if _, err := os.Stat(BaseDir); !os.IsNotExist(err) {
		return nil
	}

	err := os.Mkdir(BaseDir, os.ModePerm); 
	if err != nil {
		return err
	}
	return nil
}

func checkIfCertExists(name string) (bool, error) {
	path := getFilepath(name, true)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		glg.Tracef("checkIfCertExists | found '%s'", path)
		return true, nil
	}

	glg.Tracef("checkIfCertExists | not found '%s'", path)
	return false, nil
}

func MakePath(filename string) (string, error) {
	if _, err := os.Stat(getFilepath(filename, false)); os.IsNotExist(err) {
		if _, err = DownloadMakeFileTemplate(filename); err != nil {
			return "", err
		}

		file, err := ReadMakeFile(filename)
		if err != nil {
			return "", err
		}
		file ,err = SetMinimalConfig(filename, file)
		if err != nil {
			return "", err
		}
		
		if err = WriteFile(filename, file); err != nil {
			return "", err
		}
	}
	
	return getFilepath(filename, false), nil
}

func DownloadMakeFileTemplate(filename string) (string, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/PsclDev/automate-ssl-certificates/main/%s", filename)
	glg.Tracef("DownloadMakeFileTemplate | from '%s'", url)
	res, err:= http.Get(url)
	if err != nil  {
   		return "", err
  	}
	defer res.Body.Close()

	template, err := io.ReadAll(res.Body)
	if err != nil  {
   		return "", err
  	}

	if err = WriteFile(filename, string(template)); err != nil {
		return "", err
	}

	return string(template), nil
}

func WriteFile(filename string, content string) error {
	err := createDataFolder()
	if err != nil  {
    	return err
  	}

	path := getFilepath(filename, false)
	glg.Tracef("WriteFile | to '%s'", path)
	file, err := os.Create(path)
	if err != nil  {
    	return err
  	}

	if _, err = file.WriteString(content); err != nil {
		return err
	}
	if err = file.Sync(); err != nil {
		return err
	}

	return nil
}

func ReadMakeFile(filename string) (string, error) {
	path, err := MakePath(filename)
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func CreateCertArchive(name string) (string, error) {
	crtName := fmt.Sprintf("%s.crt", name)
	crtPath := getFilepath(crtName, true)
	zipName := fmt.Sprintf("%s.zip", name)
	zipPath := getFilepath(zipName, true)

	crtInfo, err := os.Stat(crtPath)
	if err != nil {
		return "", err
	}

	zipExists := true
	zipInfo, err := os.Stat(zipPath)
	if err != nil {
		if os.IsNotExist(err) {
			zipExists = false
		} else {
			return "", err
		}
	}

	if (zipExists && zipInfo.ModTime().After(crtInfo.ModTime())) {
		glg.Tracef("CreateCertArchive | archive already exists '%s'", zipPath)
		return zipPath, nil
	}

	archive, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	forZip := [...]string{"crt", "csr", "key"}
	
	for _, ext := range forZip {
		filename := fmt.Sprintf("%s.%s", name, ext)
		filepath := getFilepath(filename, true)
		file, err := os.Open(filepath)
		if err != nil {
		if os.IsNotExist(err) {
			glg.Tracef("CreateCertArchive | file not found, skipping '%s'", filepath)
			continue
		} else {
			return "", err
		}
	}

		writer, err := zipWriter.Create(filename)
		if err != nil {
			return "", err
		}

		if _, err := io.Copy(writer, file); err != nil {
			return "", err
		}
	}

	zipWriter.Close()
	return zipPath, nil
}

func CreateCompleteArchive() (string, error) {
	zipPath := getFilepath("archive.zip", false)

	archive, err := os.Create(zipPath)
	if err != nil {
		return "", err
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	certPath := fmt.Sprintf("%s/%s", BaseDir, certDir)
	glg.Tracef("CreateCompleteArchive | cert path '%s'", certPath)

	directories, err := os.ReadDir(certPath)
	if err != nil {
		return "", err
	}
	
	for _, directory := range directories {
		if !directory.IsDir() {
			continue
		}

		dirPath := fmt.Sprintf("%s/%s/%s", BaseDir, certDir, directory.Name())
		glg.Tracef("CreateCompleteArchive | files in '%s'", dirPath)

		files, err := os.ReadDir(dirPath)
		if err != nil {
			return "", err
		}

		for _, file := range files {
			file, err := os.Open(getFilepath(file.Name(), true))
			if err != nil {
				return "", err
			}

			writer, err := zipWriter.Create(file.Name())
			if err != nil {
				return "", err
			}

			if _, err := io.Copy(writer, file); err != nil {
				return "", err
			}
		}
	}

	zipWriter.Close()
	return zipPath, nil
}