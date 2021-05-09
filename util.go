package main

import (
	"os"
	"path"
)

type (
	Env struct {
		AppID      string
		AppSecret  string
		BaseURL    string
		ListenAddr string
	}
)

func mustEnv() *Env {
	appID := os.Getenv("FACEBOOK_APP_ID")
	appSecret := os.Getenv("FACEBOOK_APP_SECRET")
	baseURL := os.Getenv("BASE_URL")
	listenAddr := os.Getenv("LISTEN_ADDR")	
	if listenAddr == "" {
		listenAddr = ":8080"
	}
	if baseURL == "" {
		baseURL = "http://localhost:80"
	}
	if appID == "" || appSecret == "" {
		panic("Didn't defined 'FACEBOOK_APP_ID', 'FACEBOOK_APP_SECRET'")
	}
	return &Env{
		AppID:      appID,
		AppSecret:  appSecret,
		BaseURL:    baseURL,
		ListenAddr: listenAddr,		
	}
}

func getDir() string {
	return path.Dir(os.Args[0])
}
