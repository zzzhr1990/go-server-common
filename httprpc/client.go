package httprpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"errors"

	log "github.com/sirupsen/logrus"
)

//PostJSONAndUnMarshal post json
func PostJSONAndUnMarshal(url string, postData interface{}, recvData interface{}) error {
	data, err := json.Marshal(postData)
	if err != nil {
		log.Errorf("Cannot Marshal json %v", err)
		return err
	}
	log.Info("------------------QUERY----------------")
	log.Infof("POST: %v", url)
	log.Infof(string(data))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	// req.Header.Set("X-Custom-Header", "myvalue")
	log.Info("------------------QUERY----------------")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Http Api request err %v", err)
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("Decode Http body err %v", err)
		return err
	}
	err = json.Unmarshal(responseBody, recvData)
	log.Info("------------------RECV----------------")
	log.Infof("recv: %v", string(responseBody))
	log.Info("======================================")
	if err != nil {
		log.Errorf("Http Api request Unmarshal error %v", err)
		return err
	}
	return nil
}

//TryPostJSONAndUnMarshal post json
func TryPostJSONAndUnMarshal(url string, postData interface{}, recvData interface{}, tryTime int) error {
	//success := false
	for {
		err := PostJSONAndUnMarshal(url, postData, recvData)
		if err == nil {
			return nil
		}
		//
		tryTime--
		if tryTime < 1 {
			return err
		}
		if err != nil {
			log.Errorf("Request API %v error %v, retrying", url, err)
		} else {
			log.Errorf("Request API %v error, retrying", url)
		}

		time.Sleep(2 * time.Second)
	}
}

//TryPostJSONAndUnMarshalStandard post json
func TryPostJSONAndUnMarshalStandard(url string, postData interface{}, recvData interface{}, tryTime int) error {
	//success := false
	for {
		cs := &StandardResponseEntity{}
		err := PostJSONAndUnMarshal(url, postData, cs)
		if err == nil && cs.Success {
			if recvData != nil {
				err = cs.DecodeRaw(recvData)
				if err != nil {
					log.Errorf("decode inner json failed %v", err)
					return err
				}
			}
			return nil
		}
		if err == nil {
			log.Errorf("access api error %v:%v", url, cs.Message)
			return errors.New(cs.Message)
		}
		tryTime--
		if tryTime < 1 {
			return err
		}
		log.Errorf("Request API %v error %v, retrying", url, err)
		time.Sleep(2 * time.Second)
	}
}
