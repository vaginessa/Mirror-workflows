package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//Some variables
const dataDir = "data"

var ids []string

type locationData struct {
	Id       string `'json:"id"`
	Provider string `'json:"provider"`
	Date     uint64 `'json:"date"`
	Lon      string `json:"lon"`
	Lat      string `json:"lat"`
}

func getLocation(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/location/")
	w.Header().Set("Content-Type", "application/json")

	filePath := filepath.Join(dataDir, id)
	files, _ := ioutil.ReadDir(filePath)
	highest := -1
	position := -1
	for i := 0; i < len(files); i++ {
		number, _ := strconv.Atoi(files[i].Name())
		if number > highest {
			highest = number
			position = i
		}
	}
	filePath = filepath.Join(filePath, files[position].Name())
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	w.Write([]byte(fmt.Sprint(string(data))))
}

func putLocation(w http.ResponseWriter, r *http.Request) {
	var location locationData
	err := json.NewDecoder(r.Body).Decode(&location)
	if err != nil {
		fmt.Fprintf(w, "Meeep!, Error")
		return
	}
	path := filepath.Join(dataDir, location.Id)
	os.MkdirAll(path, os.ModePerm)
	files, _ := ioutil.ReadDir(path)
	highest := 0
	for i := 0; i < len(files); i++ {
		number, _ := strconv.Atoi(files[i].Name())
		if number > highest {
			highest = number
		}
	}
	highest += 1
	path = filepath.Join(path, strconv.Itoa(highest))
	file, _ := json.MarshalIndent(location, "", " ")
	_ = ioutil.WriteFile(path, file, 0644)
}

func createDevice(w http.ResponseWriter, r *http.Request) {

}

func generateNewId(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	s := make([]rune, n)
	rand.Seed(time.Now().Unix())
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	newId := string(s)
	for i := 0; i < len(ids); i++ {
		if ids[i] == newId {
			newId = generateNewId(n)
		}
	}
	return newId
}

func handleRequests() {
	http.Handle("/", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/location/", getLocation)
	http.HandleFunc("/newlocation", putLocation)
	http.HandleFunc("/newDevice", createDevice)
	//http.ListenAndServeTLS(":8001", "server.crt", "server.key", nil)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func initData() {
	fmt.Println("Init: Preparing FMD-Server...")
	filePath := filepath.Join(dataDir)
	dirs, _ := ioutil.ReadDir(filePath)
	for i := 0; i < len(dirs); i++ {
		ids = append(ids, dirs[i].Name())
	}
	fmt.Printf("Init: %d Devices registered.\n\n", len(ids))
}

func main() {
	initData()
	fmt.Println("FMD - Server")
	fmt.Println("Starting Server")
	fmt.Println("Port: 8000")
	handleRequests()
}
