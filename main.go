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

	// go listen(env)

	cred := loadCredential()
	if cred != nil {
		mustRefreshToken(env, cred.Token)
	} else {
		cred = newCredential(env)
	}
	go watchToken(env, cred)

	initFacebookSession(env, cred)

}
