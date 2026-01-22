package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Context struct {
	ProductName string   `json:"product_name"`
	LaunchDate  string   `json:"launch_date"`
	KeyFeatures []string `json:"key_features"`
	PainPoints  []string `json:"pain_points"`
}

type Post struct {
	Content       string `json:"content"`
	SuggestedDate string `json:"suggested_date"`
	CharCount     int    `json:"char_count"`
	Template      string `json:"template"`
	Type          string `json:"type"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: porter-social [generate|publish]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "generate":
		generatePosts()
	case "publish":
		publishPosts()
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func generatePosts() {
	context, err := loadContext("context.json")
	if err != nil {
		fmt.Printf("Error loading context: %v\n", err)
		os.Exit(1)
	}

	templates := []string{"problem.md", "approach.md", "demo.md", "launch.md"}
	posts := []Post{}

	launchDate, _ := time.Parse("2006-01-02", context.LaunchDate)
	now := time.Now()
	daysUntilLaunch := int(launchDate.Sub(now).Hours() / 24)

	for i, tmpl := range templates {
		content, err := loadTemplate(tmpl)
		if err != nil {
			fmt.Printf("Error loading template %s: %v\n", tmpl, err)
			continue
		}

		content = fillTemplate(content, context)

		var suggestedDate time.Time
		switch i {
		case 0:
			suggestedDate = now.AddDate(0, 0, daysUntilLaunch/6)
		case 1:
			suggestedDate = now.AddDate(0, 0, daysUntilLaunch/2)
		case 2:
			suggestedDate = now.AddDate(0, 0, int(float64(daysUntilLaunch)*0.75))
		case 3:
			suggestedDate = launchDate
		}

		post := Post{
			Content:       content,
			SuggestedDate: suggestedDate.Format("2006-01-02"),
			CharCount:     len(content),
			Template:      tmpl,
			Type:          strings.TrimSuffix(tmpl, ".md"),
		}

		posts = append(posts, post)
	}

	outputPath := filepath.Join("output", "posts.json")
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		fmt.Printf("Error creating output directory: %v\n", err)
		os.Exit(1)
	}

	data, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		fmt.Printf("Error marshaling posts: %v\n", err)
		os.Exit(1)
	}

	if err := os.WriteFile(outputPath, data, 0644); err != nil {
		fmt.Printf("Error writing output: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated %d posts\n", len(posts))
	fmt.Printf("Output written to %s\n", outputPath)
}

func loadContext(path string) (Context, error) {
	var ctx Context
	data, err := os.ReadFile(path)
	if err != nil {
		return ctx, err
	}
	err = json.Unmarshal(data, &ctx)
	return ctx, err
}

func loadTemplate(name string) (string, error) {
	path := filepath.Join("templates", name)
	data, err := os.ReadFile(path)
	return string(data), err
}

func fillTemplate(template string, ctx Context) string {
	replacements := map[string]string{
		"{{.ProductName}}":    ctx.ProductName,
		"{{.LaunchDate}}":     ctx.LaunchDate,
		"{{.PainPoint}}":      ctx.PainPoints[0],
		"{{.KeyFeature}}":     ctx.KeyFeatures[0],
		"{{.ExampleMention}}": "@porter investigate-ci-failure",
	}

	result := template
	for placeholder, value := range replacements {
		result = strings.ReplaceAll(result, placeholder, value)
	}

	return result
}

func publishPosts() {
	fmt.Println("Publish functionality not yet implemented")
	fmt.Println("Manually copy posts from output/posts.json")
}
