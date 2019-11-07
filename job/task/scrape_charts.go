package task

import "github.com/varmamsp/cello/app"

type ScrapeCharts struct {
	*app.App
}

func NewScrapeCharts(app *app.App) (*ScrapeCharts, error) {
	return &ScrapeCharts{app}, nil
}

func (s *ScrapeCharts) Call() {
	_ = map[string]string{
		"all": "https://chartable.com/charts/itunes/us-all-podcasts-podcasts",
	}
}
