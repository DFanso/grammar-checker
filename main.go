// main.go
package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

const (
	systemPrompt = `You are an English grammar expert and English teacher specializing in detailed grammar analysis. Your role is to:
1. Identify and correct grammatical errors
2. Provide comprehensive explanations for each correction
3. Focus on proper English usage, punctuation, and capitalization
4. Always include relevant grammar rules and their applications
5. Maintain a professional yet approachable teaching tone

When reviewing text, you must ALWAYS include these sections:
1. Original text (marked with **Original:**)
2. Corrected version (marked with **Corrected:**)
3. Detailed explanations for each correction (marked with **Explanation:**)
4. Applicable grammar rules with examples (marked with **Rules:**)

Format your response exactly like this example:

**Original:** "did u get the aws account"

**Corrected:** "Did you get the AWS account?"

**Explanation:**
- "u" should be "you": Informal abbreviations should be written in full
- "aws" should be "AWS": As a company name abbreviation, it should be capitalized
- Added question mark: Interrogative sentences require question marks
- Capitalized first word: Sentences should start with capital letters

**Rules:**
- Capitalization Rule: Sentences must begin with a capital letter
- Abbreviation Rule: Common abbreviations like "u" should be written in full as "you" in formal writing
- Proper Noun Rule: Names, brands, and company abbreviations (like AWS) should be capitalized
- Question Mark Rule: All direct questions must end with a question mark

Always provide detailed rules relevant to the corrections made. Each response must include ALL sections: Original, Corrected, Explanation, and Rules.`
)

// Define color styles
var (
	titleStyle       = color.New(color.FgBlue, color.Bold)
	promptStyle      = color.New(color.FgCyan)
	errorStyle       = color.New(color.FgRed)
	originalStyle    = color.New(color.FgYellow)
	correctedStyle   = color.New(color.FgGreen)
	explanationStyle = color.New(color.FgWhite)
	ruleStyle        = color.New(color.FgMagenta)
)

func formatResponse(text string) {
	lines := strings.Split(text, "\n")
	currentSection := ""

	for _, line := range lines {
		line = strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(line, "**Original:**"):
			currentSection = "original"
			titleStyle.Print("\nOriginal: ")
			originalStyle.Println(strings.TrimPrefix(line, "**Original:**"))

		case strings.HasPrefix(line, "**Corrected:**"):
			currentSection = "corrected"
			titleStyle.Print("\nCorrected: ")
			correctedStyle.Println(strings.TrimPrefix(line, "**Corrected:**"))

		case strings.HasPrefix(line, "**Explanation:**"):
			currentSection = "explanation"
			titleStyle.Println("\nExplanations:")

		case strings.HasPrefix(line, "**Rules:**"):
			currentSection = "rules"
			titleStyle.Println("\nRules:")

		case strings.HasPrefix(line, "-"):
			if currentSection == "rules" {
				ruleStyle.Printf("  %s\n", line)
			} else {
				explanationStyle.Printf("  %s\n", line)
			}

		case line != "":
			if currentSection == "rules" {
				ruleStyle.Println("  " + line)
			} else {
				explanationStyle.Println(line)
			}
		}
	}
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: Error loading .env file")
	}

	// Initialize Gemini client
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		errorStyle.Fprintf(os.Stderr, "Failed to create client: %v\n", err)
		os.Exit(1)
	}
	defer client.Close()

	// Initialize model
	model := client.GenerativeModel("gemini-pro")

	// Set up chat session with system prompt
	chat := model.StartChat()
	_, err = chat.SendMessage(ctx, genai.Text(systemPrompt))
	if err != nil {
		errorStyle.Fprintf(os.Stderr, "Failed to send system prompt: %v\n", err)
		os.Exit(1)
	}

	titleStyle.Println("\n=== Grammar Checker ===")
	promptStyle.Println("Type your text (press Ctrl+C to exit):")

	// Start input loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		promptStyle.Print("\n> ")
		if !scanner.Scan() {
			break
		}

		userInput := strings.TrimSpace(scanner.Text())
		if userInput == "" {
			continue
		}

		// Send user input to Gemini
		response, err := chat.SendMessage(ctx, genai.Text(userInput))
		if err != nil {
			errorStyle.Fprintf(os.Stderr, "Error getting response: %v\n", err)
			continue
		}

		// Print response
		fmt.Println() // Add spacing
		for _, candidate := range response.Candidates {
			for _, part := range candidate.Content.Parts {
				formatResponse(fmt.Sprintf("%v", part))
			}
		}
		fmt.Println() // Add spacing
	}

	if err := scanner.Err(); err != nil {
		errorStyle.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
