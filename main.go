package main

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

type (
	M = map[string]interface{}
)

func main() {
	_ = os.Mkdir(path.Join(getDir(), "config"), 0700)

	env := mustEnv()

	cred := loadCredential()
	if cred != nil {
		mustRefreshToken(env, cred.Token)
	} else {
		cred = newCredential(env)
	}
	go watchToken(env, cred)

	initFacebookSession(env, cred)

	res, _ := sess.Get("/me/posts", nil)

	paging, _ := res.Paging(sess)

	// append first page of results to slice of Result
	allResults = append(allResults, paging.Data()...)

	for {
		// get next page.
		noMore, err := paging.Next()
		if err != nil {
			panic(err)
		}
		if noMore {
			// No more results available
			break
		}
		// append current page of results to slice of Result
		allResults = append(allResults, paging.Data()...)
	}

	for _, data := range allResults {
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
		file, err := os.Create(folder + "/" + timestamp.Format("02-15-04-05") + "-" + regexp.MustCompile("\\s").Split(data["message"].(string),2)[0]) // open file
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
