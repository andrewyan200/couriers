package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

)

var couriers []Courier

func getCourierHandlerLocal(w http.ResponseWriter, r * http.Request) {
	//convert courier variable to json
	courierListBytes, err := json.Marshal(couriers)

	//if error, print to console + display server error
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of couriers to the response
	w.Write(courierListBytes)

}

func createCourierHandlerLocal(w http.ResponseWriter, r *http.Request) {
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

	// Get the information about the courier from the form info
	courier.Name = r.Form.Get("full_name")
	courier.City = r.Form.Get("city")
	courier.WorkHours = r.Form.Get("workHours")

	// Append our existing list of couriers with a new entry
	couriers = append(couriers, courier)

	//Finally, we redirect the user to the original HTMl page
	// (located at `/assets/`), using the http libraries `Redirect` method
	http.Redirect(w, r, "/assets/", http.StatusFound)
}