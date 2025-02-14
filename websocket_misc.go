package slack

import (
	"encoding/json"
	"fmt"
)

// AckMessage is used for messages received in reply to other messages
type AckMessage struct {
	ReplyTo   int    `json:"reply_to"`
	Timestamp string `json:"ts"`
	Text      string `json:"text"`
	RTMResponse
}

type DesktopNotification struct {
	AvatarImage string      `json:"avatarImage"`
	Channel     string      `json:"channel"`
	Content     string      `json:"content"`
	EventTs     string      `json:"event_ts"`
	ImageURI    interface{} `json:"imageUri"`
	IsShared    bool        `json:"is_shared"`
	LaunchURI   string      `json:"launchUri"`
	Msg         string      `json:"msg"`
	SsbFilename string      `json:"ssbFilename"`
	Subtitle    string      `json:"subtitle"`
	Title       string      `json:"title"`
	Type        string      `json:"type"`
}

// RTMResponse encapsulates response details as returned by the Slack API
type RTMResponse struct {
	Ok    bool      `json:"ok"`
	Error *RTMError `json:"error"`
}

// RTMError encapsulates error information as returned by the Slack API
type RTMError struct {
	Code int
	Msg  string
}

func (s RTMError) Error() string {
	return fmt.Sprintf("Code %d - %s", s.Code, s.Msg)
}

// MessageEvent represents a Slack Message (used as the event type for an incoming message)
type MessageEvent Message

// RTMEvent is the main wrapper. You will find all the other messages attached
type RTMEvent struct {
	Type string
	Data interface{}
}

// HelloEvent represents the hello event
type HelloEvent struct{}

// PresenceChangeEvent represents the presence change event
type PresenceChangeEvent struct {
	Type     string `json:"type"`
	Presence string `json:"presence"`
	User     string `json:"user"`
}

// UserTypingEvent represents the user typing event
type UserTypingEvent struct {
	Type    string `json:"type"`
	User    string `json:"user"`
	Channel string `json:"channel"`
}

// PrefChangeEvent represents a user preferences change event
type PrefChangeEvent struct {
	Type  string          `json:"type"`
	Name  string          `json:"name"`
	Value json.RawMessage `json:"value"`
}

// ManualPresenceChangeEvent represents the manual presence change event
type ManualPresenceChangeEvent struct {
	Type     string `json:"type"`
	Presence string `json:"presence"`
}

// UserChangeEvent represents the user change event
type UserChangeEvent struct {
	Type string `json:"type"`
	User User   `json:"user"`
}

// EmojiChangedEvent represents the emoji changed event
type EmojiChangedEvent struct {
	Type           string `json:"type"`
	EventTimestamp string `json:"event_ts"`
}

// CommandsChangedEvent represents the commands changed event
type CommandsChangedEvent struct {
	Type           string `json:"type"`
	EventTimestamp string `json:"event_ts"`
}

// EmailDomainChangedEvent represents the email domain changed event
type EmailDomainChangedEvent struct {
	Type           string `json:"type"`
	EventTimestamp string `json:"event_ts"`
	EmailDomain    string `json:"email_domain"`
}

// BotAddedEvent represents the bot added event
type BotAddedEvent struct {
	Type string `json:"type"`
	Bot  Bot    `json:"bot"`
}

// BotChangedEvent represents the bot changed event
type BotChangedEvent struct {
	Type string `json:"type"`
	Bot  Bot    `json:"bot"`
}

// AccountsChangedEvent represents the accounts changed event
type AccountsChangedEvent struct {
	Type string `json:"type"`
}

// ReconnectUrlEvent represents the receiving reconnect url event
type ReconnectUrlEvent struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}
