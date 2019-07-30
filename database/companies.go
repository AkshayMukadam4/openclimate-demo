package database

import (
	"encoding/json"
	edb "github.com/Varunram/essentials/database"
	globals "github.com/YaleOpenLab/openclimate/globals"
	"github.com/pkg/errors"
	"log"
)

// Our definition of "Company" includes ....
// The following struct defines the relevant fields.
type Company struct {

	// Identifying info
	Index   int
	Name    string
	Country string

	UserIDs []int

	// Contextual data
	Area        float64
	Iso         string
	Population  int
	Latitude    float64
	Longitude   float64
	Revenue     float64
	CompanySize int
	HQ          string

	MultiNational []string // an array of all the countries a company is in; if not an MNC, leave empty
	ForProfit     bool
	Industry      bool

	MRV 		string // the company's chosen MRV reporting methodology

	Pledges []Pledge

	// The entity IDs of all the company's physical assets
	Children []int

	// IDs of all the company's financial/regulatory assets (e.g. RECs, climate bonds, etc.)
	Assets []int

	// Data that is reported (through self-reporting, databases, IoT, etc.)
	// as opposed to data that is aggregated from its parts/children. Data
	// is stored on IPFS, so Reports holds the IPFS hashes.
	Reports []RepData

	Emissions  map[string]string // accept whatever emissions the frontend passes
	Mitigation map[string]string
	Adaptation map[string]string
}

// Function that creates a new company object given its name
// and country and saves the object in the countries bucket.
func NewCompany(name string, country string) (Company, error) {
	var company Company

	companies, err := RetrieveAllCompanies()
	if err != nil {
		return company, errors.Wrap(err, "could not retrieve all companies, quitting")
	}

	if len(companies) == 0 {
		company.Index = 1
	} else {
		company.Index = len(companies) + 1
	}

	company.Name = name
	company.Country = country
	return company, company.Save()
}

// Saves company object in companies bucket. Called by NewCompany
func (a *Company) Save() error {
	return edb.Save(globals.DbPath, CompanyBucket, a, a.Index)
}

// Given a key of type int, retrieves the corresponding company object
// from the database companies bucket.
func RetrieveCompany(key int) (Company, error) {
	var company Company
	companyBytes, err := edb.Retrieve(globals.DbPath, CompanyBucket, key)
	if err != nil {
		return company, errors.Wrap(err, "error while retrieving key from bucket")
	}
	err = json.Unmarshal(companyBytes, &company)
	return company, err
}

// Given a name and country, retrieves the corresponding company object
// from the database companies bucket.
func RetrieveCompanyByName(name string, country string) (Company, error) {
	var company Company
	temp, err := RetrieveAllCompanies()
	if err != nil {
		return company, errors.Wrap(err, "error while retrieving all users from database")
	}

	for _, company := range temp {
		if company.Name == name && company.Country == country {
			return company, nil
		}
	}

	return company, errors.New("company not found, quitting")
}

// RetrieveAllCompanies gets a list of all companies in the database
func RetrieveAllCompanies() ([]Company, error) {
	var companies []Company
	keys, err := edb.RetrieveAllKeys(globals.DbPath, CompanyBucket)
	if err != nil {
		log.Println(err)
		return companies, errors.Wrap(err, "could not retrieve all user keys")
	}
	for _, val := range keys {
		var x Company
		err = json.Unmarshal(val, &x)
		if err != nil {
			break
		}
		companies = append(companies, x)
	}

	return companies, nil
}
