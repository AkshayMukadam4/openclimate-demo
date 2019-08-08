package server

import (
	// "github.com/pkg/errors"
	"log"
	"net/http"

	// ipfs "github.com/Varunram/essentials/ipfs"
	erpc "github.com/Varunram/essentials/rpc"
	"github.com/YaleOpenLab/openclimate/database"
)

func setupActorsHandlers() {

	getCompanyAssetsByState()

	getAllCompanies()
	getCompany()
	getCompanyStates()
	getCompanyCountries()
	
	getAllRegions()
	getRegion()
	getAllStates()
	getState()
	getStatesByCountry()
	getAllCities()
	getCity()
	getAllCountries()
	getCountry()
}

func getCompanyAssetsByState() {
	http.HandleFunc("/company/assets/filter", func(w http.ResponseWriter, r *http.Request) {

		user, err := CheckGetAuth(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		if user.EntityType != "company" {
			log.Println("User entity type is not a company.")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		company, err := database.RetrieveCompany(user.EntityID)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		if r.URL.Query()["state"] == nil {
			log.Println("State not passed.")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		state := r.URL.Query()["state"][0]
		a, err := company.GetAssetsByState(state)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}
		erpc.MarshalSend(w, a)

		// assetsByState := make(map[string][]database.Asset)
		// for _, state := range states {
		// 	a, err := company.GetAssetsByState(state)
		// 	if err != nil {
		// 		log.Println(err)
		// 		erpc.ResponseHandler(w, erpc.StatusInternalServerError)
		// 		return
		// 	}
		// 	assetsByState[state] = a
		// }

		// erpc.MarshalSend(w, assetsByState)
	})
}




/*******************/
/* REGION HANDLERS */
/*******************/

func getAllRegions() {
	http.HandleFunc("/region/all", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		regions, err := database.RetrieveAllRegions()
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, regions)
	})
}

func getRegion() {
	http.HandleFunc("/region", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		if r.URL.Query()["region_name"] == nil || r.URL.Query()["region_country"] == nil {
			log.Println("Region_name or region_country not passed, quitting")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		name := r.URL.Query()["region_name"][0]
		country := r.URL.Query()["region_country"][0]
		region, err := database.RetrieveRegionByName(name, country) //************ STOP ***********
		if err != nil {
			log.Println("Error while retrieving region, quitting")
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, region)
	})
}

/******************/
/* STATE HANDLERS */
/******************/

func getAllStates() {
	http.HandleFunc("/state/all", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		states, err := database.RetrieveAllStates()
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, states)
	})
}

func getStatesByCountry() {
	http.HandleFunc("/state/filter", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		if r.URL.Query()["country"] == nil {
			log.Println("Country not passed, quitting")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		states, err := database.FilterStates(r.URL.Query()["country"][0])
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, states)
	})
}

func getState() {
	http.HandleFunc("/state", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		if r.URL.Query()["state_name"] == nil || r.URL.Query()["state_country"] == nil {
			log.Println("State_name or state_country not passed, quitting")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		name := r.URL.Query()["state_name"][0]
		country := r.URL.Query()["state_country"][0]
		state, err := database.RetrieveStateByName(name, country) //************ STOP ***********
		if err != nil {
			log.Println("Error while retrieving state, quitting")
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, state)
	})
}

/*****************/
/* CITY HANDLERS */
/*****************/

func getAllCities() {
	http.HandleFunc("/city/all", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		cities, err := database.RetrieveAllCities()
		if err != nil {
			log.Println("Error while retrieving all cities, quitting")
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, cities)
	})
}

func getCity() {
	http.HandleFunc("/city", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		if r.URL.Query()["city_name"] == nil || r.URL.Query()["city_region"] == nil {
			log.Println("City name or city region not passed, quitting")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		name := r.URL.Query()["city_name"][0]
		region := r.URL.Query()["city_region"][0]
		city, err := database.RetrieveCityByName(name, region)
		if err != nil {
			log.Println("Error while retrieving all cities, quitting")
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, city)
	})
}

/********************/
/* COMPANY HANDLERS */
/********************/

func getAllCompanies() {
	http.HandleFunc("/company/all", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		companies, err := database.RetrieveAllCompanies()
		if err != nil {
			log.Println("error while retrieving all companies, quitting")
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, companies)
	})
}

func getCompany() {
	http.HandleFunc("/company", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		if r.URL.Query()["company_name"] == nil || r.URL.Query()["company_country"] == nil {
			log.Println("company name or country not passed, quitting")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		name := r.URL.Query()["company_name"][0]
		country := r.URL.Query()["company_country"][0]
		company, err := database.RetrieveCompanyByName(name, country)
		if err != nil {
			log.Println("error while retrieving all companies, quitting")
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, company)
	})
}

func getCompanyStates() {
	http.HandleFunc("/company/states", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		if r.URL.Query()["company_name"] == nil || r.URL.Query()["company_country"] == nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		// Given its name and country, retrieve the company from the database

		name := r.URL.Query()["company_name"][0]
		country := r.URL.Query()["company_country"][0]
		company, err := database.RetrieveCompanyByName(name, country)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		// Get the information of the states that the company is in

		states, err := company.GetStates()
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, states)
	})
}

func getCompanyCountries() {
	http.HandleFunc("/company/countries", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		if r.URL.Query()["company_name"] == nil || r.URL.Query()["company_country"] == nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		// Given its name and country, retrieve the company from the database

		name := r.URL.Query()["company_name"][0]
		country := r.URL.Query()["company_country"][0]
		company, err := database.RetrieveCompanyByName(name, country)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		// Get the information of the states that the company is in

		countries, err := company.GetCountries()
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, countries)
	})
}

/**********************/
/* COUNTRIES HANDLERS */
/**********************/

func getAllCountries() {
	http.HandleFunc("/country/all", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		countries, err := database.RetrieveAllCountries()
		if err != nil {
			log.Println("error while retrieving all countries, quitting")
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, countries)
	})
}

func getCountry() {
	http.HandleFunc("/country", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			return
		}

		if r.URL.Query()["country_name"] == nil {
			log.Println("country name not passed, quitting")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		name := r.URL.Query()["country_name"][0]
		country, err := database.RetrieveCountryByName(name)
		if err != nil {
			erpc.ResponseHandler(w, erpc.StatusInternalServerError)
			return
		}

		erpc.MarshalSend(w, country)
	})
}
