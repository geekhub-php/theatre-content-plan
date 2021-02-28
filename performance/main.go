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
	performancesUrl = baseUrl + "/api/seasons/0/performances?locale=uk"
	performanceUrl  = "https://theatre-shevchenko.ck.ua/uk/performance/%s"
	rolesUrl        = baseUrl + "/api/performances/%s/roles?locale=uk"
)

type Performance struct {
	Title string `json:"title"`
	Desc core.Description `json:"description"`
	Slug string `json:"slug"`
	Gallery []struct{} `json:"gallery"`
	PerformanceUrl string
	Roles []Role
	GalleryImgNumb int
	RoleCount int
}
type Role struct {
	Title string `json:"title"`
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

	for i := range performances {
		performances[i].updateFields()
	}

	core.WriteCsv(performanceCsvFields)
	for _, p := range performances {
		core.WriteCsv(p.ToSlice())
	}

	core.FlushCsv()
}



var performanceCsvFields = []string{
	"Title", "Link", "Description characters without spaces", "Photos", "Roles",
}
func (p Performance) ToSlice()  []string {
	return []string{
		p.Title,
		p.PerformanceUrl,
		strconv.Itoa(p.Desc.RuneCount()),
		strconv.Itoa(len(p.Gallery)),
		strconv.Itoa(len(p.Roles)),
	}
}

func (p *Performance) updateFields() {
	p.updateRoles()
	p.PerformanceUrl = fmt.Sprintf(performanceUrl, p.Slug)
}

func (p *Performance) updateRoles() {
	url := fmt.Sprintf(rolesUrl, p.Slug)
	body, err := core.GetResource(url)
	if err != nil {
		log.Fatal(err)
	}

	roles := []Role{}
	if err := json.Unmarshal(body, &roles); err != nil {
		log.Fatal(err)
	}

	p.Roles = roles
	p.RoleCount = len(roles)
}
