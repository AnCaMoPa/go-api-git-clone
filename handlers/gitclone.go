package handlers

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Repositorios struct {
	Repositorios []struct {
		URL string `json:"url"`
	} `json:"repositorios"`
}

var path string = os.Getenv("PATH")

func GetName(url string) string {

	spliturl := strings.Split(url, "/")

	filename := spliturl[1]

	splitfilename := strings.Split(filename, ".")

	return splitfilename[0]
}

func DownloadRepository(url string) {

	name := GetName(url)

	fmt.Println(name)

	cmd := exec.Command("git", "clone", url)
	cmd.Run()

	cmd = exec.Command("mv", "./"+name, path)
	cmd.Run()

}

func HandleGitAllClone(c *fiber.Ctx) error {

	// get the Collection from the request body
	jsonBody := new(Repositorios)

	// validate the request body
	if err := c.BodyParser(jsonBody); err != nil {
		return c.Status(400).JSON(fiber.Map{"bad input": err.Error()})
	}

	for _, repo := range jsonBody.Repositorios {
		go DownloadRepository(repo.URL)
	}

	return c.Status(200).JSON(fiber.Map{"internal server error": "OK"})
}
