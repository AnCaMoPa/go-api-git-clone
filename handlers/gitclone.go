package handlers

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/gofiber/fiber/v2"
)

type Repositories struct {
	Goroutines string `json:"goroutines"`
	Path       string `json:"path"`
	/* Timeout      string `json:"time out"` */
	Repositories []struct {
		URL string `json:"url"`
	} `json:"repositories"`
}

func GetName(url string) string {

	spliturl := strings.Split(url, "/")

	filename := spliturl[1]

	splitfilename := strings.Split(filename, ".")

	return splitfilename[0]
}

func DownloadRepository(repos chan string, path string, i int) {

	if len(repos) == 0 {

		fmt.Println("Goroutine", i, "ended OK, No more Repositories to clone")
		return

	} else {

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		defer cancel()

		url := <-repos

		if _, err := git.PlainCloneContext(ctx, path+"/"+GetName(url), false, &git.CloneOptions{
			URL:      url,
			Progress: os.Stdout,
		}); err != nil {
			fmt.Println("---------------------")
			fmt.Println("URL =", url, "\nPath =", path, "\nClone Result = KO", "\nError =", err)
			fmt.Println("---------------------")
		} else {
			fmt.Println("---------------------")
			fmt.Println("URL =", url, "\nPath =", path, "\nClone Result = OK")
			fmt.Println("---------------------")
		}

		time.Sleep(5 * time.Second)

		go DownloadRepository(repos, path, i)
	}
}

func HandleGitAllClone(c *fiber.Ctx) error {

	// get the Collection from the request body
	jsonBody := new(Repositories)

	// validate the request body
	if err := c.BodyParser(jsonBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	routines, _ := strconv.Atoi(jsonBody.Goroutines)

	/* timeout, _ := time.ParseDuration(jsonBody.Timeout) */

	path := jsonBody.Path

	repos := make(chan string, len(jsonBody.Repositories))

	for _, repo := range jsonBody.Repositories {
		repos <- repo.URL
	}

	for i := 0; i < routines; i++ {
		go DownloadRepository(repos, path, i)
	}

	return c.Status(200).JSON(fiber.Map{"Operation Result": "OK"})
}
