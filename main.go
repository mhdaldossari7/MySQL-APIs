package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const addr = ":9876"

var dbClient *sqlClient

func main() {
	// initialize db client
	client, err := newSQLClient(sqlConfigs{
		Host:                 "localhost",
		Port:                 3306,
		User:                 "",
		Password:             "",
		DBName:               "",
		DialTimeoutInSeconds: 5,
	})
	if err != nil {
		log.Fatalf("unable to initialize sql client due: %v", err)
	}
	dbClient = client
	// attach handler to server
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/send_user_id", storeUserID).Methods(http.MethodPost)
	router.HandleFunc("/remove_user_id", deleteUserID).Methods(http.MethodPost)
	router.HandleFunc("/send_user_id/{user_id}", sendOneUserID).Methods(http.MethodPost)
	router.HandleFunc("/remove_user_id/{user_id}", removeOneUserID).Methods(http.MethodPost)
	router.HandleFunc("/check_user_id/{user_id}", getUserID).Methods(http.MethodGet)
	router.HandleFunc("/check_user_id", checkIfUserIDExists).Methods(http.MethodGet)
	router.Path("/get_users").Queries("limit", "{limit}").HandlerFunc(getAllID).Name("getAllID")
	//http.HandleFunc("/userID", handleUserID)
	// run server
	log.Printf("server is listening on %v", addr)
	err = http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatalf("unable to run server due: %v", err)
	}
}

// func handleUserID(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
// 	case http.MethodPost:
// 		// handle post new ad
// 		sendUserID(w, r)
// 	case http.MethodDelete:
// 		deleteUserID(w, r)
// 	}
// }

func storeUserID(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp := newBadRequestResp("invalid request body")
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	var userID postUserID
	err = json.Unmarshal(b, &userID)
	if err != nil {
		resp := newBadRequestResp(err.Error())
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	postUserID, err := dbClient.sendUserID(userID)
	if err != nil {
		resp := newErrInternalResp(err.Error())
		writeResp(w, http.StatusInternalServerError, resp)
		return
	}
	writeResp(w, http.StatusOK, newSuccessResp(postUserID, "Added"))
}

func deleteUserID(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp := newBadRequestResp("invalid request body")
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	var userID postUserID
	err = json.Unmarshal(b, &userID)
	if err != nil {
		resp := newBadRequestResp(err.Error())
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	removedUserID, err := dbClient.removeUserID(userID)
	if err != nil {
		resp := newErrInternalResp(err.Error())
		writeResp(w, http.StatusInternalServerError, resp)
		return
	}
	writeResp(w, http.StatusOK, newSuccessResp(removedUserID, "Deleted"))
}

func sendOneUserID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["user_id"]
	var isInt bool
	isInt = isNumeric(id)
	if isInt != true {
		resp := newBadRequestResp("You must pass integers")
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	userID, _ := strconv.Atoi(id)
	m := postUserID{int64(userID)}
	sentUserID, err := dbClient.sendUserID(m)
	if err != nil {
		resp := newErrInternalResp(err.Error())
		writeResp(w, http.StatusInternalServerError, resp)
		return
	}
	writeResp(w, http.StatusOK, newSuccessResp(sentUserID, "Added"))
}

func removeOneUserID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["user_id"]
	var isInt bool
	isInt = isNumeric(id)
	if isInt != true {
		resp := newBadRequestResp("You must pass integers")
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	userID, _ := strconv.Atoi(id)
	m := postUserID{int64(userID)}
	removedUserID, err := dbClient.removeUserID(m)
	if err != nil {
		resp := newErrInternalResp(err.Error())
		writeResp(w, http.StatusInternalServerError, resp)
		return
	}
	writeResp(w, http.StatusOK, newSuccessResp(removedUserID, "Deleted"))
}

func getUserID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["user_id"]
	var isInt bool
	isInt = isNumeric(id)
	if isInt != true {
		resp := newBadRequestResp("You must pass integers")
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	userID, _ := strconv.Atoi(id)
	m := postUserID{int64(userID)}
	getUserIDFromDB, err := dbClient.getUserIDIfExists(m)
	if err != nil {
		resp := newErrInternalResp(err.Error())
		writeResp(w, http.StatusInternalServerError, resp)
		return
	}
	if getUserIDFromDB > 0 {
		writeResp(w, http.StatusOK, newSuccessResp(getUserIDFromDB, "Exists in DB"))
	} else {
		writeResp(w, http.StatusOK, newSuccessRespIfUserIDDoesntExists("Doesn't Exists in DB"))
	}
}

func checkIfUserIDExists(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp := newBadRequestResp("invalid request body")
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	var userID postUserID
	err = json.Unmarshal(b, &userID)
	if err != nil {
		resp := newBadRequestResp(err.Error())
		writeResp(w, http.StatusBadRequest, resp)
		return
	}
	getUserIDFromDB, err := dbClient.getUserIDIfExists(userID)
	if err != nil {
		resp := newErrInternalResp(err.Error())
		writeResp(w, http.StatusInternalServerError, resp)
		return
	}
	if getUserIDFromDB > 0 {
		writeResp(w, http.StatusOK, newSuccessResp(getUserIDFromDB, "Exists in DB"))
	} else {
		writeResp(w, http.StatusOK, newSuccessRespIfUserIDDoesntExists("Doesn't Exists in DB"))
	}
}

func getAllID(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")
	limit2, _ := strconv.Atoi(limit)
	ids, err := dbClient.getAllUsersID(limit2)
	if err != nil {
		resp := newErrInternalResp(err.Error())
		writeResp(w, http.StatusInternalServerError, resp)
		return
	}
	writeResp(w, http.StatusOK, successUsersResp(ids))
}

func writeResp(w http.ResponseWriter, statusCode int, resp resp) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp.JSON())
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
