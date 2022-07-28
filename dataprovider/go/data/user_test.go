package data

import (
	"github.com/google/uuid"
	"github.com/tnyidea/react-admin-dataprovider/go/types"
	"log"
	"reflect"
	"testing"
)

func TestFindAllUsers(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	users, err := db.FindAllUsers()

	log.Println(users)
}

func TestFindUserByUUID(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	user, err := db.FindUserByUUID("8e78fa20-6050-4236-9f14-04e232fd04ed")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	log.Println(&user)
}

func TestCreateUserWithUUID(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	uuidString := uuid.NewString()
	newUser := types.User{
		UUID:        uuidString,
		FirstName:   "John",
		LastName:    "Smith",
		CompanyName: "Smith and Associates",
		Address:     "123 Main Street",
		City:        "Springfield",
		County:      "Springfield",
		State:       "DE",
		Zip:         "12345",
		Phone1:      "555-555-5555",
		Phone2:      "555-555-5556",
		Email:       "john.smith@email.com",
		Web:         "http://johnsmith.com",
	}
	err = db.CreateUser(newUser)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	user, err := db.FindUserByUUID(uuidString)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if !reflect.DeepEqual(user, newUser) {
		log.Println("error: user != newUser")
		t.FailNow()
	}

	log.Println(&user)
}

func TestUpdateUser(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	user, err := db.FindUserByUUID("8e78fa20-6050-4236-9f14-04e232fd04ed")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	user.FirstName = "Jane"
	err = db.UpdateUser(user)
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	user, err = db.FindUserByUUID("8e78fa20-6050-4236-9f14-04e232fd04ed")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}
	if user.FirstName != "Jane" {
		log.Println("error: update failed. Incorrect user.Firstname value:", user.FirstName)
	}

	log.Println(&user)
}

func TestDeleteAllUsers(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	err = db.DeleteAllUsers()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	users, err := db.FindAllUsers()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	if len(users) != 0 {
		log.Println("error: expected db.FindAllUsers() to return empty slice")
		t.FailNow()
	}
}

func TestDeleteUser(t *testing.T) {
	db, err := NewUserDatabase()
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	err = db.DeleteUserByUUID("8e78fa20-6050-4236-9f14-04e232fd04ed")
	if err != nil {
		log.Println(err)
		t.FailNow()
	}

	_, err = db.FindUserByUUID("8e78fa20-6050-4236-9f14-04e232fd04ed")
	if err == nil {
		log.Println("error: db.DeleteUserByUUID() failed: expected err result from db.FindUserByUUID()")
		t.FailNow()
	}
}
