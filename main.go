package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	flags "github.com/jessevdk/go-flags"
	"github.com/mackerelio/checkers"
)

type Result struct {
	Name    string
	Age     int64
	ModTime time.Time
	Status  checkers.Status
	Message string
}

var opts struct {
	Base        string `short:"b" long:"base" required:"true" description:"the base directory(required)"`
	WarningAge  int64  `short:"w" long:"warning-age" default:"21600" description:"warning if more old than(sec)"`
	CriticalAge int64  `short:"c" long:"critical-age" default:"43200" description:"critical if more old than(sec)"`
	Debug       bool   `short:"d" long:"debug" description:"debug print"`
}

func main() {
	ckr := run(os.Args[1:])
	ckr.Name = "FileAge"
	ckr.Exit()
}

func run(args []string) *checkers.Checker {
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		log.Printf("parse args: %v", err)
		os.Exit(1)
	}

	if opts.Debug {
		log.Printf("Base: %v\n", opts.Base)
		log.Printf("WarningAge: %v\n", opts.WarningAge)
		log.Printf("CriticalAge: %v\n", opts.CriticalAge)
	}

	flist, err := path(opts.Base)
	if err != nil {
		log.Printf("list file: %v", err)
		os.Exit(1)
	}

	if len(flist) < 1 {
		return checkers.Ok("Directory is empty")
	}

	if opts.Debug {
		for _, f := range flist {
			log.Println(f)
		}
	}

	results := []Result{}
	for _, f := range flist {
		stat, err := os.Stat(f)
		if err != nil {
			log.Printf("No such file. skip: %v", err)
			continue
		}
		mtime := stat.ModTime()
		age := time.Now().Unix() - mtime.Unix()

		result := checkers.OK
		if age > opts.WarningAge {
			result = checkers.WARNING
		}
		if age > opts.CriticalAge {
			result = checkers.CRITICAL
		}

		msg := fmt.Sprintf(
			"%s is %d seconds old (%04d-%02d-%02d %02d:%02d:%02d).",
			stat.Name(), age,
			mtime.Year(), mtime.Month(), mtime.Day(),
			mtime.Hour(), mtime.Minute(), mtime.Second())

		res := Result{
			Name:    stat.Name(),
			Age:     age,
			ModTime: mtime,
			Status:  result,
			Message: msg,
		}
		results = append(results, res)

		if opts.Debug {
			log.Printf("%v\n", res)
		}
	}

	for _, r := range results {
		if r.Status != checkers.OK {
			return checkers.NewChecker(r.Status, r.Message)
		}
	}

	msg := fmt.Sprintf("path: %s", opts.Base)
	return checkers.NewChecker(checkers.OK, msg)
}

func path(dir string) ([]string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	paths := []string{}
	for _, file := range files {
		if file.IsDir() {
			path, err := path(filepath.Join(dir, file.Name()))
			if err != nil {
				return paths, err
			}
			paths = append(paths, path...)
			continue
		}

		paths = append(paths, filepath.Join(dir, file.Name()))
	}

	return paths, nil
}

