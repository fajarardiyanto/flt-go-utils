package flags

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v3"
)

var (
	validate   *validator.Validate
	translator ut.Translator
)

func Init(configPath string, config interface{}) {
	if len(configPath) != 0 {
		if _, err := os.Stat(configPath); err == nil {
			if data, err := ioutil.ReadFile(filepath.Clean(configPath)); err == nil {
				if config != nil {
					if err := yaml.Unmarshal(data, config); err != nil {
						log.Fatal(err)
					}
				}
			}
		} else {
			ParseStruct(config)
		}
	}
}

func Validate(config interface{}) (err error) {
	if config != nil {
		validate = validator.New()
		if err = validate.Struct(config); err != nil {
			if val, ok := err.(validator.ValidationErrors); ok {
				for _, er := range val {
					return errors.New(er.Translate(translator))
				}
			}
		}
	}

	return err
}
