package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func main() {
	flag.Parse()

	var configFile string
	if flag.Arg(0) == "" {
		fmt.Println("Input filename")
		fmt.Scan(&configFile)
	} else {
		configFile = flag.Arg(0)
	}

	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	configArr := strings.Split(string(b), "\n")
	for e := 0; e <= len(configArr)-1; e++ {
		ele := configArr[e]
		if !strings.HasPrefix(ele, "[") && !strings.HasPrefix(ele, "#") && !strings.HasPrefix(ele, ";") {
			envDict := regexp.MustCompile("[=:]").Split(ele, -1)
			if len(envDict) == 2 {
				key := strings.Trim(envDict[0], " ")
				value := strings.Trim(envDict[1], " ")
				os.Setenv(key, value)
				fmt.Printf("Set environment variable: %v=%v\n", key, value)
			}
		}
	}
}
