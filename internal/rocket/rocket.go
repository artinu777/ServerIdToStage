package rocket

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Rocket struct {
	Cfg Config
}
var (
	config Config
	client *http.Client
	)


type Config struct {
	Url string `yaml:"Url"`
}
type Massage struct {
	Text        string `json:"text"`
	Attachments []struct {
		Title     string `json:"title"`
		TitleLink string `json:"title_link"`
		Text      string `json:"text"`
		ImageURL  string `json:"image_url"`
		Color     string `json:"color"`
	} `json:"attachments"`
}


func Init(cfg Config)  {
	config = cfg
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
}

func Send(message Massage)  {

	fmt.Println(config.Url)

	var data []byte
	var err error

	if data, err = json.Marshal(message); err != nil {
		return
	}
	req, err := http.NewRequest("POST", config.Url, bytes.NewBuffer(data))
	resp, err := client.Do(req)
	if err != nil{
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	var deletehostresponse interface{}
	jsonErr := json.Unmarshal(bodyText, &deletehostresponse)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

}