package dto

type (
	AdverseMedia struct {
		Types              []string     `json:"types"`
		MediaEntries       []MediaEntry `json:"mediaEntries"`
		Name               string       `json:"name"`
		MatchTypes         []string     `json:"matchTypes"`
		PoliticalPositions []string     `json:"politicalPositions"`
		Country            string       `json:"country"`
	}

	CheckApplicantRaw struct {
		Answer        string  `json:"answer"`
		WatchlistInfo HitsRaw `json:"watchlistInfo"`
	}

	HitsRaw struct {
		Hits []AdverseMedia `json:"hits"`
	}

	MainResponse struct {
		PersonWatchlist CheckApplicantRaw `json:"PERSON_WATCHLIST"`
	}

	MediaEntry struct {
		Date       string `json:"date"`
		Annotation string `json:"annotation"`
		Title      string `json:"title"`
		Url        string `json:"url"`
	}
)
