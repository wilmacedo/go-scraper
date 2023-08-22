package cli

import (
	"flag"
	"fmt"
	"go-scraper/src/lib"
	"os"
	"runtime"
	"time"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("	run -job [JOB] - Run determinate job to process and save data")
}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) runFetcher() {
	fmt.Printf("Fetching news data...\n")
	start := time.Now()

	_, err := lib.FetchNews(lib.SearchQuery{
		Subject: "startups AND Latin American",
	})
	if err != nil {
		panic(err)
	}

	duration := time.Since(start).Round(time.Millisecond / 100)

	fmt.Printf("Done fetching news in %s\n", duration)
}

func (cli *CommandLine) Run() {
	cli.validateArgs()

	runnerCmd := flag.NewFlagSet("run", flag.ExitOnError)

	job := runnerCmd.String("job", "", "Provide job name to run")

	switch os.Args[1] {
	case "run":
		err := runnerCmd.Parse(os.Args[2:])
		if err != nil {
			panic(err)
		}
	default:
		cli.printUsage()
		runtime.Goexit()
	}

	if runnerCmd.Parsed() {
		switch *job {
		case "fetcher":
			cli.runFetcher()
		default:
			runnerCmd.Usage()
			runtime.Goexit()
		}
	}
}
