// commands
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var lastchat int64 = 0

func CmdInterpreter(channel string, username string, usermessage string) {
	if strings.Contains(usermessage, nick) || strings.Contains(usermessage, strings.ToLower(nick)) {
		msg := strings.Replace(usermessage, nick, "", 1)
		if msg != "" && lastchat+10 <= time.Now().Unix() {
			Message(channel, askOracle(username, msg))
			lastchat = time.Now().Unix()
		}
	}
}

func askOracle(username string, message string) string {
	pohja := "http://www.lintukoto.net/viihde/oraakkeli/index.php?kysymys="
	url := pohja + url.QueryEscape(message) + "&html"
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Cannot get URL response: %s\n", err.Error())
	} else {
		defer response.Body.Close()
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatalf("Cannot read URL response: %s\n", err.Error())
		}
		answer := strings.Replace(toUtf8(contents), "\n", "", 1)
		result := fmt.Sprintf("%s: %s", username, answer)
		return result
	}
	return "Kappa"
}

func toUtf8(iso8859_1_buf []byte) string {
	buf := make([]rune, len(iso8859_1_buf))
	for i, b := range iso8859_1_buf {
		buf[i] = rune(b)
	}
	return string(buf)
}
