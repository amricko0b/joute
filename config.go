package joute

import (
	"encoding/json"
	"os"
)

// Used for testing purposes only - will be deleted
type ConfigMap map[string]any

// Used for testing purposes only - will be deleted
func LoadConfigMap() (ConfigMap, error) {

	var configMap ConfigMap

	if file, err := os.Open("./configs/.jouterc"); err == nil {
		err = json.NewDecoder(file).Decode(&configMap)
		return configMap, err
	} else {
		return nil, err
	}

}
