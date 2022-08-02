package configs

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

const (
	INPUT_DIR  string = "./input"
	OUTPUT_DIR string = "./output"

	CONFIG_EXTENSION string = ".yaml"

	MODELS_DIR   string = "models"
	HANDLERS_DIR string = "handlers"
	ROUTES_DIR   string = "routes"

	UPDATE_TIMESTAMP_TRIGGER string = "updated_timestamp"
)

var (
	DOCS_PATH    string = strings.Join([]string{OUTPUT_DIR, "docs"}, "/")
	SCHEMAS_PATH string = strings.Join([]string{DOCS_PATH, "schemas/"}, "/")
	PATHS_PATH   string = strings.Join([]string{DOCS_PATH, "paths/"}, "/")
)

var c config

type config struct {
	ENV         string
	BaseDocName string
	Repo        repo
}

type repo struct {
	User       string
	Key        string
	Service    string
	Migrations string
	Models     string
}

func init() {

	c.ENV = os.Getenv("ENV")

	c.Repo.User = os.Getenv("USER")
	c.Repo.Key = os.Getenv("KEY")
	c.Repo.Service = os.Getenv("SERVICE_REPO")
	c.Repo.Migrations = os.Getenv("MIGRATIONS_DIR")
	c.Repo.Models = os.Getenv("MODELS_REPO")

	c.BaseDocName = os.Getenv("BASE_DOC_NAME")

	err := checkConfig(c)
	if err != nil {
		panic(err)
	}
}

func GetConfig() config { return c }

func checkConfig(c config) error {

	var config map[string]interface{}
	data, _ := json.Marshal(c)
	json.Unmarshal(data, &config)

	return checkExistance(config)
}

func checkExistance(argh map[string]interface{}) error {

	for key, value := range argh {
		v, ok := value.(map[string]interface{})
		if ok {
			return checkExistance(v)
		} else {
			if value == "" {
				return fmt.Errorf("missing value for %s", key)
			}
		}
	}

	return nil
}
