package input

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/browserutils/kooky"
	"github.com/browserutils/kooky/browser/safari"
)

func Fetch(year int, day int) string {
	dir, _ := os.UserHomeDir()
	cookiesFile := dir + "/Library/Containers/com.apple.Safari/Data/Library/Cookies/Cookies.binarycookies"

	cookies, err := safari.ReadCookies(
		context.TODO(), 
		cookiesFile, 
		kooky.Valid, 
		kooky.Domain(`.adventofcode.com`), 
		kooky.Name(`session`),
	)

	if err != nil {
		panic(err)
	}

	if len(cookies) == 0 {
		panic("no session cookie found")
	}

	sessionCookie := cookies[0].Value

	url := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)
	res, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	res.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	})
	res.Header.Set("User-Agent", "github.com/nairvarun/coding-challenges/AdventOfCode by nairvarun104@gmail.com")

	resp, err := http.DefaultClient.Do(res)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		panic(fmt.Sprintf("failed to fetch input: %s", resp.Status))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return string(body)
}
