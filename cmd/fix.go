package cmd

import (
	"path/filepath"

	"github.com/spf13/cobra"
)

var FixSubs = &cobra.Command{
	Use:   "fix [filePath]",
	Short: "fix the file at the filePath",
	Long:  "fix the srt at the given filePath",
	Args:  cobra.ExactArgs(1),
	Run:   fixSrt,
}

func fixSrt(cmd *cobra.Command, args []string) {
	originalFilePath := args[0]

	srtString := ReadString(originalFilePath)

	completePrompt := CreatePrompt(srtString, originalFilePath)

	correctedContent := SendRequest(completePrompt)

	fileName := "fixed_" + filepath.Base(originalFilePath)
	CreateFile(correctedContent, fileName)

}
