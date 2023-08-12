package main

import (
	"flag"
	"fmt"
	"github.com/junjiefly/welcome"
	"os"
	"strings"
)

var srcType = "tv"
var videoDir = ""
var subtitleDir = ""
var debug = false

type ret struct {
	prefix  string
	matched string
}

var force = false
var subtitleSuffix map[string]struct{}
var videoSuffix map[string]struct{}

var subtitles = []string{
	"srt", "ass", "vtt",
	"ssa", "sub",
}

var videos = []string{
	"3gp", "avi", "flv",
	"mp4", "m4v", "mkv",
	"mov", "mpg", "mpeg",
	"rmvb", "swf",
	"webm", "wmv", "vob",
}

func init() {
	flag.StringVar(&srcType, "type", "tv", "tv or movie")
	flag.StringVar(&videoDir, "dir_v", "./", "video dir")
	flag.StringVar(&subtitleDir, "dir_s", "./", "subtitle dir")
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.BoolVar(&force, "f", false, "force executing")
	tvName = make(map[int]ret)
	subtitleSuffix = make(map[string]struct{})
	videoSuffix = make(map[string]struct{})
	for _, v := range subtitles {
		subtitleSuffix[v] = struct{}{}
	}

	for _, v := range videos {
		videoSuffix[v] = struct{}{}
	}

}

func isVideo(suffix string) bool {
	_, ok := videoSuffix[strings.ToLower(suffix)]
	return ok
}

func isSubtitle(suffix string) bool {
	_, ok := subtitleSuffix[strings.ToLower(suffix)]
	return ok
}

func rename(src, dst string) error {
	return os.Rename(src, dst)
}

func Usage() {
	fmt.Println("Usage: ./subtitleRenamer -type=movie -dir_v=./TheWanderingEarth/ -dir_s=./TheWanderingEarth/  [for movie]")
	fmt.Println("       ./subtitleRenamer -type=tv -dir_v=./TheWalkingDead/S01/ -dir_s=./TheWalkingDead/S01/  [for tv]")
}
func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] == "help" {
		Usage()
		return
	}
	welcome.Print()
	if srcType != "tv" && srcType != "movie" {
		fmt.Println("rename type has to be specified correctly, tv or movie?")
		fmt.Scanln(&srcType)
		if strings.ToLower(srcType) != "tv" || strings.ToLower(srcType) == "movie" {
			Usage()
			return
		}
	}
	var userInput string
	if force == false {
		fmt.Println("warning! this tool will change name for subtitle files!")
		fmt.Println("	1: make sure you know this changes;")
		fmt.Println("	2: you should have the read-write permission of subtitle files.")
		fmt.Print("\ninput yes or y to continue:")
		fmt.Scanln(&userInput)
		if strings.ToLower(userInput) == "yes" || strings.ToLower(userInput) == "y" {
			fmt.Println("")
		} else {
			fmt.Println("exit with doing nothing")
			return
		}
	}
	if srcType == "tv" {
		RenameTV()
	} else {
		RenameMovie()
	}
}
