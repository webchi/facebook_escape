package main

import (
	"encoding/json"
	"github.com/huandu/facebook/v2"
	"log"
	"os"
	"time"
)

type (
	Feed struct {
		Data []struct {
			ID          string `json:"id"`
			ObjectID    string `json:"object_id"`
			Message     string `json:"message"`
			Application struct {
				ID        string `json:"id"`
				Name      string `json:"name"`
				NameSpace string `json:"namespace"`
				Category  string `json:"category"`
				Link      string `json:"link"`
			} `json:"application"`
		} `json:"data"`
		Paging struct {
			Previous string `json:"previous"`
			Next     string `json:"next"`
		} `json:"paging"`
	}
	Image struct {
		ID     string `json:"id"`
		Name     string `json:"name"`
		Images []struct {
			Width  int    `json:"width"`
			Height int    `json:"height"`
			Source string `json:"source"`
		} `json:"images"`
	}
	Result struct {
		Message  string
	}
)

var sess *facebook.Session

func initFacebookSession(env *Env, cred *Credential) {
	app := facebook.New(env.AppID, env.AppSecret)
	sess = app.Session(cred.Token)
}

func mustRefreshToken(env *Env, token string) *Credential {
	app := facebook.New(env.AppID, env.AppSecret)
	tmpSess := app.Session(token)
	res, err := tmpSess.Get("/oauth/access_token", M{
		"grant_type":        "fb_exchange_token",
		"client_id":         env.AppID,
		"client_secret":     env.AppSecret,
		"fb_exchange_token": token,
	})
	if err != nil {
		log.Println("ロングターム トークンの取得に失敗しました:", err)
		os.Exit(3)
	}
	expiresIn, _ := res.Get("expires_in").(json.Number).Int64()
	expireAt := time.Now().Add(time.Duration(expiresIn) * time.Second)
	cred := &Credential{
		ExpireAt: expireAt,
		Token:    res.Get("access_token").(string),
	}
	saveCredential(cred)
	return cred
}

func watchToken(env *Env, cred *Credential) {
	for {
		diff := cred.ExpireAt.Sub(time.Now())
		if diff.Hours() < 24*7 {
			cred = mustRefreshToken(env, cred.Token)
			saveCredential(cred)
			initFacebookSession(env, cred)
		}
		time.Sleep(time.Minute)
	}
}
