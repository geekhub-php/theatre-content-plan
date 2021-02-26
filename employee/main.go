package main

import (
	"github.com/grokify/html-strip-tags-go"

	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"unicode/utf8"
)

const (
	baseUrl      = "https://theatre-shevchenko.ck.ua/uk"
	employeesUrl = "http://apistaging.theatre.pp.ua/api/employees?locale=uk"
)
type Bio string
type Employee struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Position  string `json:"position"`
	Slug      string `json:"slug"`
	Bio       Bio    `json:"biography"`
	Link      string
}
var employeesFields = []string{
	"Employee Name", "Link", "Bio characters without spaces",
}
func (em Employee) ToSlice() []string {
	return []string{em.FirstName+" "+em.LastName, em.Link, strconv.Itoa(em.Bio.CharactersWithoutSpaces())}
}

func main() {
	body, err := getResource(employeesUrl)

	if err != nil {
		log.Fatal(err)
	}

	employees := []Employee{}
	if err := json.Unmarshal(body, &employees); err != nil {
		fmt.Println(err)
		panic("Cannot deserialize json")
	}

	generateEmployeeLink(employees)

	w := csv.NewWriter(os.Stdout)
	w.Write(employeesFields)
	for _, em := range employees {
		if err := w.Write(em.ToSlice()); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}
	w.Flush()
}

func (bio Bio) CharactersWithoutSpaces() int {
	striped := strip.StripTags(string(bio))
	count := 0
	for _, str := range strings.Fields(striped) {
		count += utf8.RuneCountInString(str)
	}
	return count
}

func generateEmployeeLink(employees []Employee) {
	for i, employee := range employees {
		employees[i].Link = fmt.Sprintf("%v/persons/%v", baseUrl, employee.Slug)
	}
}

func getResource(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
