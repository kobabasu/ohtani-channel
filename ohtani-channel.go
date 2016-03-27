package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const verbose = true
const path = "http://live.baseball.yahoo.co.jp/npb/game/"
const irkit = "http://192.168.12.51/messages"

const no = "#11"
const team = "日本ハム"
const next = "大谷翔平"

const channel_cs = `[17421,8755,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,3228,1190,1037,1190,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,1037,1190,3228,1190,3228,1190,1037,1190,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,1037,1190,1037,1190,3228,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,3228,1190]`

const channel_4 = `[17421,8755,1150,1150,1150,1150,1150,1150,1150,1150,1150,1150,1150,1150,1150,3228,1150,1150,1150,3228,1150,3228,1150,3228,1150,3228,1150,3228,1150,3228,1150,1150,1150,3228,1150,1150,1150,1150,1150,3228,1150,1150,1150,1150,1150,1150,1150,1150,1150,1150,1150,3228,1150,3228,1150,1150,1150,3228,1150,3228,1150,3228,1150,3228,1150,3228,1150,65535,0,13693,17421,4400,1150]`

const channel_6 = `[17421,8755,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,3228,1190,1037,1190,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190,1037,1190,3228,1190,1037,1190,3228,1190,3228,1190,1037,1190,1037,1190,1037,1190,1037,1190,1037,1190,3228,1190,1037,1190,1037,1190,3228,1190,3228,1190,3228,1190,3228,1190,3228,1190]`

func query(today string) string {
	data := make(map[string]string)
	js := json.NewDecoder(os.Stdin)
	js.Decode(&data)

	return data[today]
}

func scrape(query string) bool {
	doc, err := goquery.NewDocument(path + query)
	if err != nil {
		log.Fatal(err)
	}

	cNo := doc.Find("#batter .playerNo").Text()
	cTeam := doc.Find(".score .act").Text()
	cNext := doc.Find("#nextR p").Text()

	if verbose {
		fmt.Printf("no: %s\n", cNo)
		fmt.Printf("team: %s\n", cTeam)
		fmt.Printf("next: %s\n", cNext)
	}

	if next == cNext || (team == cTeam && no == cNo) {
		return true
	} else {
		return false
	}
}

func send(ir string) {
	json := `{"format":"raw","freq":38","data":` + ir + `}`
	var ch = bytes.NewBuffer([]byte(json))
	req, err := http.NewRequest("POST", irkit, ch)

	req.Header.Set("X-Requested-With", "curl")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	time.Sleep(1 * time.Second)
	fmt.Println("status:", resp.Status)
}

func main() {
	today := time.Now().Format("20060102")
	// today := "20160326"

	query := query(today)

	if query != "" {
		res := scrape(today + query)
		if res {
			send(channel_cs)
			send(channel_4)
		}
	} else {
		fmt.Println("today has no game")
	}
}
