package mail

import (
	"bytes"
	"github.com/resend/resend-go/v2"
	"html/template"
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

func readTemplate(emailTemplate string) (string, string) {
	var bf bytes.Buffer
	htmlPath, err := filepath.Abs(emailMap[emailTemplate] + "index.html")
	tmpl := template.Must(template.ParseFiles(htmlPath))
	textPath, err := filepath.Abs(emailMap[emailTemplate] + "text.txt")

	if err := tmpl.Execute(&bf, nil); err != nil {
		panic(err)
	}

	htmlStr := bf.String()
	if err != nil {
		log.Println("template path mess up")
	}
	text, err := os.ReadFile(textPath)
	if err != nil {
		log.Println("failed to read plain text")
	}
	return htmlStr, string(text)
}

func SendMail(input Params, data map[string]string) {
	client := resend.NewClient(os.Getenv("RESEND_API"))
	htmlStr, textStr := readTemplate(input.Template)
	for k, v := range data {
		log.Println(v)
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
