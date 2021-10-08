package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// App export
type App struct {
	Router *mux.Router
}

type Resume struct {
	Basics       Basics         `json:"basics"`
	Work         []Work         `json:"work"`
	Volunteer    []Volunteer    `json:"volunteer"`
	Education    []Education    `json:"education"`
	Awards       []Awards       `json:"awards"`
	Publications []Publications `json:"publications"`
	Skills       []Skills       `json:"skills"`
	Languages    []Languages    `json:"languages"`
	Interests    []Interests    `json:"interests"`
	References   []References   `json:"references"`
	Projects     []Projects     `json:"projects"`
}
type Location struct {
	Address     string `json:"address"`
	PostalCode  string `json:"postalCode"`
	City        string `json:"city"`
	CountryCode string `json:"countryCode"`
	Region      string `json:"region"`
}
type Profiles struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}
type Basics struct {
	Name     string     `json:"name"`
	Label    string     `json:"label"`
	Image    string     `json:"image"`
	Email    string     `json:"email"`
	Phone    string     `json:"phone"`
	URL      string     `json:"url"`
	Summary  string     `json:"summary"`
	Location Location   `json:"location"`
	Profiles []Profiles `json:"profiles"`
}
type Work struct {
	Name       string   `json:"name"`
	Position   string   `json:"position"`
	URL        string   `json:"url"`
	StartDate  string   `json:"startDate"`
	EndDate    string   `json:"endDate"`
	Summary    string   `json:"summary"`
	Highlights []string `json:"highlights"`
}
type Volunteer struct {
	Organization string   `json:"organization"`
	Position     string   `json:"position"`
	URL          string   `json:"url"`
	StartDate    string   `json:"startDate"`
	EndDate      string   `json:"endDate"`
	Summary      string   `json:"summary"`
	Highlights   []string `json:"highlights"`
}
type Education struct {
	Institution string   `json:"institution"`
	URL         string   `json:"url"`
	Area        string   `json:"area"`
	StudyType   string   `json:"studyType"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	Score       string   `json:"score"`
	Courses     []string `json:"courses"`
}
type Awards struct {
	Title   string `json:"title"`
	Date    string `json:"date"`
	Awarder string `json:"awarder"`
	Summary string `json:"summary"`
}
type Publications struct {
	Name        string `json:"name"`
	Publisher   string `json:"publisher"`
	ReleaseDate string `json:"releaseDate"`
	URL         string `json:"url"`
	Summary     string `json:"summary"`
}
type Skills struct {
	Name     string   `json:"name"`
	Level    string   `json:"level"`
	Keywords []string `json:"keywords"`
}
type Languages struct {
	Language string `json:"language"`
	Fluency  string `json:"fluency"`
}
type Interests struct {
	Name     string   `json:"name"`
	Keywords []string `json:"keywords"`
}
type References struct {
	Name      string `json:"name"`
	Reference string `json:"reference"`
}
type Projects struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Highlights  []string `json:"highlights"`
	Keywords    []string `json:"keywords"`
	StartDate   string   `json:"startDate"`
	EndDate     string   `json:"endDate"`
	URL         string   `json:"url"`
	Roles       []string `json:"roles"`
	Entity      string   `json:"entity"`
	Type        string   `json:"type"`
}

func getResume() *Resume {
	return &Resume{
		Basics: Basics{
			Name:    "Teis Angel Clausen",
			Label:   "DevOps engineer / Sysadmin",
			Image:   "",
			Email:   "teisangel@tsrv.pw",
			Phone:   "+45 26 37 27 62",
			URL:     "https://tsrv.pw",
			Summary: "Linux / UNIX Enthusiast. IT octopus. Runner. I'm a passionate and solution oriented DevOps'er / Sysadm based in Koege, Denmark. I'm optimistic and solution driven. I want to solve problems, and deliver the most value to the users. I'm building positive team environments and greatly enjoy working with others. I enjoy continuous learning and exploring new technologies. When I'm not doing IT stuff at home or work, I enjoy time with my family and running.",
			Location: Location{
				Address:     "Steinmannsvej 10",
				PostalCode:  "4600",
				City:        "Koege",
				CountryCode: "DK",
				Region:      "Sjaelland",
			},
			Profiles: []Profiles{
				{
					Network:  "Linkedin",
					Username: "Teis Angel Clausen",
					URL:      "https://www.linkedin.com/in/teisangelclausen/",
				}, {
					Network:  "Github",
					Username: "AngelFreak",
					URL:      "https://github.com/AngelFreak",
				}, {
					Network:  "Twitter",
					Username: "Teis_Angel",
					URL:      "https://twitter.com/Teis_Angel",
				},
			},
		},
		Work: []Work{
			{
				Name:       "Danish army of defence",
				Position:   "IT-Specialist",
				URL:        "https://fmn.dk//",
				StartDate:  "2020-01-04",
				EndDate:    "N/A",
				Summary:    "",
				Highlights: []string{""},
			},
		},
		Volunteer: []Volunteer{
			{
				Organization: "Danish home guard",
				Position:     "Sergeant",
				URL:          "https://www.hjv.dk",
				StartDate:    "2015-11-01",
				EndDate:      "N/A",
				Summary:      "Currently i command a squad of infantry soldiers. In one of the most active parts of The Danish Home Guard.",
				Highlights:   []string{"Domestic Responce Force, Arctic Respoce Force"},
			},
		},
		Education:    []Education{},
		Awards:       []Awards{},
		Publications: []Publications{},
		Skills:       []Skills{},
		Languages:    []Languages{},
		Interests:    []Interests{},
		References:   []References{},
		Projects:     []Projects{},
	}
}

func respondWithJSON(w http.ResponseWriter, r *http.Request) {
	log.Printf("Handling request for %s from %s", r.URL.Path, r.RemoteAddr)
	resume := getResume()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resume)
}

func (app *App) initialiseRoutes() {
	app.Router = mux.NewRouter()
	app.Router.HandleFunc("/", respondWithJSON)
}

func (app *App) run() {
	port := loadEnvVar("WEB_PORT")
	log.Printf("Starting http server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
}

func loadEnvVar(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {

	app := App{}
	app.initialiseRoutes()
	app.run()
}