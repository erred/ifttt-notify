package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const baseURL string = "https://maker.ifttt.com/trigger/%s/with/key/%s"
const contentType = "application/json"

type data struct {
	Value1 string `json:"value1,omitempty"`
	Value2 string `json:"value2,omitempty"`
	Value3 string `json:"value3,omitempty"`
}

func main() {
	event := flag.String("event", os.Getenv("IFTTT_EVENT"), "event to trigger, defaults to env IFTTT_EVENT")
	key := flag.String("key", os.Getenv("IFTTT_KEY"), "key for event, defaults to env IFTTT_KEY")
	flag.Parse()
	url := fmt.Sprintf(baseURL, *event, *key)

	d := data{}
	switch len(flag.Args()) {
	case 5:
		d.Value3 = flag.Arg(3)
		fallthrough
	case 4:
		d.Value2 = flag.Arg(2)
		fallthrough
	case 3:
		d.Value1 = flag.Arg(1)
	default:
		log.Fatal("need 1-3 args, found: ", len(flag.Args()))
	}

	b, err := json.Marshal(d)
	if err != nil {
		log.Fatal("error marshaling data: ", err)
	}

	_, err = http.Post(url, contentType, bytes.NewBuffer(b))
	if err != nil {
		log.Fatal("error posting data: ", err)
	}
}
