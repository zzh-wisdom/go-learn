package main
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)
func main() {
	resp, err := http.PostForm("http://www.baidu.com", url.Values{"wd": {"golang"}})
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}