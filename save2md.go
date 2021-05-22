package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"github.com/huandu/facebook/v2"
)

func save2md(results []facebook.Result){
	for _, data := range results {
		// loop all posts
		if data["message"] == nil {
			// check is post have message (or only pic)
			continue
		}
		timestamp, _ := time.Parse("2006-01-02T15:04:05-0700", data["created_time"].(string)) // parse post timestamp
		folder := "data/" + strings.Split(data["id"].(string), "_")[0] + "/" + timestamp.Format("2006/01")
		err := os.MkdirAll(folder, 0755)
		if err != nil {
			fmt.Println(err)
			continue
		}
		file, err := os.Create(folder + "/" + timestamp.Format("02-15-04-05") + "-" + regexp.MustCompile("[A-zА-я]+").FindString(data["message"].(string))) // open file
		if err != nil {
			fmt.Println(err)
			continue
		}
		_, err = file.WriteString(data["message"].(string))
		if err != nil {
			fmt.Println(err)
			continue
		}
		err = file.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}