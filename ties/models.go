package ties

import (
	"log"

	"database/sql"
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

func InsertTie(stmt *sql.Stmt, t Tie) {
	_, err := stmt.Exec(t.Login, t.GameName, t.BetBanker, t.BetPlayer, t.BetTie, t.TotalPayout, t.GameNumber, t.GameTime, t.DealerCards, t.PlayerCards)
	if err != nil {
		log.Fatal(err.Error())
	}
}
