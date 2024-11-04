package models

import (
	et "authRestApis/models/entities"
	"sync"
	"time"
)


type Tokens struct{
	sync.Mutex
	TokenMap map[string]time.Time
}

var (
	UserData map[string]*et.User
	BlacklistedTokens Tokens
)

func InitializeDb(){
	UserData = map[string]*et.User{}
	BlacklistedTokens = Tokens{TokenMap: map[string]time.Time{}}
}