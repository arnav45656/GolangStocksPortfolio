package gemini

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const API_KEY = ""

func initClient() *genai.Client {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(API_KEY))
	if err != nil {
		log.Fatalf("genai.NewClient: %v", err)
	}

	defer client.Close()

	return client

}

func GetPmt1() string {
	ctx := context.Background()
	prompt := "I need to develop a trading strategy that will help me build a portfolio with the least amount of risk." +
		"This prompt is an effective way to create a trading strategy that minimizes risk while still allowing for potential gains." +

		"To use this prompt, consider a few key variables:" +

		"Risk tolerance: Assess your risk tolerance and understand how much risk you are comfortable taking on." +
		"Investment horizon: Determine how long you plan to hold your investments." +
		"Portfolio diversification: Consider diversifying your portfolio with different asset classes, such as stocks, bonds, and commodities." +
		"Once you have a clear understanding of these variables, develop a trading strategy that incorporates them and allows you to build a portfolio with the least amount of risk. This could include strategies such as dollar-cost averaging, hedging, or investing in low-risk securities." +

		"Using this prompt will help you create a trading strategy that is tailored to your needs and minimizes the risk associated with investing."

	return GetData(prompt, ctx)

}

func GetPmt2() string {
	ctx := context.Background()
	prompt := "Get the active condition of stocks"
	return GetData(prompt, ctx)

}

func GetData(prompt string, ctx context.Context) string {
	client := initClient()
	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatalf("model.GenerateContent: %v", err)
	}

	formattedContent := formatResponse(resp)

	return formattedContent

}

// format resposne
func formatResponse(resp *genai.GenerateContentResponse) string {
	var formattedContent strings.Builder
	if resp != nil && resp.Candidates != nil {
		for _, cand := range resp.Candidates {
			if cand.Content != nil {
				for _, part := range cand.Content.Parts {
					formattedContent.WriteString(fmt.Sprintf("%v", part))
				}
			}
		}
	}

	return formattedContent.String()
}
