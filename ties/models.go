package ties

import (
	"log"

	"github.com/avecost/ibt/config"
)

type Tie struct {
	Id          int
	Login       string
	GameName    string
	BetBanker   float64
	BetPlayer   float64
	BetTie      float64
	TotalPayout float64
	GameNumber  string
	GameTime    string
	DealerCards string
	PlayerCards string
}

/*

Ex: how to get count of rows

func Rows(strSQL string) (count int) {
	row := db.QueryRow("SELECT COUNT(*) FROM table_name").Scan(&count)
}

*/

func InsertTie(t Tie) {
	stmt, err := config.DB.Prepare("INSERT INTO ties (login, game_name, bet_banker, bet_player, bet_tie, total_payout, game_number, game_time, dealer_cards, player_cards) VALUES(?,?,?,?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(t.Login, t.GameName, t.BetBanker, t.BetPlayer, t.BetTie, t.TotalPayout, t.GameNumber, t.GameTime, t.DealerCards, t.PlayerCards)
	if err != nil {
		log.Fatal(err.Error())
	}
}
