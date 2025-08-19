package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/genai"
)

func ReadString(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	srtString := string(content)
	return srtString
}

func CreatePrompt(srtString string, originalFilePath string) string {
	basePrompt := `You will be given the text of a srt subtitle file that has some corrupted characters because of encoding of special character issues, it will be provided in the following way: 
	First the name of the subtitles to give you context, then the subtitles. Please proceed and only output the corrected srt file`

	completePrompt := fmt.Sprintf(`%s
	FILENAME:
	%s

	SRT CONTENT:
	%s
	`, basePrompt, originalFilePath, srtString)

	return completePrompt
}

func SendRequest(prompt string) string {
	ctx := context.Background()
	client, err := genai.NewClient(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash-lite",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Fatal(err)
	}
	correctedContent := result.Text()
	return correctedContent
}

func CreateFile(correctedContent, fileName string) {
	err := os.WriteFile(fileName, []byte(correctedContent), 0644)
	if err != nil {
		log.Fatal("Failed to write to file", err)
	}
	fmt.Printf("Successfully saved %s", fileName)

}
