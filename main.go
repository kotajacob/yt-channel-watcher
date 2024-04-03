package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

const baseFeedURL = `https://www.youtube.com/feeds/videos.xml?channel_id=`

type Config struct {
	Channels []Channel
}

type Channel struct {
	ID   string
	Path string
}

func main() {
	log.SetPrefix("")
	log.SetFlags(0)

	configPath := flag.String(
		"config",
		"/etc/yt-channel-watcher/config.toml",
		"Path to config",
	)
	videoPath := flag.String(
		"vids",
		"/var/www/fog/vids",
		"Path to video folder",
	)
	flag.Parse()

	config := new(Config)
	_, err := toml.DecodeFile(*configPath, config)
	if err != nil {
		log.Fatalf("failed reading %v: %v\n", *configPath, err)
	}

	// List existing videos.
	// Fetch recent RSS videos.
	// Download unmatched videos.
	for _, channel := range config.Channels {
		channelPath := filepath.Join(*videoPath, channel.Path)
		existing, err := existingVideos(channelPath)
		if err != nil {
			log.Fatalf("failed reading %v: %v\n", channelPath, err)
		}
		fmt.Println("existing", existing)
		latest, _ := latestDate(existing)
		fmt.Println("latest", latest)
		// err = fetch(channel.ID)
		// if err != nil {
		// 	log.Fatalln(err)
		// }
	}
}

// existingVideos returns a list of downloaded videos located in a path.
func existingVideos(path string) ([]string, error) {
	var videos []string
	entries, err := os.ReadDir(path)
	if err != nil {
		return videos, err
	}

	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		videos = append(videos, e.Name())
	}
	return videos, err
}

// latestDate gets the most recent date from a list of video names.
// The video names are assumed to be in the format: YYYYMMDD - Title.
// If none can be found with a valid date an error is returned.
func latestDate(vids []string) (string, error) {
	var dates []string
	for _, vid := range vids {
		date, _, ok := strings.Cut(vid, " ")
		if !ok {
			continue
		}
		_, err := time.Parse("20060102", date)
		if err != nil {
			continue
		}
		dates = append(dates, date)
	}

	if len(dates) == 0 {
		return "", fmt.Errorf("downloaded video names did not begin with a YYYYMMDD date")
	}

	sortableDates := sort.StringSlice(dates)
	sort.Sort(sort.Reverse(sortableDates))
	return sortableDates[0], nil
}
