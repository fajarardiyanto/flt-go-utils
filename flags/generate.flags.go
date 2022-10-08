package flags

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GenerateConfig(cfg interface{}, output string) error {
	if err := os.MkdirAll(filepath.Clean(filepath.Dir(output)), 0750); err != nil {
		return err
	}
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	if stat, err := os.Stat(output); err == nil {
		if !stat.IsDir() {
			if b, err := ioutil.ReadFile(filepath.Clean(output)); err == nil {
				if err = yaml.Unmarshal(b, cfg); err == nil {
					if bs, _ := yaml.Marshal(cfg); err == nil {
						data = bs
					}
				}
			}
		}
	}

	if err = ioutil.WriteFile(output, data, 0600); err != nil {
		return err
	}

	return nil
}
