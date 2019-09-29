package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type mySQLDatabase struct {
	db *sql.DB
}

func (db *mySQLDatabase) Open() {
	db.db, _ = sql.Open("mysql", "kosev:veliko123@/statser?charset=utf8")
}

func (db *mySQLDatabase) Close() {
	db.db.Close()
}

func (db *mySQLDatabase) insertUser(username, psw, email string) {
	date := time.Now().Format("2006-01-02 15:04:05")
	hash, _ := hashPassword(psw)
	stmt, err := db.db.Prepare("INSERT users SET username=?,psw=?, email=?, date=?")
	if err != nil {
		log.Println("statement failed:", err)
	}

	_, err = stmt.Exec(username, hash, email, date)
	if err != nil {
		log.Println("failed to insert user data:", err)
	}
}

func (db *mySQLDatabase) getUsername(s string) (string, error) {
	var result string
	err := db.db.QueryRow("SELECT username FROM users WHERE username=?", s).Scan(&result)
	if err != nil {
		return result, err
	}
	return result, nil

}

// Gets the password of the user given a username
func (db *mySQLDatabase) getPassword(user string) (string, error) {
	var psw string
	err := db.db.QueryRow("SELECT psw FROM users where username=?", user).Scan(&psw)
	if err != nil {
		log.Println("Error getting the password", err)
		return psw, err
	}
	return psw, nil
}

func (db *mySQLDatabase) insertMatch(league, date, homeTeam, awayTeam string, homeTeamGoals, awayTeamGoals int) {
	stmt, err := db.db.Prepare("INSERT match_data SET league=?,date=?, home_team=?, away_team=?, home_team_goals=?, away_team_goals=?")
	if err != nil {
		log.Println("statement failed:", err)
	}

	_, err = stmt.Exec(league, date, homeTeam, awayTeam, homeTeamGoals, awayTeamGoals)
	if err != nil {
		log.Println("failed to insert match data:", err)
	}

}

func (db *mySQLDatabase) getLike(keyword string) []string {
	s := []string{}
	result, err := db.db.Query("SELECT DISTINCT home_team FROM match_data WHERE home_team LIKE '%" + keyword + "%'")
	if err != nil {
		log.Println(err)
	}
	for result.Next() {
		var q string
		err = result.Scan(&q)
		fmt.Println(q)
		s = append(s, q)
		if err != nil {
			log.Println("Error:", err)
		}

	}
	fmt.Println(s)
	return s

}
