package handlers

import (
	"thedekk/AIReview/internal/api"
	"os/exec"
	"fmt"
)


func Request(CurrentBranch, MainBranch string, CustomPromt bool, OutFile string, SupplementationPromt string) error {
	cmd := exec.Command("git", "diff", fmt.Sprintf("%s..%s", MainBranch, CurrentBranch))
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing git diff command:", err)
		return err
	}



	if err := api.Request(string(output), SupplementationPromt, CustomPromt); err != nil {
		return err
	}
	return nil
}
