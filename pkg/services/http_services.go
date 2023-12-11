package services

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/IshanSaha05/Web_Scrapper_Election/pkg/config"
	"github.com/IshanSaha05/Web_Scrapper_Election/pkg/models"
	"github.com/PuerkitoBio/goquery"
)

func GetSite(url string) (*http.Response, error) {
	return http.Get(url)
}

func GetLink_ByName(constituency_name string) (string, error) {
	url := config.DefaultSite
	response, err := GetSite(url)

	if err != nil {
		fmt.Println("Error while fetching the site: \"", url, "\"")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("Error while parsing the body for GoQuery.")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	info := doc.Find("#ctl00_ContentPlaceHolder1_Result1_ddlState").Find("option")

	var id string
	found := false

	info.Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		if i > 0 {
			result := strings.Split(text, " - ")

			if result[0] == constituency_name {
				id = result[1]
				found = true
				return
			}
		}
	})

	if found {
		fmt.Println("Link: ", fmt.Sprintf("https://results.eci.gov.in/AcResultGenDecNew2023/candidateswise-S29%s.htm", id))
		return fmt.Sprintf("https://results.eci.gov.in/AcResultGenDecNew2023/candidateswise-S29%s.htm", id), nil
	}

	fmt.Println("Error: Could not find any constituency id with constituency name: ", constituency_name)

	return "", fmt.Errorf(fmt.Sprintf("Invalid constituency name \"%s\"", constituency_name))

}

func GetLink_ByID(constituency_id string) (string, error) {
	url := config.DefaultSite
	response, err := GetSite(url)

	if err != nil {
		fmt.Println("Error while fetching the site: \"", url, "\"")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("Error while parsing the body for GoQuery.")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	info := doc.Find("#ctl00_ContentPlaceHolder1_Result1_ddlState").Find("option")

	var id string
	found := false

	info.Each(func(i int, s *goquery.Selection) {
		text := s.Text()

		if i > 0 {
			result := strings.Split(text, " - ")

			if result[1] == constituency_id {
				id = result[1]
				found = true
				return
			}
		}
	})

	if found {
		fmt.Println("Link: ", fmt.Sprintf("https://results.eci.gov.in/AcResultGenDecNew2023/candidateswise-S29%s.htm", id))
		return fmt.Sprintf("https://results.eci.gov.in/AcResultGenDecNew2023/candidateswise-S29%s.htm", id), nil
	}

	fmt.Println("Error: Could not find any constituency id with constituency name: ", constituency_id)

	return "", fmt.Errorf(fmt.Sprintf("Invalid constituency name \"%s\"", constituency_id))
}

func parseResults(res *http.Response) []models.Results {
	doc, err := goquery.NewDocumentFromReader(res.Body)

	if err != nil {
		fmt.Println("Error while parsing the body for GoQuery.")
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	candBox := doc.Find(".cand-box")

	results := []models.Results{}

	candBox.Each(func(i int, p *goquery.Selection) {
		candInfo := p.Find(".cand-info")
		statusInfo := candInfo.Find(".status")

		result := models.Results{}

		result.Status = statusInfo.Find("div").Eq(0).Text()
		result.VoteShare = statusInfo.Find("div").Eq(1).Text()

		partyInfo := candInfo.Find(".nme-prty")

		result.CandidateName = partyInfo.Find("h5").Text()
		result.CandidateParty = partyInfo.Find("h6").Text()

		results = append(results, result)
	})

	return results
}

func printResults(results []models.Results) {
	// Printing the results.
	for _, result := range results {
		fmt.Println("-------------------------------------------------")
		fmt.Println("Candidate Name: ", result.CandidateName)
		fmt.Println("Candidate Party Name: ", result.CandidateParty)
		fmt.Println("Status: ", result.Status)
		fmt.Println("Vote Share: ", result.VoteShare)
		fmt.Println("-------------------------------------------------")
	}
}

func ParsePrintResults(response *http.Response) {
	// Parsing the results.
	results := parseResults(response)

	// Printing the results.
	printResults(results)
}
