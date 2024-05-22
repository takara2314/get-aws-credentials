package main

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/ini.v1"
)

type AWSCredential struct {
	AccessKeyID     string
	SecretAccessKey string
	SessionToken    string
}

func getAWSCredentialFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".aws/credentials"), nil
}

func readFileAll(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	return os.ReadFile(path)
}

func main() {
	profiles := map[string]AWSCredential{}

	credentialFilePath, err := getAWSCredentialFilePath()
	if err != nil {
		panic(err)
	}

	credentialFile, err := ini.Load(credentialFilePath)
	if err != nil {
		panic(err)
	}

	sections := credentialFile.Sections()
	for i := range sections {
		profiles[sections[i].Name()] = AWSCredential{
			AccessKeyID:     sections[i].Key("aws_access_key_id").String(),
			SecretAccessKey: sections[i].Key("aws_secret_access_key").String(),
			SessionToken:    sections[i].Key("aws_session_token").String(),
		}
	}

	// とりあえず default のみ表示
	fmt.Printf("AWS_ACCESS_KEY_ID=%s\n", profiles["default"].AccessKeyID)
	fmt.Printf("AWS_SECRET_ACCESS_KEY=%s\n", profiles["default"].SecretAccessKey)
	fmt.Printf("AWS_SESSION_TOKEN=%s\n", profiles["default"].SessionToken)
}
