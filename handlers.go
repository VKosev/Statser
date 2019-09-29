package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("username")
	if err != nil {
		tpl.ExecuteTemplate(w, "home", nil)
	} else {
		tplData.Username = cookie.Value
		tpl.ExecuteTemplate(w, "home", tplData)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.Redirect(w, r, "/", 302)
	case "POST":
		r.ParseForm()
		err := checkLogin(r.FormValue("username"), r.FormValue("password"))
		if err != nil {
			tplData.LoginError = err.Error()
			tpl.ExecuteTemplate(w, "home", tplData)
		} else {
			cookie := &http.Cookie{
				Name:   "username",
				Value:  r.FormValue("username"),
				MaxAge: 6000,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", 302)
		}
	}
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	// Deleting the cookie and redirecting to the index page.
	cookie := &http.Cookie{
		Name:   "username",
		Value:  r.FormValue("username"),
		MaxAge: 0,
	}

	http.SetCookie(w, cookie)
	tplData.LoginError = ""
	http.Redirect(w, r, "/", 302)
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		tpl.ExecuteTemplate(w, "register", nil)

	case "POST":
		// Check if username is valid
		_, err := validateUser(r.FormValue("username"))

		if err != nil {
			tplData.RegError = err.Error()
			tpl.ExecuteTemplate(w, "register", tplData)

			// Check if password is valid
		} else if err = validatePsw(r.FormValue("password")); err != nil {
			fmt.Println(r.FormValue("passwrd"))
			tplData.RegError = err.Error()
			tpl.ExecuteTemplate(w, "register", tplData)
		} else {
			// If everything is valid insert the data and create cookie
			db.insertUser(r.FormValue("username"), r.FormValue("password"), r.FormValue("email"))

			cookie := &http.Cookie{
				Name:   "username",
				Value:  r.FormValue("username"),
				MaxAge: 6000,
			}
			http.SetCookie(w, cookie)
			http.Redirect(w, r, "/", 302)
		}
	}
}

func hthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var keyword string
		data := r.FormValue("team1")

		if strings.Count(data, "") > 2 {
			result, err := db.db.Query("SELECT DISTINCT home_team FROM match_data WHERE home_team LIKE '%" + data + "%'")
			if err != nil {
				log.Println(err)
			}

			var response string

			for result.Next() {
				err = result.Scan(&keyword)
				response += `<a href="#" class="list-group-item list-group-item-action py-1" onclick ="fillTeam1('` + keyword + `')">` + keyword + `</a>`

			}
			w.Write([]byte(response))

		}

		data = r.FormValue("team2")
		if strings.Count(data, "") > 2 {
			result, err := db.db.Query("SELECT DISTINCT home_team FROM match_data WHERE home_team LIKE '%" + data + "%'")
			if err != nil {
				log.Println(err)
			}

			var response string

			for result.Next() {
				err = result.Scan(&keyword)
				response += `<a href="#" class="list-group-item list-group-item-action py-1" onclick ="fillTeam2('` + keyword + `')">` + keyword + `</a>`

			}
			w.Write([]byte(response))
		}
	case "GET":
		var team1 teamData
		var team2 teamData
		league := "English Premier League " + r.FormValue("season")
		inputTeam1 := r.FormValue("team1")
		inputTeam2 := r.FormValue("team2")
		fmt.Println(league, inputTeam1, inputTeam2)
		//var matchData []matchData

		result, err := db.db.Query("SELECT league, date, home_team, away_team, home_team_goals, away_team_goals FROM match_data WHERE home_team = ? AND away_team = ? AND league = ? OR home_team = ? AND away_team = ? AND league = ?", inputTeam1, inputTeam2, league, inputTeam2, inputTeam1, league)
		if err != nil {
			log.Println("error fetching data for match_data:", err)
		}

		for result.Next() {
			var league string
			var date string
			var homeTeam string
			var awayTeam string
			var homeTeamGoals int
			var awayTeamGoals int
			err = result.Scan(&league, &date, &homeTeam, &awayTeam, &homeTeamGoals, &awayTeamGoals)
			//log.Println(err)

			if homeTeamGoals > awayTeamGoals {
				team1.Wins++
				team2.Loses++
			}

			team1.Name = homeTeam
			fmt.Println("team1:", team1.Name)
			fmt.Println("team1 wins:", team1.Wins)
			team2.Name = awayTeam
			fmt.Println("team2:", team2.Name)
			fmt.Println("team2 wins:", team2.Wins)
		}
	}

}
