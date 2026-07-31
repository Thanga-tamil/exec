package config

import (
	"encoding/json"
	"os"

	logger "github.com/Thanga-tamil/exec/internal/utils"
)

type Conf struct{
	Host string
	Port string
}

func LoadConfig(path string) string {

	file, err := os.Open(path)
	if err != nil {
		logger.Fatalf("Err while opening %s file\n", path)
	}

	var res Conf
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&res); err != nil {
		logger.Fatalf("Err while unmarshal: %s\n", err.Error())
	}

	return res.Host + ":" + res.Port
}
