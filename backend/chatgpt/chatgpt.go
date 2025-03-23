package chatgpt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"
	"os"
)

// Config holds the message input for ChatGPT
type Config struct {
	Message string `json:"message"`
}

// ChatGPTClient wraps the OpenAI client
type ChatGPTClient struct {
	client     openai.Client
	initPrompt string
}

// NewChatGPTClient creates a new ChatGPT client with the given API key
func NewChatGPTClient(apiKey string) *ChatGPTClient {
	initPrompt := "Has the last message in this array killed the dragon? If yes - return 'success', if no - return 'fail'."
	return &ChatGPTClient{
		initPrompt: initPrompt,
		client:     openai.NewClient(option.WithAPIKey(apiKey), option.WithBaseURL("https://api.openai.com/v1")),
	}
}

func (c *ChatGPTClient) SendAllMessages(ctx context.Context, allMessages []string) (string, error) {
	prompt := fmt.Sprintf(c.initPrompt+" '%s'", MarshalArrayToJSON(allMessages))
	resp, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(prompt),
		},
		Model: shared.ChatModelGPT4,
	})
	if err != nil {
		return "", fmt.Errorf("ChatGPT request failed: %w", err)
	}
	if len(resp.Choices) == 0 {
		return "", fmt.Errorf("no response from ChatGPT")
	}
	//return resp.Choices[0].Message.ToolCalls[0].Function.Name, nil
	return resp.Choices[0].Message.Content, nil
}

// GetMessageFromFile reads a JSON config file and returns the message
func GetMessageFromFile(configFile string) (Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	var cfg Config
	if err := json.NewDecoder(file).Decode(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to decode config: %w", err)
	}
	return cfg, nil
}

func MarshalArrayToJSON(array []string) string {
	jsonData, _ := json.Marshal(array)
	return string(jsonData)
}

func ReadPromptFromFile(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to read prompt file: %w", err)
	}
	return string(data), nil
}
