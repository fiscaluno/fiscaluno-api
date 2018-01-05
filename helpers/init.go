package helpers

import (
	"bufio"
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"os/user"
	"syscall"
)

var F_DIRPATH string = GetUserHome() + "/.fiscaluno"
var F_CONFFILE string = F_DIRPATH + "/config.json"
var F_LOGFILE string = F_DIRPATH + "/fiscaluno.log"

func init() {
	CheckFiscalunoDir()
	CheckConfigFile()
}

// User home directory
func GetUserHome() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return usr.HomeDir
}

// Checks if fiscaluno directory exists or not
// if not exists, creates directory
func CheckFiscalunoDir() {
	if _, err := os.Stat(F_DIRPATH); os.IsNotExist(err) {
		fmt.Println("Creating fiscaluno directory...")
		os.Mkdir(F_DIRPATH, os.ModePerm)
	} else {
		fmt.Println("Fiscaluno directory already exists...")
	}
}

// Checks if database config file exists or not
// if not exists, creates file
func CheckConfigFile() {
	if _, err := os.Stat(F_CONFFILE); os.IsNotExist(err) {
		createConfigFile()
	}
}

// Puts database information at config file
func createConfigFile() {
	configFile, err := os.OpenFile(F_CONFFILE, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	fileContent := configFileContent()
	if _, err := configFile.Write(fileContent); err != nil {
		log.Fatal(err)
	}
	defer configFile.Close()
}

// Config file json content
func configFileContent() []byte {
	var fileContentMap map[string]string

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\x1b[96;1mDatabase host: \x1b[m")
	databaseHost, _ := reader.ReadString('\n')
	fmt.Print("\x1b[96;1mDatabase name: \x1b[m")
	databaseName, _ := reader.ReadString('\n')
	fmt.Print("\x1b[96;1mDatabase user: \x1b[m")
	databaseUser, _ := reader.ReadString('\n')
	fmt.Print("\x1b[96;1mDatabase password: \x1b[m")
	databasePassInput, _ := terminal.ReadPassword(int(syscall.Stdin))
	databasePass := string(databasePassInput)

	fmt.Println("")
	fileContentMap = map[string]string{
		"DB_HOST":     databaseHost[0 : len(databaseHost)-1],
		"DB_NAME":     databaseName[0 : len(databaseName)-1],
		"DB_USER":     databaseUser[0 : len(databaseUser)-1],
		"DB_PASSWORD": databasePass,
	}

	fileContentJson, err := json.Marshal(fileContentMap)
	if err != nil {
		log.Fatal(err)
	}

	return fileContentJson
}
