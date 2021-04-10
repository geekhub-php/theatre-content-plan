package main

import (
	"encoding/json"
	"fmt"
	"github.com/geekhub-php/theatre-content-plan/core"
	"log"
	"strconv"
)

const (
	baseUrl         = "https://127.0.0.1:8000"
	performancesUrl = baseUrl + "/api/seasons/-1/performances?locale=uk"
	performanceUrl  = "https://theatre-shevchenko.ck.ua/uk/performance/%s"
	rolesUrl        = baseUrl + "/api/performances/%s/roles?locale=uk"
)
type Performance struct {
	Title string `json:"title"`
	Slug string `json:"slug"`
	PerformanceUrl string
	AgeLimit int `json:"age_limit"`
}

func (p *Performance) updateFields() {
	p.PerformanceUrl = fmt.Sprintf(performanceUrl, p.Slug)
}

var performanceCsvFields = []string{
	"Title", "Link", "Age Limit",
}
func (p Performance) ToSlice()  []string {
	return []string{
		p.Title,
		p.PerformanceUrl,
		strconv.Itoa(p.AgeLimit),
	}
}

func main() {
	body, err := core.GetResource(performancesUrl)
	if err != nil {
		log.Fatal(err)
	}

	performances := []Performance{}
	if err := json.Unmarshal(body, &performances); err != nil {
		log.Fatal(err)
	}

	core.WriteCsv(performanceCsvFields)
	for _, p := range performances {
		p.updateFields()
		core.WriteCsv(p.ToSlice())
	}

	core.FlushCsv()
}
