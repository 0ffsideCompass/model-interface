package modelinterface

import "time"

type JSONFixtureArticle struct {
	ID         string          `json:"id,omitempty"`
	Title      string          `json:"title"`
	Body       string          `json:"body"`
	HashTags   []string        `json:"hashtags"`
	Date       time.Time       `json:"date"`
	Categories []string        `json:"categories"`
	Data       JSONFixtureData `json:"data"`
	Summary    string          `json:"summary"`
}

type JSONFixtureData struct {
	League                 string `json:"league"`
	TeamOne                string `json:"team_one"`
	TeamTwo                string `json:"team_two"`
	TeamOneWins            int    `json:"team_one_wins"`
	TeamTwoWins            int    `json:"team_two_wins"`
	TeamOneDraws           int    `json:"team_one_draws"`
	TeamTwoDraws           int    `json:"team_two_draws"`
	TeamOneLoses           int    `json:"team_one_loses"`
	TeamTwoLoses           int    `json:"team_two_loses"`
	TeamOneGoals           int    `json:"team_one_goals"`
	TeamTwoGoals           int    `json:"team_two_goals"`
	TeamOneGoalsAgainst    int    `json:"team_one_goals_against"`
	TeamTwoGoalsAgainst    int    `json:"team_two_goals_against"`
	TeamOneCleanSheets     int    `json:"team_one_clean_sheets"`
	TeamTwoCleanSheets     int    `json:"team_two_clean_sheets"`
	TeamOneFailedToScore   int    `json:"team_one_failed_to_score"`
	TeamTwoFailedToScore   int    `json:"team_two_failed_to_score"`
	TeamOnePenaltiesScored int    `json:"team_one_penalties_scored"`
	TeamTwoPenaltiesScored int    `json:"team_two_penalties_scored"`
}

type GetFixtureResponse struct {
	Referee         string       `json:"referee"`
	Venue           string       `json:"venue"`
	City            string       `json:"city"`
	Status          string       `json:"status"`
	StatusLong      string       `json:"status_long"`
	HomeTeam        string       `json:"home_team"`
	AwayTeam        string       `json:"away_team"`
	HomeGoals       int          `json:"home_goals"`
	AwayGoals       int          `json:"away_goals"`
	HomeTeamLogo    string       `json:"home_team_logo"`
	AwayTeamLogo    string       `json:"away_team_logo"`
	League          string       `json:"league"`
	LeagueCountry   string       `json:"league_country"`
	LeagueLogo      string       `json:"league_logo"`
	CountryFlag     string       `json:"country_flag"`
	TimeElapsed     int          `json:"time_elapsed"`
	HomeLineup      LinueupStuct `json:"home_lineup"`
	AwayLineup      LinueupStuct `json:"away_lineup"`
	HomeCoach       string       `json:"home_coach"`
	AwayCoach       string       `json:"away_coach"`
	Events          []Event      `json:"events"`
	Statistics      []Statistic  `json:"statistics"`
	HomePlayerStats PlayerStats  `json:"home_player_stats"`
	AwayPlayerStats PlayerStats  `json:"away_player_stats"`
}

type LinueupStuct struct {
	Formation   string   `json:"formation"`
	StartXI     []Player `json:"start_xi"`
	Substitutes []Player `json:"substitutes"`
}

type Event struct {
	Time struct {
		Elapsed int         `json:"elapsed"`
		Extra   interface{} `json:"extra"`
	} `json:"time"`
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	Player struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"player"`
	Assist struct {
		ID   interface{} `json:"id"`
		Name interface{} `json:"name"`
	} `json:"assist"`
	Type     string `json:"type"`
	Detail   string `json:"detail"`
	Comments string `json:"comments"`
}

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Number int    `json:"number"`
	Pos    string `json:"pos"`
	Grid   string `json:"grid"`
}

type Statistic struct {
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Logo string `json:"logo"`
	} `json:"team"`
	Statistics []struct {
		Type  string      `json:"type"`
		Value interface{} `json:"value"`
	} `json:"statistics"`
}

type PlayerStats struct {
	Team struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"team"`
	Players []PlayerStat `json:"statistics"`
}

type PlayerStat struct {
	Player struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Photo string `json:"photo"`
	} `json:"player"`
	Statistics []struct {
		Games struct {
			Minutes    int    `json:"minutes"`
			Number     int    `json:"number"`
			Position   string `json:"position"`
			Rating     string `json:"rating"`
			Captain    bool   `json:"captain"`
			Substitute bool   `json:"substitute"`
		} `json:"games"`
		Offsides interface{} `json:"offsides"`
		Shots    struct {
			Total interface{} `json:"total"`
			On    interface{} `json:"on"`
		} `json:"shots"`
		Goals struct {
			Total    interface{} `json:"total"`
			Conceded int         `json:"conceded"`
			Assists  interface{} `json:"assists"`
			Saves    int         `json:"saves"`
		} `json:"goals"`
		Passes struct {
			Total    int         `json:"total"`
			Key      interface{} `json:"key"`
			Accuracy string      `json:"accuracy"`
		} `json:"passes"`
		Tackles struct {
			Total         interface{} `json:"total"`
			Blocks        interface{} `json:"blocks"`
			Interceptions interface{} `json:"interceptions"`
		} `json:"tackles"`
		Duels struct {
			Total interface{} `json:"total"`
			Won   interface{} `json:"won"`
		} `json:"duels"`
		Dribbles struct {
			Attempts interface{} `json:"attempts"`
			Success  interface{} `json:"success"`
			Past     interface{} `json:"past"`
		} `json:"dribbles"`
		Fouls struct {
			Drawn     interface{} `json:"drawn"`
			Committed interface{} `json:"committed"`
		} `json:"fouls"`
		Cards struct {
			Yellow int `json:"yellow"`
			Red    int `json:"red"`
		} `json:"cards"`
		Penalty struct {
			Won      interface{} `json:"won"`
			Commited interface{} `json:"commited"`
			Scored   int         `json:"scored"`
			Missed   int         `json:"missed"`
			Saved    int         `json:"saved"`
		} `json:"penalty"`
	}
}
