package helpers

import (
    "io"
    "os"
    "fmt"
    "log"
    "os/user"
)

var F_DIRPATH string = GetUserHome() + "/.fiscaluno"

func init() {
    CheckFiscalunoDir()
    CopyConfigFile()
}

func GetUserHome() string {
    usr, err := user.Current()

    if err != nil {
        log.Fatal(err)
    }

    return usr.HomeDir
}

func CopyConfigFile() error {
    originalPath := os.Getenv("GOPATH")
    configFilePath := originalPath + "/src/github.com/fiscaluno/fiscaluno-api/config/config.json"

    inputFile, err := os.Open(configFilePath)
    if err != nil {
        fmt.Println(err)
    }
    defer inputFile.Close()

    _, err = os.Stat(F_DIRPATH + "/config.json")
    if err != nil {
        _, _ = os.Create(F_DIRPATH + "/config.json")
    }

    outputFile, err := os.OpenFile(F_DIRPATH + "/config.json", os.O_WRONLY, os.ModePerm)
    if err != nil {
        fmt.Println(err)
    }
    defer outputFile.Close()

    _, err = io.Copy(outputFile, inputFile)
    if err != nil {
        fmt.Println(err)
    }

    return outputFile.Close()
}

func CheckFiscalunoDir() {
    if _, err := os.Stat(F_DIRPATH); os.IsNotExist(err) {
        fmt.Println("Creating fiscaluno directory...")
        os.Mkdir(F_DIRPATH, os.ModePerm)
    } else {
        fmt.Println("Fiscaluno directory already exists...")
    }
}