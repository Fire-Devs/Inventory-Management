package utils

import (
	"context"
	"fmt"
	"github.com/resend/resend-go/v2"
)

func SendEmail(url string, email string) error {
	ctx := context.TODO()
	client := resend.NewClient("re_JqPCDAa5_EDUUrzEYHo1tSBR3uJtuQjS9")

	params := &resend.SendEmailRequest{
		From:    "noreply<onboarding@resend.dev>",
		To:      []string{email},
		Subject: "Confirm your mail",
		Html:    fmt.Sprintf("Please click this click to verify your account <a href='%s'>Click Here</a>", url),
	}

	_, err := client.Emails.SendWithContext(ctx, params)

	if err != nil {
		return err
	}

	return nil
}
