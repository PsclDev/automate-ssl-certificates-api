package services

import (
	"api/models"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/kpango/glg"
)

func GetAllCertsAsJson() ([]*models.Certificate, error) {
	certPath := fmt.Sprintf("%s/%s", BaseDir, certDir)
	glg.Tracef("GetAllCertsAsJson | get direcotries from cert path ''", certPath)

	directories, err := os.ReadDir(certPath)
	if err != nil {
		return nil, err
	}

	certs := []*models.Certificate{}

	for _, directory := range directories {
		jsonPath := fmt.Sprintf("%s/%s/%s.json", certPath, directory.Name(), directory.Name())
		glg.Tracef("GetAllCertsAsJson | read file from json path ''", jsonPath)

		content, err := os.ReadFile(jsonPath)
		if err != nil {
			continue
		}

		cert := new(models.Certificate)
		err = json.Unmarshal(content, &cert)
		if err != nil {
			return nil, err
		}
		certs = append(certs, cert)
    }

	return certs, nil
}

func GetCertAsJson(name string) (*models.Certificate, error) {
	jsonPath := fmt.Sprintf("%s/%s/%s/%s.json", BaseDir, certDir, name, name)
	glg.Tracef("GetCertAsJson | read file from json path ''", jsonPath)

	content, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}

	cert := new(models.Certificate)
	err = json.Unmarshal(content, &cert)
	if err != nil {
		return nil, err
	}

	return cert, nil
}

func CheckRootCertificate() error {
	exists, err := checkIfCertExists("root.key")
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	if err := createRootCert(); err != nil {
		return err
	}

	return nil
}

func createRootCert() error {
	if _, err := exec.Command("bash", getFilepath(MakeRootFileName, false)).Output(); err != nil {
		return err
	}

	return nil
}

func CreateCert(cert *models.Certificate, forceCreate bool) error {
	if err := CheckRootCertificate(); err != nil {
		return err
	}

	exists, err := checkIfCertExists(fmt.Sprintf("%s.key", cert.Name))
	if err != nil {
		return err
	}
	if exists && !forceCreate {
		glg.Trace("CreateCert | Cert already exists and was not forced to recreate")
		return errors.New("CreateCert | cert already exists, use PATCH to recreate")
	}

	if _, err := exec.Command("bash", getFilepath(MakeCertFileName, false), "-d", cert.DNS, "-i", cert.IP, "-n", cert.Name).Output(); err != nil {
		return err
	}
	glg.Trace("CreateCert | created")

	return nil
}

func DeleteCert(certName string) error {
	path := fmt.Sprintf("%s/%s/%s", BaseDir, certDir, certName)
	glg.Tracef("DeleteCert | remove file from path ''", path)

	if err := os.RemoveAll(path); err != nil {
		return err
	}
	return nil
}