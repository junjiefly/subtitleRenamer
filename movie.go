package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"strings"
)

var movieName ret

func readMovieVideoFile() error {
	movieName.matched = "no"
	videoFiles, err := ioutil.ReadDir(videoDir)
	if err != nil {
		fmt.Println("fail to read video dir:", err, "videoDir:", videoDir)
		return err
	}
	for _, videoFile := range videoFiles {
		if !videoFile.IsDir() {
			name := videoFile.Name()
			idx := strings.LastIndex(name, ".")
			if idx < 0 {
				fmt.Println("invalid video name")
				continue
			}
			suffix := name[idx+1:]
			if !isVideo(suffix) {
				continue
			}
			vname := name[:idx]
			movieName.prefix = vname
			if debug {
				fmt.Println("find video file:", vname)
			}
			break
		}
	}
	return nil
}

func readMovieSubtitleFile() error {
	subtitleFiles, err := ioutil.ReadDir(subtitleDir)
	if err != nil {
		fmt.Println("fail to read subtitle dir:", err, "subtitleDir:", videoDir)
		return err
	}
	for _, subtitleFile := range subtitleFiles {
		if !subtitleFile.IsDir() {
			name := subtitleFile.Name()
			idx := strings.LastIndex(name, ".")
			if idx < 0 {
				fmt.Println("invalid subtitle name:", subtitleFile.Name())
				continue
			}
			if !isSubtitle(name[idx+1:]) {
				continue
			}
			idx = strings.LastIndex(name[:idx], ".")
			if idx < 0 {
				fmt.Println("invalid subtitle name", subtitleFile.Name())
				continue
			}
			suffix := strings.ToLower(name[idx:])
			newName := movieName.prefix + suffix
			movieName.matched = "yes"
			oldFullpath := path.Join(subtitleDir, name)
			newFullpath := path.Join(subtitleDir, newName)
			err = rename(oldFullpath, newFullpath)
			if err != nil {
				fmt.Println("fail to rename subtitle file:", name, "to:", newName)
				continue
			}
			if debug {
				fmt.Println("match subtitle for movie:", movieName.prefix)
			}
			if debug {
				fmt.Println("fail to match subtitle file:", name)
			}
		}
	}
	return nil
}

func RenameMovie() {
	if readMovieVideoFile() != nil {
		return
	}
	if readMovieSubtitleFile() != nil {
		return
	}
	if movieName.prefix != "" {
		fmt.Printf("find move: %s, subtitle matched? %s \n", movieName.prefix, movieName.matched)
	} else {
		fmt.Printf("no movie was found in this directory")
	}
}
