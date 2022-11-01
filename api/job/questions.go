package main

import (
	"context"
	b64 "encoding/base64"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"

	"api/database"
	"api/models"

	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	ctx := context.Background()

	database.Connect()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("AccessToken")},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	questions, err := getAllQuestions(client, ctx)

	if err != nil {
		panic(err)
	}

	status, err := storeQuestionContent(client, ctx, questions)

	if err != nil {
		panic(err)
	}

	fmt.Println(status)
}

func getAllQuestions(client *github.Client, ctx context.Context) ([]string, error) {
	_, dir, _, err := client.Repositories.GetContents(ctx, "always-maap", "Python-Leet-Code", "/", nil)
	if err != nil {
		panic(err)
	}

	question := []string{}
	for _, file := range dir {
		if strings.HasSuffix(*file.Name, ".py") {
			question = append(question, *file.Name)
		}
	}

	return question, nil
}

func storeQuestionContent(client *github.Client, ctx context.Context, questions []string) (string, error) {

	var wg sync.WaitGroup
	for _, q := range questions {
		wg.Add(1)
		go func(q string) error {
			file, _, _, err := client.Repositories.GetContents(ctx, "always-maap", "Python-Leet-Code", q, nil)
			if err != nil {
				return err
			}

			decoded, _ := b64.StdEncoding.DecodeString(*file.Content)
			decodedStr := string(decoded)

			question := parseQuestion(decodedStr)
			database.DB.Create(&question)

			wg.Done()
			return nil
		}(q)
	}

	wg.Wait()

	return "success", nil
}

func parseQuestion(contents string) models.Question {
	question := models.Question{}

	r, _ := regexp.Compile(`"""(.|[\r\n])*"""`)

	idx := r.FindStringIndex(contents)
	start, end := idx[0], idx[1]

	desc := contents[start:end]
	desc = strings.ReplaceAll(desc, "\"\"\"", "")
	desc = strings.ReplaceAll(desc, "\r", "")

	desc = strings.TrimSpace(desc)

	temp := strings.Split(desc, "\n")
	question.No, _ = strconv.Atoi(strings.Split(temp[0], ".")[0])
	question.Name = strings.Split(temp[0], ". ")[1]
	question.Difficulty = temp[1]
	question.Subject = temp[2]
	question.Problem = strings.Join(temp[3:], "\n")

	problem := strings.Split(desc, "---")[1]
	problem = strings.TrimSpace(problem)
	question.Problem = problem

	solution := contents[end:]
	solution = strings.TrimSpace(solution)
	question.Solution = solution

	return question
}
