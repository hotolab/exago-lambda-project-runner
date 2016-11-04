package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strings"

	lambda "github.com/eawsy/aws-lambda-go/service/lambda/runtime"
	"github.com/hotolab/exago-runner/task"
)

var (
	binFolder     string
	gitFolder     string
	gitCoreFolder string

	goPath = "/tmp/gopath"
)

func handler(e json.RawMessage, ctx *lambda.Context) (interface{}, error) {
	var v map[string]string
	json.Unmarshal(e, &v)

	if _, ok := v["repository"]; !ok {
		return nil, errors.New("The repository is missing")
	}

	m := task.NewManager(v["repository"])
	// Whether we use shallow cloning or not
	if val, ok := v["shallow"]; !ok {
		if val == "1" {
			m.DoShallow()
		}
	}
	// Whether we use a reference or not
	if val, ok := v["reference"]; !ok {
		if val != "" {
			m.UseReference(val)
		}
	}

	// Cleanup gopath
	cleanup()

	return m.ExecuteRunners()
}

func init() {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	binFolder = pwd + "/bin/" + runtime.GOOS + "-amd64"
	gitFolder = pwd + "/git/bin"
	gitCoreFolder = pwd + "/git/lib/git-core"

	// Set some environment variables
	paths := []string{os.Getenv("PATH"), binFolder, gitFolder, goPath + "/bin"}
	os.Setenv("PATH", strings.Join(paths, ":"))
	os.Setenv("GOGC", "off")
	os.Setenv("GOPATH", goPath)
	os.Setenv("GODEBUG", "sbrk=1")
	os.Setenv("GIT_SSL_NO_VERIFY", "1")
	os.Setenv("CGO_ENABLED", "0")
	os.Setenv("HOME", pwd)
	os.Setenv("GIT_EXEC_PATH", gitCoreFolder)

	lambda.HandleFunc(handler)
}

// cleanup removes $GOPATH
func cleanup() {
	os.RemoveAll(goPath)
}

func main() {}
