package main

import (
	"bufio"
	"fmt"
	gemini_setup "gemini/example/internal/gemini-setup"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("O que voce quer saber com o Gemini ?")
	prompt, _ := reader.ReadString('\n')
	prompt = prompt[:len(prompt)-1]

	resp := gemini_setup.PromptGemini(prompt)

	fmt.Println("Response: ", resp.Candidates[0].Content)
}
