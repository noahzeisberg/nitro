package main

import (
	"context"
	"github.com/NoahOnFyre/gengine/color"
	"github.com/NoahOnFyre/gengine/filesystem"
	"github.com/NoahOnFyre/gengine/logging"
	"github.com/google/go-github/github"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func RegisterCommand(name string, description string, args []string, runnable func([]string)) {
	arguments := Args{
		Count: len(args),
		Get:   args,
	}
	commands = append(commands, Command{
		Name:        name,
		Description: description,
		Args:        arguments,
		Run:         runnable,
	})
}

func GetCommand(args []string) {
	repository := ParseRepository(args[0])

	packageName := ToPackageName(repository)

	owner, name := SplitRepositoryID(repository)

	logging.Log("Checking package " + color.Green + repository + color.Reset + "...")
	if !filesystem.Exists(pkgDir + "\\" + packageName + "\\") {
		err := os.MkdirAll(pkgDir+"\\"+packageName+"\\", os.ModeDir)
		if err != nil {
			logging.Error("Failed to create local package folder:", err)
			return
		}
	} else {
		logging.Error("Package already downloaded!")
		return
	}

	_, directoryContent, response, err := gitHubClient.Repositories.GetContents(context.Background(), owner, name, "/", &github.RepositoryContentGetOptions{})
	if err != nil {
		logging.Error("Failed to make request to GitHub API:", err)
		return
	}

	logging.Log("Starting download of package " + color.Green + repository + color.Reset + " via " + color.Green + response.Proto)

	var wg sync.WaitGroup
	wg.Add(len(directoryContent))

	logging.Log("Downloading... This may take a while.")

	for _, fileProperties := range directoryContent {
		go func(props *github.RepositoryContent) {
			if *props.Type != "file" {
				wg.Done()
				return
			}
			file, _ := os.Create(pkgDir + "\\" + packageName + "\\" + *props.Name)

			logging.Log("Fetching " + color.Green + *props.Name + "... " + color.Gray + "(" + strconv.Itoa(props.GetSize()) + " Bytes" + ")" + color.Gray + " - " + color.Green + response.Status)

			res, err := http.Get(*props.DownloadURL)
			if err != nil {
				logging.Error("GET request failed: " + err.Error())
				return
			}

			defer res.Body.Close()

			body, err := io.ReadAll(res.Body)

			if err != nil {
				logging.Error("Reading body failed: " + err.Error())
				return
			}

			_, err = file.Write(body)
			if err != nil {
				logging.Error("Error while writing file: " + err.Error())
				return
			}
			file.Close()
			wg.Done()
		}(fileProperties)
	}
	wg.Wait()

	logging.Log("Successfully collected package " + color.Green + repository + color.Reset + "!")
}

func RemoveCommand(args []string) {
	repository := ParseRepository(args[0])
	packageName := ToPackageName(repository)

	logging.Log("Removing package " + color.Green + repository + color.Reset + "...")

	err := os.RemoveAll(pkgDir + "\\" + packageName)
	if err != nil {
		logging.Error("Failed to remove package:", err)
		return
	}

	logging.Log("Successfully removed package from local storage!")
}

func HelpCommand(args []string) {
	for _, props := range commands {
		logging.Log(color.Green + strings.ToUpper(props.Name) + color.Green + " - " + color.Reset + props.Description)
	}
}

func ListCommand(args []string) {
	dirContent, err := os.ReadDir(pkgDir)

	if err != nil {
		logging.Error("The package directory couldn't be read.")
	}

	for _, pkg := range dirContent {
		if pkg.IsDir() {
			pkgInfo, err := pkg.Info()

			if err != nil {
				logging.Error("The file information couldn't be read.")
			}
			logging.Log(color.Green + strings.Replace(pkg.Name(), ".", "/", -1) + color.Gray + " - " + color.Reset + strconv.Itoa(int(pkgInfo.Size())) + " Bytes" + color.Gray + " - " + color.Reset + pkgInfo.ModTime().Format("15:04:05 - 02.01.2006"))
		}
	}
}

func ExitCommand(args []string) {
	Exit(0)
}
