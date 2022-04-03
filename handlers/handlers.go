package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/andrewyan200/couriers/cloud_db"
	"github.com/microcosm-cc/bluemonday"
	"net/http"

)

type Courier struct {
	Name string `json: "full_name"`
	City string `json: "city"`
	WorkHours string `json: "workHours"`
}

func GetCourierHandler(w http.ResponseWriter, r * http.Request) {
	//fetch couriers from cloud database
	couriers := cloud_db.Query()
	courierListBytes, err := json.Marshal(couriers)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	// If all goes well, write the JSON list of couriers to the response
	w.Write(courierListBytes)

}

func CreateCourierHandler(w http.ResponseWriter, r *http.Request) {
	// Create a new instance Courier
	courier := Courier{}

	// We send all our data as HTML form data
	// the `ParseForm` method of the request, parses the
	// form values
	err := r.ParseForm()

	// In case of any error, we respond with an error to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Get the information about the courier from the form info and sanitize it 
	p := bluemonday.StrictPolicy() // here we use the default policy StrictPolicy
	courier.Name = p.Sanitize(r.Form.Get("full_name"))
	courier.City = p.Sanitize(r.Form.Get("city"))
	courier.WorkHours = p.Sanitize(r.Form.Get("workHours"))

	//Post to the cloud database the received info
	cloud_db.Post_request(p.Sanitize(courier.Name), p.Sanitize(courier.City), p.Sanitize(courier.WorkHours))

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}