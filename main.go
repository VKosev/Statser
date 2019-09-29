package main

import (
	"html/template"
	"net/http"
)

type teamData struct {
	Name  string
	Wins  int
	Loses int
	Draws int
}

type matchData struct {
	League        string
	Date          string
	HomeTeam      string
	AwayTeam      string
	HomeTeamGoals int
	AwayTeamGoals int
}

// This struct is used to pass data around the html templates.
type templateData struct {
	Username   string
	RegError   string
	LoginError string
	Team1      teamData
	team2      teamData
	MatchData  []matchData
}

var tplData = new(templateData)
var db = new(mySQLDatabase)
var s = new(Season)

// Caching all templates.
var tpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {

	db.Open()
	defer db.Close()

	// Serving static files from the assets folder like CSS and JS files.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Routing and handlers
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/register", registerHandler)
	http.HandleFunc("/hth", hthHandler)

	http.ListenAndServe(":8000", nil)
}
