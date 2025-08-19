package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "srt-fixer",
	Short: "CLI that fixes broken srt subtitles ",
	Long: `CLI that fixes broken srt subtitles
	You need a gemini API key, you can get for free at: https://aistudio.google.com/app/apikey 

	Before using the tool, please set the google API KEY as an env variable:

	For Linux/macOS:
		export GEMINI_API_KEY="YOUR_API_KEY

	For Windows:
		$env:GEMINI_API_KEY="YOUR_API_KEY"

	Usage:
		srt-fixer fix path/to/your/subtitles.srt
	`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.AddCommand(FixSubs)
}
