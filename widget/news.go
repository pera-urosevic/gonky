package widget

import (
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/pera-urosevic/gonky/runner"
	"github.com/pera-urosevic/gonky/util"
)

// News //
type News struct {
	URL        string
	EverySlide time.Duration
	EveryFetch time.Duration
	index      int
	titles     []string
	State      string
	Render     func(*News)
}

// Start //
func (news News) Start() {
	go runner.Run(news.EveryFetch, news.fetch)
	go runner.RunAfter(news.EverySlide, news.slide)
}

func (news *News) fetch() {
	var e error
	news.titles, e = news.sensor(news.URL)
	if e != nil {
		util.ErrorLog(e)
		return
	}
	news.index = 0
	news.Render(news)
}

func (news *News) sensor(url string) ([]string, error) {
	if util.Debug {
		sensor := []string{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
			"Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
			"Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.",
		}
		return sensor, nil
	}
	fp := gofeed.NewParser()
	feed, e := fp.ParseURL(url)
	if e != nil {
		return nil, e
	}
	sensor := []string{}
	for _, item := range feed.Items {
		sensor = append(sensor, item.Title)
	}
	return sensor, nil
}

func (news *News) slide() {
	news.index = (news.index + 1) % len(news.titles)
	news.State = news.titles[news.index]
	news.Render(news)
}
