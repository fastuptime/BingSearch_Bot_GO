package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type SearchResult struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

type SearchResponse struct {
	Results     []SearchResult `json:"results"`
	CurrentPage int            `json:"current_page"`
	NextPage    int            `json:"next_page"`
}

func searchBing(query string, page int, numResults int, language string, country string) (*SearchResponse, error) {
	var results []SearchResult

	searchURL := fmt.Sprintf("https://www.bing.com/search?q=%s&first=%d&setlang=%s&cc=%s", url.QueryEscape(query), (page-1)*numResults+1, language, country)

	client := &http.Client{}
	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")
	req.Header.Set("Accept-Language", language)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("li.b_algo").Each(func(i int, s *goquery.Selection) {
		if i >= numResults {
			return
		}

		title := s.Find("h2").Text()
		description := s.Find("p").Text()
		link := s.Find("a").AttrOr("href", "")

		if title != "" && link != "" {
			results = append(results, SearchResult{
				Title:       strings.TrimSpace(title),
				Description: strings.TrimSpace(description),
				URL:         link,
			})
		}
	})

	nextPage := page + 1
	if len(results) < numResults {
		nextPage = 0
	}

	return &SearchResponse{
		Results:     results,
		CurrentPage: page,
		NextPage:    nextPage,
	}, nil
}

func main() {
	var keyword string
	fmt.Print("Aramak istediğiniz anahtar kelimeyi girin: ")
	fmt.Scanln(&keyword)

	var language string
	fmt.Print("Arama dilini girin (örneğin, 'en' için İngilizce, 'tr' için Türkçe): ")
	fmt.Scanln(&language)

	var country string
	fmt.Print("Bölge kodunu girin (örneğin, 'US' için Amerika, 'TR' için Türkiye): ")
	fmt.Scanln(&country)

	page := 1

	for {
		fmt.Printf("Bing'de arama yapılıyor (Sayfa %d)...\n", page)
		response, err := searchBing(keyword, page, 10, language, country)
		if err != nil {
			fmt.Println("Arama sırasında hata:", err)
			return
		}

		jsonData, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			fmt.Println("JSON dönüşümü sırasında hata:", err)
			return
		}
		fmt.Println(string(jsonData))

		if response.NextPage == 0 {
			fmt.Println("Son sayfaya ulaşıldı.")
			break
		}

		fmt.Print("Sonraki sayfayı görmek istiyor musunuz? (e/h): ")
		var input string
		fmt.Scanln(&input)
		if strings.ToLower(input) != "e" {
			break
		}

		page = response.NextPage

		time.Sleep(5 * time.Second)
	}
}
