package services

import (
	"api/models"
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

func SetConfig(domainName *models.DomainName) error {
	if err := setConfig(MakeRootFileName, domainName); err != nil {
		return err
	}
	if err := setConfig(MakeCertFileName, domainName); err != nil {
		return err
	}

	return nil
}

func setConfig(filename string, domainName *models.DomainName) error {
	rootFile, err := ReadFile(filename)
	if err != nil {
		return err
	}

	rootFile = replaceAllVariables(domainName, rootFile)
	if err = WriteFile(filename, rootFile); err != nil {
		return err
	}

	return nil
}

func replaceAllVariables(domainName *models.DomainName, fileContent string) string {
	fileContent = replaceVariable(fileContent, "yourPath", BaseDir)
	fileContent = replaceVariable(fileContent, "rootPath", fmt.Sprintf("%s/%s/root", BaseDir, certDir))
	fileContent = replaceVariable(fileContent, "C", domainName.Country)
	fileContent = replaceVariable(fileContent, "ST", domainName.State)
	fileContent = replaceVariable(fileContent, "L", domainName.Location)
	fileContent = replaceVariable(fileContent, "O", domainName.Domain)
	fileContent = replaceVariable(fileContent, "CN", fmt.Sprintf("%s.%s", domainName.Domain, domainName.Tld))

	return fileContent
}

func replaceVariable(content string, variable string, value string) string {
	arr := strings.Split(content, "\n")
 	idx := slices.IndexFunc(arr, func(s string) bool { return strings.HasPrefix(s, variable)})
	if idx == -1 {
		return arrToString(arr)
	}
	
	// check that no variables will be changed which depends on bash vars
	if strings.HasPrefix(arr[idx], fmt.Sprintf("%s=$", variable)){
		return arrToString(arr)
	}

 	arr[idx] = fmt.Sprintf("%s=%s", variable, value)
 	return arrToString(arr)
 }

 func arrToString(arr []string) string {
	return strings.Join(arr, "\n")
 }