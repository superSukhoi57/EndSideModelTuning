package main

import (
	"context"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

func main() {
	client := openai.NewClient(
		//option.WithAPIKey(os.Getenv("DASHSCOPE_API_KEY")),
		option.WithAPIKey("sk-5a52357f6b234c5987cb1bef8cba3756"),
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"),
	)
	chatCompletion, err := client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage("你是谁"),
			},
			Model: "qwen-plus",
		},
	)

	if err != nil {
		panic(err.Error())
	}

	println(chatCompletion.Choices[0].Message.Content)
}
