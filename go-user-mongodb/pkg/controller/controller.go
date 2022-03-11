package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-user-mongodb/pkg/models"
	"github.com/go-user-mongodb/pkg/utils"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	if !bson.IsObjectIdHex(params["id"]) {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	UserDetails, _ := models.GetUser(params["id"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UserDetails)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	fmt.Println(CreateUser)
	b := CreateUser.CreateUser()
	res, _ := json.Marshal(b)
	fmt.Println(string(res))
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
