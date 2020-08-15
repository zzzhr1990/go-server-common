package httprpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"errors"
	// log "github.com/sirupsen/logrus"
)

//PostJSONAndUnMarshal post json
func PostJSONAndUnMarshal(url string, postData interface{}, recvData interface{}, errorCallback func(string)) error {
	data, err := json.Marshal(postData)
	if err != nil {
		if errorCallback != nil {
			errorCallback(fmt.Sprintf("cannot Marshal json %v", err))
		}
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	// req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if errorCallback != nil {
			errorCallback(fmt.Sprintf("Http API request err %v", err))
		}
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if errorCallback != nil {
			errorCallback(fmt.Sprintf("decode Http body err %v", err))
		}
		return err
	}
	err = json.Unmarshal(responseBody, recvData)
	if err != nil {
		if errorCallback != nil {
			errorCallback(fmt.Sprintf("Http Api request Unmarshal error %v, recv: %v", err, string(responseBody)))
		}
		return err
	}
	return nil
}

//GetJSONAndUnMarshal post json
func GetJSONAndUnMarshal(url string, recvData interface{}, errorCallback func(string)) error {
	req, err := http.NewRequest("GET", url, nil)
	// req.Header.Set("X-Custom-Header", "myvalue")
	// req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		if errorCallback != nil {
			errorCallback(fmt.Sprintf("Http Api request err %v", err))
		}
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		if errorCallback != nil {
			errorCallback(fmt.Sprintf("decode Http body err %v", err))
		}
		return err
	}
	err = json.Unmarshal(responseBody, recvData)
	if err != nil {
		if errorCallback != nil {
			errorCallback(fmt.Sprintf("Http api request Unmarshal error %v, recv: %v", err, string(responseBody)))
		}
		return err
	}
	return nil
}

// TryPostJSONAndUnMarshal post json
func TryPostJSONAndUnMarshal(url string, postData interface{}, recvData interface{}, tryTime int, errorCallback func(string)) error {
	//success := false
	for {
		err := PostJSONAndUnMarshal(url, postData, recvData, errorCallback)
		if err == nil {
			return nil
		}
		//
		tryTime--
		if tryTime < 1 {
			return err
		}
		if errorCallback != nil {
			if err != nil {
				errorCallback(fmt.Sprintf("request API %v error %v, retrying", url, err))
			} else {
				errorCallback(fmt.Sprintf("request API %v error, retrying", url))
			}
		}

		time.Sleep(2 * time.Second)
	}
}

//TryPostJSONAndUnMarshalStandard post json
func TryPostJSONAndUnMarshalStandard(url string, postData interface{}, recvData interface{}, tryTime int, errorCallback func(string)) error {
	//success := false
	for {
		cs := &StandardResponseEntity{}
		err := PostJSONAndUnMarshal(url, postData, cs, errorCallback)
		if err == nil && cs.Success {
			if recvData != nil {
				err = cs.DecodeRaw(recvData)
				if err != nil {
					if errorCallback != nil {
						errorCallback(fmt.Sprintf("decode inner json failed %v", err))
					}
					return err
				}
			}
			return nil
		}
		if err == nil {
			if errorCallback != nil {
				errorCallback(fmt.Sprintf("access api error %v:%v", url, cs.Message))
			}
			return errors.New(cs.Message)
		}
		tryTime--
		if tryTime < 1 {
			return err
		}
		errorCallback(fmt.Sprintf("request API %v error %v, retrying", url, err))
		time.Sleep(2 * time.Second)
	}
}
