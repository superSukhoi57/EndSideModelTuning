package llm

import (
	"context"

	"log"

	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
)

type BLClient struct {
	Client *openai.Client
	Model  string
}

// sk-5a52357f6b234c5987cb1bef8cba3756,qwen-plus
func CreateLLMClient(APIKey, model string) *BLClient {
	client := openai.NewClient(
		//option.WithAPIKey(os.Getenv("DASHSCOPE_API_KEY")),
		option.WithAPIKey(APIKey),
		option.WithBaseURL("https://dashscope.aliyuncs.com/compatible-mode/v1"),
	)
	return &BLClient{
		Client: &client,
		Model:  model,
	}
}

func (b *BLClient) Chat(msg string) (answer string) {

	chatCompletion, err := b.Client.Chat.Completions.New(
		context.TODO(), openai.ChatCompletionNewParams{
			Messages: []openai.ChatCompletionMessageParamUnion{
				openai.UserMessage(msg),
			},
			Model: b.Model,
		},
	)

	if err != nil {
		log.Println(err.Error())
	}

	answer = chatCompletion.Choices[0].Message.Content
	println(answer)
	return

}
