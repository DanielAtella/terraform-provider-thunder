package go_thunder

import (
	"bytes"
	"fmt"
	"github.com/clarketm/json" // "encoding/json"
	"io/ioutil"
	"util"
)

type HTTP2 struct {
	UUID HTTP2Instance `json:"http2,omitempty"`
}

type SamplingEnable34 struct {
	Counters1 string `json:"counters1,omitempty"`
}
type HTTP2Instance struct {
	Counters1 []SamplingEnable34 `json:"sampling-enable,omitempty"`
	UUID      string             `json:"uuid,omitempty"`
}

func GetSlbHTTP2(id string, host string) (*HTTP2, error) {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	resp, err := DoHttp("GET", "https://"+host+"/axapi/v3/slb/http2", nil, headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
		return nil, err
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var m HTTP2
		erro := json.Unmarshal(data, &m)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
			return nil, err
		} else {
			fmt.Print(m)
			logger.Println("[INFO] GET REQ RES..........................", m)
			err := check_api_status("GetSlbHTTP2", data)
			if err != nil {
				return nil, err
			}
			return &m, nil
		}
	}
}

func PostSlbHTTP2(id string, vc HTTP2, host string) error {

	logger := util.GetLoggerInstance()

	var headers = make(map[string]string)
	headers["Accept"] = "application/json"
	headers["Content-Type"] = "application/json"
	headers["Authorization"] = id

	payloadBytes, err := json.Marshal(vc)

	logger.Println("[INFO] input payload bytes - " + string((payloadBytes)))

	if err != nil {
		logger.Println("[INFO] Marshalling failed with error \n", err)
	}
	resp, err := DoHttp("POST", "https://"+host+"/axapi/v3/slb/http2", bytes.NewReader(payloadBytes), headers)

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		logger.Println("The HTTP request failed with error \n", err)
	} else {
		data, _ := ioutil.ReadAll(resp.Body)
		var v HTTP2
		erro := json.Unmarshal(data, &v)
		if erro != nil {
			fmt.Printf("Unmarshal error %s\n", err)
		} else {
			fmt.Println("response Body:", string(data))
			logger.Println("response Body:", string(data))
			err := check_api_status("PostSlbHTTP2", data)
			if err != nil {
				return err
			}
		}
	}
	return err
}
