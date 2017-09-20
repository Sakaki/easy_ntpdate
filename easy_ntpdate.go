package main

import (
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"regexp"
	"os/exec"
)

func main() {
	r := regexp.MustCompile("<BODY>\n(.*)\\.")
	for {
		resp, err := http.Get("http://ntp-a1.nict.go.jp/cgi-bin/jst")
		if err != nil {
			fmt.Printf("Connection failed: %s", err)
			time.Sleep(5 * time.Second)
		} else {
			byteArray, _ := ioutil.ReadAll(resp.Body)
			group := r.FindSubmatch(byteArray)
			epoc_time := string(group[1])
			fmt.Println(epoc_time)
			resp.Body.Close()
			err := exec.Command("sh", "-c", fmt.Sprintf("date  --set @%s", epoc_time)).Run()
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}
