package config

import (
	"encoding/json"
	"github.com/fiscaluno/fiscaluno-api/helpers"
	"os"
	"reflect"
	"strings"
)

// Configuration file path
var FILE_PATH = helpers.F_DIRPATH

// Config struct
type Config struct {
	DB_HOST     string // DB_HOST variable
	DB_NAME     string // DB_NAME variable
	DB_USER     string // DB_USER variable
	DB_PASSWORD string // DB_PASSWORD variable
}

// Configuration struct initializer
var configuration Config

// SetUp configuration values
func SetUp() {
	config_file, _err := os.Open(FILE_PATH + "/config.json")

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
func GetConfig(conf string) (conf_field interface{}) {
	// Check if struct has values or not
	if reflect.DeepEqual(Config{}, configuration) {
		SetUp()
	}

	valr := reflect.ValueOf(configuration)
	conf_field = reflect.Indirect(valr).FieldByName(strings.ToUpper(conf)).Interface()

	return conf_field
}
