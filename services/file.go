package services

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const BaseDir string = "_data"
const certDir string = "cert"
const MakeRootFileName string = "makeRoot.sh"
const MakeCertFileName string = "makeCertificate.sh"

func getFilepath(filename string, isCert bool) string {
	if isCert {
		return fmt.Sprintf("%s/%s/%s", BaseDir, certDir, filename)	
	}

	return fmt.Sprintf("%s/%s", BaseDir, filename)
}

func createDataFolder() error {
	if _, err := os.Stat(BaseDir); !os.IsNotExist(err) {
		return nil
	}

	err := os.Mkdir(BaseDir, os.ModePerm); 
	if err == nil {
		return nil
	}
	return err
}

func MakePath(filename string) (string, error) {
	if _, err := os.Stat(getFilepath(filename, false)); os.IsNotExist(err) {
		if _, err = DownloadMakeFileTemplate(filename); err != nil {
			return "", err
		}
	}
	
	return getFilepath(filename, false), nil
}

func DownloadMakeFileTemplate(filename string) (string, error) {
	url := fmt.Sprintf("https://raw.githubusercontent.com/PsclDev/automate-ssl-certificates/main/%s", filename)
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

	file, err := os.Create(getFilepath(filename, false))
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

func ReadFile(filename string) (string, error) {
	path, err := MakePath(filename)
	if err != nil {
		return "", nil
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return "", nil
	}

	return string(content), nil
}