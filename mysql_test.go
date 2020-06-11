package main

import (
	"log"
	"testing"
)

var dbClient2 *sqlClient

func TestStoreUserID(t *testing.T) {
	client, err := newSQLClient(sqlConfigs{
		Host:                 "localhost",
		Port:                 3306,
		User:                 "root",
		Password:             "root",
		DBName:               "post_notifications",
		DialTimeoutInSeconds: 5,
	})
	if err != nil {
		log.Fatalf("unable to initialize sql client due: %v", err)
	}
	dbClient2 = client

	numIds := 4
	var ids []postUserID

	for i := 0; i < numIds; i++ {
		ids = append(ids, postUserID{
			UserID: 123,
		})
	}
	for i := 0; i < len(ids); i++ {
		_, err = dbClient2.sendUserID(ids[i])
		if err != nil {
			t.Fatalf("unable to insert to the DB due : %v", err)
		}
	}
	storedID, err := dbClient2.getDocCount()
	if err != nil {
		t.Fatalf("unable to get records from table due : %v", err)
	}
	if storedID != numIds {
		t.Fatalf("got unexpected number of records, expected %v and got %v", numIds, storedID)
	}
	for i := 0; i < len(ids); i++ {
		_, err = dbClient2.removeUserID(ids[i])
		if err != nil {
			t.Fatalf("unable to delete from table due : %v", err)
		}
	}
}

func TestDeleteUserID(t *testing.T) {
	client, err := newSQLClient(sqlConfigs{
		Host:                 "localhost",
		Port:                 3306,
		User:                 "root",
		Password:             "root",
		DBName:               "post_notifications",
		DialTimeoutInSeconds: 5,
	})
	if err != nil {
		log.Fatalf("unable to initialize sql client due: %v", err)
	}
	dbClient2 = client

	sampleUserID := postUserID{
		UserID: 123,
	}
	_, err = dbClient2.sendUserID(sampleUserID)
	if err != nil {
		t.Fatalf("unable to store user id due : %v", err)
	}
	_, err = dbClient2.removeUserID(sampleUserID)
	if err != nil {
		t.Fatalf("unable to delete user id due : %v", err)
	}
	var storedID int
	storedID, err = dbClient2.getDocCount()
	if err != nil {
		t.Fatalf("unable to get records from db due : %v", err)
	}
	if storedID != 0 {
		t.Fatalf("user id still exists in db")
	}
}

func TestGetUserIfExists(t *testing.T) {
	client, err := newSQLClient(sqlConfigs{
		Host:                 "localhost",
		Port:                 3306,
		User:                 "root",
		Password:             "root",
		DBName:               "post_notifications",
		DialTimeoutInSeconds: 5,
	})
	if err != nil {
		log.Fatalf("unable to initialize sql client due: %v", err)
	}
	dbClient2 = client

	sampleUserID := postUserID{
		UserID: 123,
	}
	_, err = dbClient2.sendUserID(sampleUserID)
	if err != nil {
		t.Fatalf("unable to store user in db due : %v", err)
	}
	var postExists int
	postExists, err = dbClient2.getUserIDIfExists(sampleUserID)
	if err != nil {
		t.Fatalf("unable to get data from db due : %v", err)
	}
	if postExists != 1 {
		t.Fatalf("the user id not stored in db")
	}
	_, err = dbClient2.removeUserID(sampleUserID)
	if err != nil {
		t.Fatalf("unable to delete user id due : %v", err)
	}
}

func TestGetLimitedUserID(t *testing.T) {
	client, err := newSQLClient(sqlConfigs{
		Host:                 "localhost",
		Port:                 3306,
		User:                 "root",
		Password:             "root",
		DBName:               "post_notifications",
		DialTimeoutInSeconds: 5,
	})
	if err != nil {
		log.Fatalf("unable to initialize sql client due: %v", err)
	}
	dbClient2 = client

	numIds := 10
	var sampleIds []postUserID
	for i := 0; i < numIds; i++ {
		sampleIds = append(sampleIds, postUserID{
			UserID: 123,
		})
	}
	for i := 0; i < len(sampleIds); i++ {
		_, err = dbClient2.sendUserID(sampleIds[i])
		if err != nil {
			t.Fatalf("unable to store user id in db due : %v", err)
		}
	}
	limit := 5
	var ids []postUserID
	ids, err = dbClient2.getAllUsersID(limit)
	if err != nil {
		t.Fatalf("unable to get limited user ids due : %v", err)
	}
	if len(ids) != 5 {
		t.Fatalf("got unexpected number of results, expected %v and got %v", limit, len(ids))
	}
}
