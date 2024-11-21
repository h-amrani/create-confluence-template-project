package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Confluence API configuration
const (
	baseURL     = "https://<your-confluence-domain>/wiki/rest/api/content"
	username    = "<your-username>"
	apiToken    = "<your-api-token>"
	spaceKey    = "<space-key>"
	contentType = "application/json"
)

// Page structure for Confluence
type Page struct {
	Type       string      `json:"type"`
	Title      string      `json:"title"`
	Space      Space       `json:"space"`
	Ancestors  []Ancestor  `json:"ancestors,omitempty"`
	Body       PageBody    `json:"body"`
}

type Space struct {
	Key string `json:"key"`
}

type Ancestor struct {
	ID string `json:"id"`
}

type PageBody struct {
	Storage Content `json:"storage"`
}

type Content struct {
	Value          string `json:"value"`
	Representation string `json:"representation"`
}

// Function to create a Confluence page
func createPage(title, parentID, body string) (string, error) {
	page := Page{
		Type:  "page",
		Title: "go_" + title, // Add the "go_" prefix
		Space: Space{Key: spaceKey},
		Body: PageBody{
			Storage: Content{
				Value:          body,
				Representation: "storage",
			},
		},
	}

	if parentID != "" {
		page.Ancestors = []Ancestor{{ID: parentID}}
	}

	// Convert page structure to JSON
	jsonData, err := json.Marshal(page)
	if err != nil {
		return "", fmt.Errorf("error marshalling JSON: %v", err)
	}

	// Make HTTP POST request
	client := &http.Client{}
	req, err := http.NewRequest("POST", baseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}
	req.SetBasicAuth(username, apiToken)
	req.Header.Set("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("error response from server: %s", string(bodyBytes))
	}

	// Parse response JSON
	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	pageID := result["id"].(string)
	fmt.Printf("Page '%s' created successfully with ID: %s\n", "go_"+title, pageID)
	return pageID, nil
}

func main() {
	// Create Home Page
	homePageTitle := "Project Name Overview"
	homePageBody := "<h1>Welcome to the Project</h1><p>This space is dedicated to organizing the entire project lifecycle.</p>"

	homePageID, err := createPage(homePageTitle, "", homePageBody)
	if err != nil {
		log.Fatalf("Failed to create home page: %v", err)
	}

	// Define main sections with their subcategories
	sections := []struct {
		Title        string
		Body         string
		Subcategories []struct {
			Title string
			Body  string
		}
	}{
		{
			Title: "Project Overview",
			Body:  "<h1>Project Overview</h1><p>Objectives, scope, and assumptions for the project.</p>",
			Subcategories: []struct {
				Title string
				Body  string
			}{
				{"Project Objectives", "<p>High-level description of the project's goals.</p>"},
				{"Scope", "<p>In-scope and out-of-scope items.</p>"},
				{"Assumptions", "<p>Key dependencies and constraints.</p>"},
			},
		},
		{
			Title: "Requirements",
			Body:  "<h1>Requirements</h1><p>Functional and non-functional requirements, constraints, and workflows.</p>",
			Subcategories: []struct {
				Title string
				Body  string
			}{
				{"Functional Requirements", "<p>Core functionalities and user stories.</p>"},
				{"Non-Functional Requirements", "<p>Performance, scalability, and usability requirements.</p>"},
				{"Constraints", "<p>Technical and timeline limitations.</p>"},
				{"AS-IS and TO-BE Processes", "<p>Current and future workflows.</p>"},
			},
		},
		{
			Title: "Architecture and Technical Design",
			Body:  "<h1>Architecture and Technical Design</h1><p>System architecture and technical solutions.</p>",
			Subcategories: []struct {
				Title string
				Body  string
			}{
				{"System Architecture", "<p>High-level diagrams and descriptions.</p>"},
				{"Data Flow Diagrams", "<p>Data movement workflows.</p>"},
				{"ETL Process Design", "<p>ETL pipelines for data integration.</p>"},
				{"Deployment Architecture", "<p>Cloud infrastructure details.</p>"},
			},
		},
		{
			Title: "Development Documentation",
			Body:  "<h1>Development Documentation</h1><p>Technical guides for developers.</p>",
			Subcategories: []struct {
				Title string
				Body  string
			}{
				{"API Documentation", "<p>Endpoints and schemas.</p>"},
				{"Database Schema", "<p>Table design and indexing.</p>"},
				{"CI/CD Pipeline", "<p>Build, test, and deployment processes.</p>"},
			},
		},
		{
			Title: "Testing and QA",
			Body:  "<h1>Testing and QA</h1><p>Testing plans and criteria.</p>",
			Subcategories: []struct {
				Title string
				Body  string
			}{
				{"Testing Strategy", "<p>Unit, integration, and end-to-end tests.</p>"},
				{"Test Cases", "<p>Detailed test cases.</p>"},
				{"UAT Plans", "<p>User acceptance testing.</p>"},
			},
		},
		{
			Title: "Reporting and Metrics",
			Body:  "<h1>Reporting and Metrics</h1><p>Monitoring and reporting requirements.</p>",
			Subcategories: []struct {
				Title string
				Body  string
			}{},
		},
		{
			Title: "Change Management",
			Body:  "<h1>Change Management</h1><p>Track changes and approvals.</p>",
			Subcategories: []struct {
				Title string
				Body  string
			}{},
		},
		{
			Title: "Training and Support",
			Body:  "<h1>Training and Support</h1><p>User guides and support contacts.</p>",
			Subcategories: []struct {
				Title string
				Body  string
			}{},
		},
	}

	// Create main sections and their subcategories
	for _, section := range sections {
		sectionID, err := createPage(section.Title, homePageID, section.Body)
		if err != nil {
			log.Printf("Failed to create section '%s': %v", section.Title, err)
			continue
		}

		// Create subcategories
		for _, sub := range section.Subcategories {
			_, err := createPage(sub.Title, sectionID, sub.Body)
			if err != nil {
				log.Printf("Failed to create subcategory '%s': %v", sub.Title, err)
			}
		}
	}
}