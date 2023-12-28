package mail

import (
	"github.com/resend/resend-go/v2"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Params struct {
	To       string
	Template string
	Subject  string
}

var emailMap = map[string]string{
	"verify": "internal/pkg/mail/templates/verify/",
}

func readTemplate(template string) (string, string) {
	htmlPath, err := filepath.Abs(emailMap[template] + "index.html")
	if err != nil {
		log.Println("template path mess up")
	}
	html, err := os.ReadFile(htmlPath)
	if err != nil {
		log.Println("failed to read plain text")
	}
	textPath, err := filepath.Abs(emailMap[template] + "text.txt")
	if err != nil {
		log.Println("template path mess up")
	}
	text, err := os.ReadFile(textPath)
	if err != nil {
		log.Println("failed to read plain text")
	}
	return string(html), string(text)
}

func SendMail(input Params, data map[string]string) {
	client := resend.NewClient(os.Getenv("resend_api"))
	htmlStr, textStr := readTemplate(input.Template)
	for k, v := range data {
		log.Println(v)
		log.Println("{" + k + "}")
		htmlStr = strings.ReplaceAll(htmlStr, "{"+k+"}", v)
		textStr = strings.ReplaceAll(textStr, "{"+k+"}", v)
	}
	options := &resend.SendEmailRequest{
		From:    "Team Reach <me@iresharma.com>",
		To:      []string{input.To},
		Html:    htmlStr,
		Text:    textStr,
		Subject: input.Subject,
	}
	_, err := client.Emails.Send(options)
	if err != nil {
		panic(err)
		log.Println("error sending email")
	}

}
