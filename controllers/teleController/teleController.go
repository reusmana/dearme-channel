package teleController

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type WebHookReqBody struct {
	Message struct {
		Message_id int `json:message_id`
		From       struct {
			Username string `json:"username"`
			Id       int    `json:"Id"`
		} `json:"from"`
		Text string `json:"text"`
		Chat struct {
			ID int `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

type SendMessageReqBody struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}

func WebHookHandler(c *gin.Context) {
	// Create our web hook request body type instance
	body := &WebHookReqBody{}

	// Decodes the incoming request into our cutom webhookreqbody type
	if err := json.NewDecoder(c.Request.Body).Decode(body); err != nil {
		log.Printf("An error occured (webHookHandler)")
		log.Panic(err)
		return
	}

	fmt.Println("sending request")

	action := strings.ToLower(body.Message.Text)
	chatId := body.Message.Chat.ID
	msgId := body.Message.Message_id
	SendingReply(chatId, msgId, action)
}

func SendingReply(chatID int, msgId int, command string) error {
	//get message in request command
	msg := RequestCommand(command)

	reqBody := &SendMessageReqBody{
		ChatID: chatID,
		Text:   msg,
	}

	// Convert our custom type into json format
	reqBytes, err := json.Marshal(reqBody)

	if err != nil {
		return err
	}

	RequestApiTelegram(reqBytes, chatID, msgId)

	return err

}

func RequestCommand(command string) (action string) {
	if command == "/joke" {
		action = "imam ganteng"
	} else {
		action = "command not found"
	}
	return
}

func RequestApiTelegram(reqBytes []byte, chatID int, msgId int) error {
	// Make a request to send our message using the POST method to the telegram bot API
	var API = os.Getenv("API_TELEGRAM")         //Config("TELEGRAM_API")
	var TOKEN = os.Getenv("TOKEN_BOT_TELEGRAM") //Config("TELEGRAM_TOKEN")
	oriApi := API + TOKEN + "/" + "sendMessage?chat_id="
	concat := fmt.Sprintf("%s%d", oriApi, chatID)
	// reply message
	apiUrls := concat + "&reply_to_message_id="
	concatedApi := fmt.Sprintf("%s%d", apiUrls, msgId)
	// fmt.Println(concatedApi)
	resp, err := http.Post(
		concatedApi,
		"application/json",
		bytes.NewBuffer(reqBytes),
	)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + resp.Status)
	}

	return err
}
