package cmd

import (
	"thedekk/AIReview/internal/api"
	"os/exec"
	"fmt"
	"os"
	"strings"
)


func Request(CurrentBranch, MainBranch string, CustomPromt bool, OutFile string, SupplementationPromt string) error {
	cmd := exec.Command("git", "diff", fmt.Sprintf("%s..%s", MainBranch, CurrentBranch))
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing git diff command:", err)
		return err
	}



	answer, err := api.Request(strings.SplitAfterN(OutFile, ".", 2)[1],string(output), SupplementationPromt, CustomPromt)
	if err != nil {
		return err
	}

	file, err := os.Create(OutFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return err
	}

	defer file.Close()

	_, err = file.WriteString(*answer)
	if err != nil {
		fmt.Println("Error writing to output file:", err)
		return err
	}

	fmt.Println("Request sent successfully")
	return nil
}
