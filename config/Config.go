package config

import (
    "os"
    "strings"
    "reflect"
    "encoding/json"
)

// Configuration file path
var FILE_PATH, _ = os.Getwd()

// Config struct
type Config struct {
    DB_HOST string // DB_HOST variable
    DB_NAME string // DB_NAME variable
}

// Configuration struct initializer
var configuration Config

// SetUp configuration values
func SetUp() {
    config_file, _err := os.Open(FILE_PATH + "/config/config.json")

    if _err != nil {
        panic(_err)
    }

    decoder := json.NewDecoder(config_file)
    _err = decoder.Decode(&configuration)

    if _err != nil {
        panic(_err)
    }
}

// Get Config Struct value
func GetConfig(conf string) (conf_field interface {}){
    // Check if struct has values or not
    if reflect.DeepEqual(Config {}, configuration) {
        SetUp()
    }

    valr := reflect.ValueOf(configuration)
    conf_field = reflect.Indirect(valr).FieldByName(strings.ToUpper(conf)).Interface()

    return conf_field
}