package main

import (
	"os"
	"path"
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

    save2md(allResults)
}
