package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
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
	settingFile, err := getSettingFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	settings := make(map[string]string)
	f, err := os.Open(settingFile)
	if os.IsNotExist(err) {
		cf, err := os.Create(settingFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		cf.WriteString("{}")
		cf.Close()

		return settings, nil
	}

	if err != nil {
		fmt.Println("other error")
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
	settingFile, err := getSettingFile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bits, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	f, err := os.Create(settingFile)
	if err != nil {
		return err
	}

	_, err = f.Write(bits)
	return err
}

func getSettingFile() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	dir := usr.HomeDir

	return fmt.Sprintf("%s/.bulbs", dir), nil
}
