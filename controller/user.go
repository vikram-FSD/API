package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/atomedgesoft/scheduler/inputvalidator"
	"github.com/atomedgesoft/scheduler/model"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	// inputvalidator.IsMethodValid(w, r, "POST")
	w.Header().Set("Content-Type", "application/json")
	inputvalidator.IsMethodValid(w, r, "POST")

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Print(err)
	}
	//validating input if it is empty
	if user.FirstName == "" {
		http.Error(w, "FirstName is required", http.StatusBadRequest)
		return
	}
	if user.LastName == "" {
		http.Error(w, "LastName is required", http.StatusBadRequest)
		return
	}
	if user.EmailAddress == "" {
		http.Error(w, "EmailId is required", http.StatusBadRequest)
		return
	}
	if user.Signinthrough == "" {
		http.Error(w, "SignInThrough is required", http.StatusBadRequest)
		return
	}
	if user.TimeZone == "" {
		http.Error(w, "Timezone is required", http.StatusBadRequest)
		return
	}
	//Validating emailAddress and send to model to store with DB
	email := user.EmailAddress
	isEmailValid := inputvalidator.IsEmailAddressValid("English", email, 100)
	if isEmailValid == "valid" {
		user.EmailAddress = email
	} else {
		http.Error(w, "Email ID is invalid", http.StatusBadRequest)
		log.Fatal("Email ID is Invalid Please enter the valid one!")
	}
	//Autogenerating Id with length of 12 character & only contains letters & numbers and send to model to store with DB.
	user.Id = inputvalidator.GenerateRandomKey("generateID12345")
	//validating country string with space
	Country := user.Country
	res, err := inputvalidator.IsStringWitSpaceValid("EN", Country, 50, false, "country")
	if res == "valid" {
		user.Country = Country
	} else {
		log.Fatal(err)
	}

	user.CreatedAt = inputvalidator.Timenow(user.CreatedAt.Local().Location())
	response := model.InsertUser(user)
	user.IsActive = true
	fmt.Printf("\n IsActive: %v", user.IsActive)
	fmt.Println("\n",response, "inserted successfully !")
	fmt.Fprint(w, `inserted`, "\t", response)

}

// func GetUser() {

// }
