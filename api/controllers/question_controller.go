package controllers

import (
	"context"
	b64 "encoding/base64"
	"os"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func AllQuestions(c *fiber.Ctx) error {
	ctx := context.Background()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("AccessToken")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	_, dir, _, err := client.Repositories.GetContents(ctx, "always-maap", "Python-Leet-Code", "/", nil)
	if err != nil {
		return c.SendString(err.Error())
	}

	question := []string{}
	for _, file := range dir {
		if strings.HasSuffix(*file.Name, ".py") {
			question = append(question, *file.Name)
		}
	}

	contents := []string{}
	var wg sync.WaitGroup
	for _, q := range question {
		wg.Add(1)
		go func(q string, contents *[]string) error {
			file, _, _, err := client.Repositories.GetContents(ctx, "always-maap", "Python-Leet-Code", q, nil)
			if err != nil {
				return c.SendString(err.Error())
			}

			decoded, _ := b64.StdEncoding.DecodeString(*file.Content)
			decodedStr := string(decoded)

			*contents = append(*contents, decodedStr)
			wg.Done()
			return nil
		}(q, &contents)
	}

	wg.Wait()
	return c.JSON(contents)
}
