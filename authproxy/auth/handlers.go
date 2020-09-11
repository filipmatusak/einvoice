package auth

import (
	"encoding/json"
	"net/http"
)

func WithToken(manager UserManager, f func(res http.ResponseWriter, req *http.Request)) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		if manager.Exists(token) {
			req.Header.Del("Authorization")
			f(res, req)
		} else {
			res.WriteHeader(401)
		}
	}
}

type UserInfo struct {
	Token string `json:"token"`
	Id    string `json:"id"`
}

func HandleLogin(manager UserManager) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		user := manager.Create()

		info := UserInfo{user.Token, user.Id}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(info)
	}
}

func HandleLogout(manager UserManager) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		if err := manager.Remove(token); err != nil {
			res.WriteHeader(401)
			return
		}

		res.WriteHeader(200)
	}
}

func HandleMe(manager UserManager) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		token := req.Header.Get("Authorization")
		if user := manager.GetUser(token); user != nil {
			info := UserInfo{user.Token, user.Id}
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(info)
		} else {
			res.WriteHeader(401)
		}
	}
}
