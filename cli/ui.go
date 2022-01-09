package cli

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func RenderUI() error {

	kubeConfig, err := GetKubeConfig()

	if err != nil {
		return err
	}

	currentPosition := 0
	for k, v := range kubeConfig.Contexts {
		if v == kubeConfig.CurrentContext {
			currentPosition = k
			kubeConfig.Contexts[currentPosition] = "✅ " + v
			break
		}
	}

	prompt := promptui.Select{
		Label:             "🔥 Switch a context",
		Items:             kubeConfig.Contexts,
		Size:              len(kubeConfig.Contexts),
		StartInSearchMode: false,
		CursorPos:         currentPosition,
	}

	_, result, err := prompt.Run()

	if err != nil {
		return err
	}

	if err := RunCommand(result); err != nil {
		return err
	}

	fmt.Println("You choose:", result)
	return nil
}
