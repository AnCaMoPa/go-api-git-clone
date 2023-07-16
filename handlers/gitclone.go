package handlers

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Repositorios struct {
	Repositorios []struct {
		URL string `json:"url"`
	} `json:"repositorios"`
}

const destiny string = "path in which you want to have your repositories"

func GetName(url string) string {

	spliturl := strings.Split(url, "/")

	filename := spliturl[1]

	splitfilename := strings.Split(filename, ".")

	return splitfilename[0]
}

func DownloadRepository(repos chan string) string {

	if len(repos) == 0 {

		fmt.Println("No more Repositories to clone")
		return ""

	} else {

		url := <-repos

		name := GetName(url)

		/* fmt.Println(url) */

		cmd := exec.Command("git", "clone", url)

		if err := cmd.Run(); err != nil {

			fmt.Println("Git Repository :", url, "Ko Result :", err.Error())

		} else {
			fmt.Println(destiny)
			cmd = exec.Command("mv", "./"+name, destiny)

			if err := cmd.Run(); err != nil {

				fmt.Println("Git Repository :", url, "Ko Result :", err.Error())

			}

		}

		time.Sleep(2 * time.Second)

		go DownloadRepository(repos)
	}

	return ""

}

func HandleGitAllClone(c *fiber.Ctx) error {

	routines, _ := strconv.Atoi(c.Params("routines"))

	// get the Collection from the request body
	jsonBody := new(Repositorios)

	// validate the request body
	if err := c.BodyParser(jsonBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	repos := make(chan string, len(jsonBody.Repositorios))

	for _, repo := range jsonBody.Repositorios {
		repos <- repo.URL
	}

	/* 	wg.Add(routines)
	 */
	for i := 0; i < routines; i++ {
		go DownloadRepository(repos)
	}

	return c.Status(200).JSON(fiber.Map{"Operation Result": "OK"})
}
