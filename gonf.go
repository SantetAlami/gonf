package gonf

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func Load(conf interface{}, configPath string, environment string) error {
    // Read YAML file
	if len(environment) > 0 {
		configPath = configPath[:len(configPath)-5] + "." + environment + ".yaml"
	}

    data, err := ioutil.ReadFile(configPath)
    if err != nil {
        return err
    }

    // Parse YAML data into struct
    return yaml.Unmarshal(data, conf)
}
