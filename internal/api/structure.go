package api


type Message struct {
    Role    string `json:"role"`
    Content string `json:"content"`
}

type Reasoning struct {
    Enabled bool `json:"enabled"`
}
type RequestBody struct {
    Model     string    `json:"model"`
    Messages  []Message `json:"messages"`
    Reasoning Reasoning `json:"reasoning"`
}

type AnswerMessage struct {
	Content string `json:"content"`
}

type AnswerChoice struct {
	Message AnswerMessage `json:"message"`
}
type Answer struct {
	Choices []AnswerChoice `json:"choices"`
}
