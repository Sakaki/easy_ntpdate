package easy_ntpdate

import (
	"net/http"
	"fmt"
	"time"
	"io/ioutil"
	"regexp"
	"os/exec"
)

func SetDate() {
	r := regexp.MustCompile("<BODY>\n(.*)\\.")
	for {
		resp, err := http.Get("http://ntp-a1.nict.go.jp/cgi-bin/jst")
		if err != nil {
			fmt.Printf("Connection failed: %s", err)
			time.Sleep(5 * time.Second)
		} else {
			byteArray, _ := ioutil.ReadAll(resp.Body)
			group := r.FindSubmatch(byteArray)
			epochTime := string(group[1])
			fmt.Println(epochTime)
			resp.Body.Close()
			err := exec.Command("sh", "-c", fmt.Sprintf("date  --set @%s", epochTime)).Run()
			if err != nil {
				fmt.Println(err)
			}
			return
		}
	}
}
