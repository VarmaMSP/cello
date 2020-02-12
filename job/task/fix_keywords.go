package task

import "github.com/varmamsp/cello/app"

type FixKeywords struct {
	*app.App,
}

func NewFixKeywords(app *app.App) (*FixKeywords, error) {
	return &FixKeywords{
		App: app,
	}
}

func (s *FixKeywords) Call() {
	fmt.Println("Fix keywords started")
	go func() {
		limit := 1000
		lastId := int64(0)

		for {
			keywords, err := s.Store.Keyword().GetAllPaginated(lastId, limit)
			if err != nil {
				fmt.Println(err)
				break
			}

			for _, keyword := range keywords {
				if model.IsValidKeyword(keyword.Text) {
					if err := s.Store.Keyword.UpdateText(keyword.Id, strings.ToLower(keyword.Text)); err != nil {
						fmt.Println(err)
						break
					}
				} else {
					if err := s.Store.Keyword.Delete(keyword.Id); err != nil {
						fmt.Println(err)
						break
					}
				}
			}

			if len(keywords) < limit {
				break
			}
			lastId = keywords[len(keywords)-1]
		}
	}()
}