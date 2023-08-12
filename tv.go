package main

import (
	"fmt"
	"io/ioutil"
	"path"
	"regexp"
	"strconv"
	"strings"
)

var tvName map[int]ret

func StringToInt(s string) int {
	c, err := strconv.Atoi(s)
	if err != nil {
		return -1
	}
	return c
}

func parseEpisode(fileName string) int {
	c := matchSE(fileName)
	if c < 0 {
		c = matchEP(fileName)
		if c < 0 {
			c = matchJiShu(fileName)
		}
	}
	return c
}

func test4() int {
	fileNames := []string{
		"S01EP132",
		"S1E134",
		"s01e3",
		"s1e03",
		"s02e04",
		"S1E0134",
		"S01E00134",
	}
	regexPattern := "(?i)s([0-9]+)e(p)?([0-9]+)"
	for _, fileName := range fileNames {
		re := regexp.MustCompile(regexPattern)
		match := re.FindStringSubmatch(fileName)
		//	fmt.Println("match:", match, "len:", len(match))
		if len(match) > 3 {
			season := match[1]
			episode := StringToInt(match[3])
			fmt.Printf("test4 %s season: %s, episode: %d\n", fileName, season, episode)
			//	return episode
		} else {
			fmt.Printf("test4 fail\n")
		}
	}
	return -1
	//}
}
func matchSE(fileName string) int {
	/*fileNames := []string{
		"S01EP132",
		"S1E134",
		"s01e3",
		"s1e03",
		"s02e04",
		"S1E0134",
		"S01E00134",
	}*/
	regex := regexp.MustCompile(`(?i)S(\d+)E(p)?(\d+)`)
	//for _, fileName := range fileNames {
	match := regex.FindStringSubmatch(fileName)
	//	fmt.Println("match:", match, "len:", len(match))
	if len(match) > 3 {
		seasonNumber := match[1]
		episodeNumber := StringToInt(match[3])
		if debug {
			fmt.Printf("matchSE fileName: %s, season: %s, episode: %d\n", fileName, seasonNumber, episodeNumber)
		}
		return episodeNumber
	} else {
		if debug {
			fmt.Printf("matchSE fileName: %s fail \n", fileName)
		}
		return -1
	}
	//	}

}

func matchEP(fileName string) int {
	/*fileNames := []string{
		"eP01",
		"EP002",
		"EP1",
		"Ep2",
		"EP124",
	}*/
	regex := regexp.MustCompile(`(?i)E(p)?(\d+)`)
	match := regex.FindStringSubmatch(fileName)
	if len(match) > 2 {
		episodeNumber := StringToInt(match[2])
		if debug {
			fmt.Printf("matchEP fileName: %s, episode: %d \n", fileName, episodeNumber)
		}
		return episodeNumber
	}
	if debug {
		fmt.Printf("matchEP fileName: %s fail\n", fileName)
	}
	return -1
}

func matchJiShu(fileName string) int {
	/*fileNames := []string{
		"第1集",
		"第135集",
		"第01集",
		"第002集",
	}*/
	regex := regexp.MustCompile(`第(\d+)集`)
	//for _, fileName := range fileNames {
	match := regex.FindStringSubmatch(fileName)
	if len(match) > 1 {
		episodeNumber := StringToInt(match[1])
		if debug {
			fmt.Printf("matchJiShu fileName: %s, episode: %d\n", fileName, episodeNumber)
		}
		return episodeNumber
	}
	if debug {
		fmt.Printf("matchJiShu fileName: %s fail\n", fileName)
	}
	return -1
	//}
}

func readTVVideoFile() error {
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
			ep := parseEpisode(vname)
			if ep > 0 {
				tvName[ep] = ret{prefix: vname, matched: "no"}
				if debug {
					fmt.Println("find video file:", name, "episode number:", ep)
				}
				continue
			}
			if debug {
				fmt.Println("can not match video file:", name)
			}
		}
	}
	return nil
}

func readTVSubtitleFile() error {
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
			vname := name[:idx]
			suffix := strings.ToLower(name[idx:])
			ep := parseEpisode(vname)
			if ep > 0 {
				cc, ok := tvName[ep]
				if ok {
					newName := cc.prefix + suffix
					oldFullpath := path.Join(subtitleDir, name)
					newFullpath := path.Join(subtitleDir, newName)
					err = rename(oldFullpath, newFullpath)
					if err != nil {
						fmt.Println("fail to rename subtitle file:", name, "to:", newName)
						continue
					}
					cc.matched = "yes"
					tvName[ep] = cc
					if debug {
						fmt.Println("match subtitle for episode:", ep)
					}
				} else {
					if debug {
						fmt.Println("fail to match video file for subtitle file:", name)
					}
				}
			}
			if debug {
				fmt.Println("fail to match subtitle file:", name)
			}
		}
	}
	return nil
}

func RenameTV() {
	if readTVVideoFile() != nil {
		return
	}
	var max = -1
	for k := range tvName {
		if max < k {
			max = k
		}
	}
	if readTVSubtitleFile() != nil {
		return
	}
	for i := 0; i <= max; i++ {
		v, ok := tvName[i]
		if ok {
			fmt.Printf("find episode: %3d, subtitle matched? %s \n", i, v.matched)
		}
	}
}
