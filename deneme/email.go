package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func main() {
	// JSON dosyasını okuyun
	jsonFile, err := os.Open("credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()

	// JSON dosyasındaki verileri okuyun
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var credentials struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		RefreshToken string `json:"refresh_token"`
	}
	json.Unmarshal(byteValue, &credentials)

	// OAuth2 yapılandırması
	config := &oauth2.Config{
		ClientID:     credentials.ClientID,
		ClientSecret: credentials.ClientSecret,
		Endpoint:     google.Endpoint,
		Scopes: []string{
			gmail.MailGoogleComScope,
		},
	}

	// Access token alınması
	token := &oauth2.Token{
		RefreshToken: credentials.RefreshToken,
		TokenType:    "Bearer",
	}
	client := config.Client(context.Background(), token)

	// Gmail istemcisini oluşturun
	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Gmail client: %v", err)
	}

	// E-posta mesajı oluşturun
	to := "muhammetyusufkaydi@gmail.com"
	subject := "Golang ile E-posta gönderme"
	body := "Merhaba,\nBu bir test mesajıdır."
	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body + "\r\n")

	// E-posta gönderme işlemi
	_, err = srv.Users.Messages.Send("me", &gmail.Message{
		Raw: string(message),
	}).Do()
	if err != nil {
		log.Fatalf("Unable to send email: %v", err)
	}
	fmt.Println("E-posta başarıyla gönderildi!")
}
