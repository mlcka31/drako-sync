package chatgpt

import (
	"backend/contract"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/openai/openai-go"
	"github.com/openai/openai-go/option"
	"github.com/openai/openai-go/shared"
	"go.uber.org/zap"
	"os"
)

// Config holds the message input for ChatGPT
type Config struct {
	Message string `json:"message"`
}

// ChatGPTClient wraps the OpenAI client
type ChatGPTClient struct {
	client      openai.Client
	initPrompt  string
	logger      *zap.Logger
	aiResponses []string
	aiAddress   string
}

// NewChatGPTClient creates a new ChatGPT client with the given API key
func NewChatGPTClient(apiKey string, initPrompt string, logger *zap.Logger) *ChatGPTClient {
	return &ChatGPTClient{
		initPrompt: initPrompt,
		logger:     logger,
		client:     openai.NewClient(option.WithAPIKey(apiKey), option.WithBaseURL("https://api.openai.com/v1")),
	}
}

func (c *ChatGPTClient) SetAIAddress(aiAddress string) {
	c.aiAddress = aiAddress
}

func (c *ChatGPTClient) SendAllMessages(ctx context.Context, allMessages []contract.AlephGameStateMessage) (string, error) {
	//prompt := fmt.Sprintf(c.initPrompt+" '%s'", marshalArrayToJSON(allMessages))

	c.logger.Debug("Sending messages to ChatGPT")

	userMessages, agentMessages := c.splitMessagesIntoUserAndAgentOnes(allMessages)
	promptArray := c.buildChatCompletionMessage(userMessages, agentMessages)

	c.logger.Debug("promptArray", zap.Any("promptArray", marshalChatCompletionMessageParamUnionArrayToJSON(promptArray)))

	resp, err := c.client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: promptArray,
		Model:    shared.ChatModelGPT4,
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

func (c *ChatGPTClient) buildChatCompletionMessage(userMessages []string, agentMessages []string) []openai.ChatCompletionMessageParamUnion {
	messages := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(c.initPrompt),
	}
	// join two arrays one by one
	// at first goes one user message, then one agent messages
	// then again one user message, then one agent messages
	// if there are no more messages from agent add to the end of the array all left messages from user
	// if there are no more messages from user add to the end of the array all left messages from agent
	for i := 0; i < len(userMessages) || i < len(agentMessages); i++ {
		if i < len(userMessages) {
			messages = append(messages, openai.UserMessage(userMessages[i]))
		}
		if i < len(agentMessages) {
			messages = append(messages, openai.AssistantMessage(agentMessages[i]))
		}
	}

	return messages
}

func (c *ChatGPTClient) splitMessagesIntoUserAndAgentOnes(allMessages []contract.AlephGameStateMessage) ([]string, []string) {
	// code which separates allMessages into two arrays
	// one for user messages and one for agent messages
	userMessages := make([]string, 0)
	agentMessages := make([]string, 0)
	for _, message := range allMessages {
		if message.Sender == common.HexToAddress(c.aiAddress) {
			userMessages = append(userMessages, message.Content)
		} else {
			agentMessages = append(agentMessages, message.Content)
		}
	}
	return userMessages, agentMessages
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

func marshalArrayToJSON(array []any) string {
	jsonData, _ := json.Marshal(array)
	return string(jsonData)
}

// marshal `[]openai.ChatCompletionMessageParamUnion` to JSON
func marshalChatCompletionMessageParamUnionArrayToJSON(array []openai.ChatCompletionMessageParamUnion) string {
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
