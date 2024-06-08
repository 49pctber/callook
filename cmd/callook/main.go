package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/49pctber/callook"
)

func main() {
	flag.Parse()
	callsign := strings.ToUpper(flag.Arg(0))
	if len(callsign) < 4 {
		log.Fatal("invalid callsign")
	}

	info := callook.GetCallsignInfoText(callsign)
	fmt.Println(info)
}
