package handlers

import (
	"thedekk/AIReview/internal/api"
	"os/exec"
	"fmt"
)


func Request(CurrentBranch, MainBranch string, CustomPromt bool, OutFile string, SupplementationPromt string) error {
	cmd := exec.Command("git", "diff", fmt.Sprintf("%s..%s", MainBranch, CurrentBranch))
	output, _ := cmd.Output()


	fmt.Println("Git diff output:", string(output))

	if err := api.Request(string(output), SupplementationPromt, CustomPromt); err != nil {
		return err
	}
	return nil
}
