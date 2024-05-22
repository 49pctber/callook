package callook

import (
	"io"
	"log"
	"net/http"
	"strings"
)

func GetCallsignInfoText(callsign string) string {
	url := "https://callook.info/" + callsign + "/text"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return buf.String()
}
