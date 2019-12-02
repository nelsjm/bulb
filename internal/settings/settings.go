package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func GetURL(name string) (string, error) {
	settings, err := GetSettings()
	if err != nil {
		return "", err
	}

	return settings[name], nil
}

func AddURL(name string, url string) error {
	settings, err := GetSettings()
	if err != nil {
		return err
	}

	settings[name] = url
	return setSettings(settings)
}

func GetSettings() (map[string]string, error) {
	settings := make(map[string]string)
	f, err := os.Open("settings.json")
	if os.IsNotExist(err) {
		cf, err := os.Create("settings.json")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cf.WriteString("{}")
		cf.Close()

		return settings, nil
	}

	if err != nil {
		return settings, err
	}

	defer f.Close()
	bits, err := ioutil.ReadAll(f)
	if err != nil {
		return settings, err
	}

	err = json.Unmarshal(bits, &settings)
	return settings, err
}

func setSettings(settings map[string]string) error {
	bits, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	f, err := os.Create("settings.json")
	if err != nil {
		return err
	}

	_, err = f.Write(bits)
	return err
}
