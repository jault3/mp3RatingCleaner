package main

import (
	"errors"
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"strings"

	"github.com/DHowett/go-plist"
)

var (
	iTunesPath string
	rating     int
	disliked   bool
)

type Library struct {
	Tracks map[string]Track `plist:"Tracks"`
}

type Track struct {
	Disliked       bool   `plist:"Disliked"`
	Rating         int    `plist:"Rating"`
	RatingComputed bool   `plist:"Rating Computed"`
	Artist         string `plist:"Artist"`
	Name           string `plist:"Name"`
	Location       string `plist:"Location"`
}

func main() {
	flag.StringVar(&iTunesPath, "itunes", "", "Path to the iTunes XML file in plist format")
	flag.IntVar(&rating, "rating", -1, "All songs with the given rating will be removed [1 - 5]")
	flag.BoolVar(&disliked, "disliked", false, "Also remove all songs that have been disliked (This does not affect the rating removal process)")
	flag.Parse()
	if iTunesPath == "" {
		panic(errors.New("Invalid iTunes path: -itunes is required"))
	}
	if rating != -1 && (rating < 1 || rating > 5) {
		panic(errors.New("Invalid rating: must be between 1 and 5"))
	}
	itunesFile, err := os.Open(iTunesPath)
	if err != nil {
		panic(err)
	}
	defer itunesFile.Close()
	decoder := plist.NewDecoder(itunesFile)
	var l Library
	err = decoder.Decode(&l)
	if err != nil {
		panic(err)
	}
	removing := []Track{}

	for _, t := range l.Tracks {
		if disliked && t.Disliked {
			removing = append(removing, t)
		}
		if !t.RatingComputed && t.Rating == rating*20 {
			removing = append(removing, t)
		}
	}
	for i, t := range removing {
		fmt.Printf("Removing %d/%d songs (%s - %s)\n", i+1, len(removing), t.Name, t.Artist)
		mp3Path, err := url.QueryUnescape(strings.TrimPrefix(t.Location, "file://"))
		if err != nil {
			fmt.Printf("Error parsing file location: %s - skipping: %s\n", t.Location, err)
			continue
		}
		cmd := exec.Command("trash", mp3Path)
		cmd.Env = os.Environ()
		cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		err = cmd.Run()
		if err != nil {
			fmt.Printf("Error removing file %s: %s\n", mp3Path, err)
			continue
		}
	}
	if len(removing) == 0 {
		fmt.Println("Nothing to remove")
		return
	}
}
