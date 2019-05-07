package httprpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

//PostJSONAndUnMarshal post json
func PostJSONAndUnMarshal(url string, postData interface{}, recvData interface{}) error {
	data, err := json.Marshal(postData)
	if err != nil {
		log.Errorf("Cannot Marshal json %v", err)
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	// req.Header.Set("X-Custom-Header", "myvalue")
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
		log.Errorf("Request API %v error %v, retrying", url, err)
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
		// if cs.Success {}
		//
		tryTime--
		if tryTime < 1 {
			return err
		}
		log.Errorf("Request API %v error %v, retrying", url, err)
		time.Sleep(2 * time.Second)
	}
}
