package main

import (
	"encoding/json"
	"fmt"
	"github.com/cossim/hipush/api/http/v1/dto"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var (
	url    = "http://127.0.0.1:7070/api/v1/push"
	method = "POST"
)

func main() {
	payload := dto.PushRequest{
		//AppID:    "xxx",
		AppName:  "cossim",
		Platform: "xiaomi",
		Token: []string{
			"vAmkC65U7IthVLCyK0udEfUiXCiRAP5DXZtV2h0qt/vgo8k0foiQ8YS3VKcbPMa/",
		},
		Data: dto.XiaomiPushRequestData{
			Title:         "cossim",
			Subtitle:      "hello",
			Content:       "hello",
			Icon:          "",
			TTL:           time.Minute,
			NotifyType:    -1,
			ScheduledTime: 0,
			IsScheduled:   false,
			Foreground:    true,
			ClickAction: dto.ClickAction{
				Action: 1,
			},
			Data: nil,
		},
		Option: dto.PushOption{
			DryRun:        false,
			Retry:         1,
			RetryInterval: 1,
		},
	}

	// Marshal the request object to JSON
	reqBody, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(reqBody)))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
