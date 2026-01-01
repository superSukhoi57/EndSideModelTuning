package main

import "backend/common/llm"

func main() {
	client := llm.CreateLLMClient("sk-5a52357f6b234c5987cb1bef8cba3756", "qwen-plus")
	client.Chat("你是谁")
}
