package config

import (
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"encoding/xml"

	log "github.com/sirupsen/logrus"

	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func parseYaml(yamlPath string, out interface{}) error {
	//yaml.UnmarshalStrict
	yamlFile, err := ioutil.ReadFile(yamlPath)
	if err != nil {
		log.Printf("Read yaml file error %v #%v ", yamlPath, err)
		return err
	}

	dErr := yaml.Unmarshal(yamlFile, out)
	if dErr != nil {
		log.Printf("Decode yaml file error %v #%v ", yamlPath, dErr)
		return dErr
	}
	return nil
}

func loadYamlFromURL(configURL string, out interface{}) error {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(configURL)
	if err != nil {
		// handle error
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("http: read file failed")
	}
	body, err := ioutil.ReadAll(resp.Body)
	dErr := yaml.Unmarshal(body, out)
	if dErr != nil {
		log.Printf("Decode yaml file error %v #%v ", configURL, dErr)
		return dErr
	}
	return nil
}

// LoadYamlFromURL load remote
func LoadYamlFromURL(configURL string, out interface{}) error {
	err := loadYamlFromURL(configURL, out)
	if err != nil {
		tryTime := 30
		for tryTime > 0 && err != nil {
			tryTime = tryTime - 1
			time.Sleep(time.Second * 5)
			err = loadYamlFromURL(configURL, out)
			if err != nil {
				log.Errorf("load config: %v err %v, retry...%v", configURL, err, tryTime)
			}
		}
		if err != nil {
			log.Errorf("load config give up %v", err)
		}
	}
	return err
}

// LoadXMLFromURL load remote
func LoadXMLFromURL(configURL string, out interface{}) error {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(configURL)
	if err != nil {
		// handle error
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("http: read file failed")
	}
	body, err := ioutil.ReadAll(resp.Body)
	dErr := xml.Unmarshal(body, out)
	if dErr != nil {
		log.Printf("Decode yaml file error %v #%v ", configURL, dErr)
		return dErr
	}
	return nil
}

// LoadYaml load config
func LoadYaml(out interface{}) error {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	configPath := ""
	configFile := ""
	flag.StringVar(&configPath, "config-path", "", "Config file path")
	flag.StringVar(&configFile, "config-file", "", "Config file path")
	configURL := ""
	flag.StringVar(&configURL, "config-url", "", "Config file configURL")
	flag.Parse()
	if len(configURL) > 0 {
		log.Printf("Loading config file from URL: %v", configURL)
		return LoadYamlFromURL(configURL, out)
	}

	if len(configPath) > 0 {
		log.Printf("Loading config file from %v", configPath)
		return parseYaml(configPath, out)
	}

	if len(configFile) > 0 {
		log.Printf("Loading config file from %v", configFile)
		return parseYaml(configFile, out)
	}
	//
	configURLInEnv := os.Getenv("CONFIG_URL")
	if len(configURLInEnv) > 0 {
		log.Printf("Loading config file from URL: %v", configURLInEnv)
		return LoadYamlFromURL(configURLInEnv, out)
	}
	configPathInEnv := os.Getenv("CONFIG_PATH")
	if len(configPathInEnv) > 0 {
		configPath = configPathInEnv
	}
	if len(configPath) < 1 {
		//
		log.Println("Config file path not set. Trying to use default config file.")
		// read system to load file
		fPath, err := loadDefaultConfigFile()
		if err != nil {
			log.Println("Config file path load error.")
			return err
		}
		configPath = fPath
	}
	log.Printf("Loading config file from %v", configPath)
	err := parseYaml(configPath, out)

	return err
}

func loadDefaultConfigFile() (string, error) {
	s, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	i := strings.LastIndex(s, string(os.PathSeparator))
	path := string(s[0:i+1]) + "default-config.yml"
	return path, nil
}
