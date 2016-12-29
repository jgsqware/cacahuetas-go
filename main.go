package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"encoding/json"

	"github.com/jgsqware/cacahuetas-go/cacahuetas"
)

const ID = 12
const DATAFOLDER = "/tmp/cacahuetas"

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/generate", handlerGenerate)
	http.HandleFunc("/display/", handlerDisplay)
	http.HandleFunc("/cacahueta/", handlerCacahueta)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.ListenAndServe(":"+port, nil)

}

func createDatafolder() {
	os.MkdirAll(DATAFOLDER, os.ModePerm)
}

func handler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}

func handlerGenerate(w http.ResponseWriter, r *http.Request) {
	users := parseUsersForm(r)
	restrictions := parseRestrictionForm(r)
	cacahuetas.Init(users, restrictions)
	couples := cacahuetas.GenerateCouples()
	createDatafolder()
	generatedURLs := make(map[string]string)
	for _, couple := range couples {
		uniqueID := randomID()
		path := filepath.Join(DATAFOLDER, uniqueID)
		err := ioutil.WriteFile(path, []byte(couple.String()), 0644)
		if err != nil {
			fmt.Printf("Cannot write %v", path)
			break
		}
		generatedURLs[couple.Giver] = uniqueID
	}

	generalUniqueID := randomID()
	generatedURLsJSON, _ := json.Marshal(generatedURLs)
	path := filepath.Join(DATAFOLDER, generalUniqueID)
	err := ioutil.WriteFile(path, []byte(string(generatedURLsJSON)), 0644)
	if err != nil {
		fmt.Fprintf(w, "Something wrong happened...")
		fmt.Printf("Cannot write %v", path)
	}
	http.Redirect(w, r, "/display/"+generalUniqueID, http.StatusFound)
	return
}

func handlerDisplay(w http.ResponseWriter, r *http.Request) {
	uniqueID := r.URL.Path[len("/display/"):]
	path := filepath.Join(DATAFOLDER, uniqueID)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(w, "Something wrong happened...")
			fmt.Printf("Cannot read %v", path)
			return
		}
		var generatedURLs map[string]string
		err = json.Unmarshal(dat, &generatedURLs)
		if err != nil {
			fmt.Fprintf(w, "Something wrong happened...")
			fmt.Printf("Cannot unmarshal %v", path)
		}
		t, _ := template.ParseFiles("templates/display.html")
		data := struct {
			GeneratedURLs map[string]string
			BaseURL       string
		}{
			generatedURLs,
			"https://cacahuetas.herokuapp.com/cacahueta",
		}
		t.Execute(w, data)
	} else {
		fmt.Fprintf(w, "This cacahuetas generation does not exists or was revoked!")
	}

}

func parseUsersForm(req *http.Request) cacahuetas.Users {
	u := strings.Split(strings.Replace(req.FormValue("cacahuetas"), "\r\n", "\n", -1), "\n")
	us := make(cacahuetas.Users)
	for i, user := range u {
		us[user] = i
	}
	return us
}

func parseRestrictionForm(req *http.Request) cacahuetas.Restrictions {
	r := strings.Split(strings.Replace(req.FormValue("restrictions"), "\r\n", "\n", -1), "\n")
	restric := make(cacahuetas.Restrictions)
	for _, restriction := range r {
		s := strings.Split(restriction, ":")
		restric[s[0]] = s[1]
	}
	return restric
}

func handlerCacahueta(w http.ResponseWriter, r *http.Request) {
	uniqueID := r.URL.Path[len("/cacahueta/"):]
	path := filepath.Join(DATAFOLDER, uniqueID)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		dat, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(w, "Something wrong happened...")
			fmt.Printf("Cannot read %v", path)
			return
		}
		c := strings.Split(string(dat), ":")
		couple := cacahuetas.Couple{
			Giver:    c[0],
			Receiver: c[1],
		}
		t, _ := template.ParseFiles("templates/cacahueta.html")
		t.Execute(w, couple)
	} else {
		fmt.Fprintf(w, "This cacahueta does not exists or was revoked!")
	}

}

func randomID() string {
	rand.Seed(time.Now().UTC().UnixNano())
	const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, ID)
	for i := 0; i < ID; i++ {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
