/*
 * Copyright (c) 2020 Devtron Labs
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

const (
	SSH_PRIVATE_KEY_DIR       = ".ssh"
	SSH_PRIVATE_KEY_FILE_NAME = "id_rsa"
	GIT_CREDENTIAL_FILE_NAME  = ".git-credentials"
	SSH_CONFIG_FILE_NAME      = "config"
)

func CreateSshPrivateKeyOnDisk(fileId int, sshPrivateKeyContent string) error {

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	sshPrivateKeyFilePath := path.Join(userHomeDirectory, SSH_PRIVATE_KEY_DIR, SSH_PRIVATE_KEY_FILE_NAME)

	// if file exists then delete file
	if _, err := os.Stat(sshPrivateKeyFilePath); os.IsExist(err) {
		os.Remove(sshPrivateKeyFilePath)
	}

	// create file with content
	err = ioutil.WriteFile(sshPrivateKeyFilePath, []byte(sshPrivateKeyContent), 0600)
	if err != nil {
		return err
	}

	sshConfigFilePath := path.Join(userHomeDirectory, SSH_PRIVATE_KEY_DIR, SSH_CONFIG_FILE_NAME)

	if _, err := os.Stat(sshConfigFilePath); os.IsExist(err) {
		err = os.Chmod(sshConfigFilePath, 0644)
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateGitCredentialFileAndWriteData(data string) error {

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fileName := path.Join(userHomeDirectory, GIT_CREDENTIAL_FILE_NAME)

	// if file exists then delete file
	if _, err := os.Stat(fileName); os.IsExist(err) {
		os.Remove(fileName)
	}

	// create file with content
	err = ioutil.WriteFile(fileName, []byte(data), 0600)
	if err != nil {
		return err
	}

	return nil
}

func CleanupAfterFetchingHttpsSubmodules() error {

	userHomeDirectory, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// remove ~/.git-credentials
	gitCredentialsFile := path.Join(userHomeDirectory, GIT_CREDENTIAL_FILE_NAME)
	if _, err := os.Stat(gitCredentialsFile); os.IsExist(err) {
		os.Remove(gitCredentialsFile)
	}

	return nil
}

func LogStage(name string) {
	stageTemplate := `
------------------------------------------------------------------------------------------------------------------------
STAGE:  %s
------------------------------------------------------------------------------------------------------------------------`
	log.Println(fmt.Sprintf(stageTemplate, name))
}

var chars = []rune("abcdefghijklmnopqrstuvwxyz0123456789")

//Generates random string
func Generate(size int) string {
	rand.Seed(time.Now().UnixNano())
	var b strings.Builder
	for i := 0; i < size; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()
	return str
}
