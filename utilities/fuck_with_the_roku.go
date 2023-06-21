package utilities

import (
	"net/http"
	"net/url"
	"strings"
	"time"
)

func GetIPNameLike(partialName string) (ip string) {
	return ""
}

var roku_url = "192.168.1.158"

func TypeLetter(str string) {
	// makes a curl request to the roku, in Go
	// ex. curl -d '' "http://$ROKU_DEV_TARGET:8060/keypress/Lit_r"

	urlEncodedStr := url.QueryEscape(str)

	http.Post("http://"+roku_url+":8060/keypress/Lit_"+urlEncodedStr, "", nil)
}

func SpellWord(word string) int {
	letters := strings.SplitN(word, "", -1)
	for _, letter := range letters {
		TypeLetter(letter)
		time.Sleep(time.Millisecond * 150)
	}

	return len(letters)
}

func Erase(num int) {
	for i := 0; i < num; i++ {
		http.Post("http://"+roku_url+":8060/keypress/backspace", "", nil)
	}
}

func CreepOnRoku(script []string) {
	for _, line := range script {
		SpellWord(line)
		dur := len(line)/15 + 2
		time.Sleep(time.Second * time.Duration(dur))
		Erase(len(line))
		time.Sleep(time.Second * 1)
	}
}

func LaunchNetFlix() {
	http.Post("http://"+roku_url+":8060/launch/12", "", nil)
}

func Control(key string) {
	http.Post("http://"+roku_url+":8060/keypress/"+key, "", nil)
}

func GoToNetflixAndCreepOnRoku(script []string) {
	LaunchNetFlix()
	time.Sleep(time.Second * 5)

	// Navigate to search from home
	Control("left")
	Control("left")
	Control("left")
	Control("left")
	Control("left")
	Control("left")
	Control("up")
	Control("select")

	CreepOnRoku(script)
}
