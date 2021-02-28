package main

import (
	"encoding/json"
	"fmt"
	"github.com/geekhub-php/theatre-content-plan/core"
	"log"
	"strconv"
)

const (
	baseUrl      = "https://theatre-shevchenko.ck.ua/uk"
	employeesUrl = "http://apistaging.theatre.pp.ua/api/employees?locale=uk"
)
type Employee struct {
	FirstName string     	   `json:"first_name"`
	LastName  string    	   `json:"last_name"`
	Position  string   		   `json:"position"`
	Slug      string   		   `json:"slug"`
	Bio       core.Description `json:"biography"`
	Link      string
}
var employeesFields = []string{
	"Employee Name", "Link", "Description characters without spaces",
}
func (em Employee) ToSlice() []string {
	return []string{em.FirstName+" "+em.LastName, em.Link, strconv.Itoa(em.Bio.RuneCount())}
}

func main() {
	body, err := core.GetResource(employeesUrl)

	if err != nil {
		log.Fatal(err)
	}

	employees := []Employee{}
	if err := json.Unmarshal(body, &employees); err != nil {
		log.Fatalf("Cannot deserialize json: %v\n", err)
	}

	generateEmployeeLink(employees)

	core.WriteCsv(employeesFields)
	for _, em := range employees {
		core.WriteCsv(em.ToSlice())
	}
	core.FlushCsv()
}

func generateEmployeeLink(employees []Employee) {
	for i, employee := range employees {
		employees[i].Link = fmt.Sprintf("%v/persons/%v", baseUrl, employee.Slug)
	}
}
