package main

import "fmt"
import "strconv"
import "net/http"
import "github.com/scottkf/steam"
import "encoding/json"
import "reflect"

type Response map[string]interface{}

func (r Response) String() (s string) {
	b, err := json.Marshal(r)
	if err != nil {
		s = ""
		return
	}
	s = string(b)
	return
}

const (
	ApiKey = "<CHANGEME>"
)

var worker = make(chan int)

func main() {

	// for i := 0; i < 10; i++ {
	// 	go func(i int, ci <-chan profileParams) {
	// 		for params := range ci {
	// 		}
	// 	}(i, worker)
	// }

	http.HandleFunc("/ISteamUser/GetFriendList/v1/", func(w http.ResponseWriter, r *http.Request) {
		handleProfile(w, r, "GetFriends")
	})
	http.HandleFunc("/IPlayerService/GetOwnedGames/v0001/", func(w http.ResponseWriter, r *http.Request) {
		handleProfile(w, r, "GetGames")
	})
	http.HandleFunc("/ISteamUser/GetPlayerSummaries/v0002/", func(w http.ResponseWriter, r *http.Request) {
		handleProfile(w, r, "GetPlayerSummary")
	})
	err := http.ListenAndServe(":8080", nil)
	fmt.Println(err) // don't ignore errors
}

// type profileParams struct {
// 	steam_id steam.SteamId
// 	key      string
// 	method   string
// 	w        http.ResponseWriter
// }

func handleProfile(w http.ResponseWriter, r *http.Request, method string) {
	id, err := strconv.Atoi(r.FormValue("steamid"))
	key := r.FormValue("key")

	steam_id := steam.SteamId{Id64: id}

	if err != nil || !steam_id.ValidSteamId() {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintln(w, id, err)
	} else {
		profile := steam.Profile{Id: steam_id}
		reflection := reflect.ValueOf(&profile).MethodByName(method)
		if reflection.IsValid() {
			reflection.Call([]reflect.Value{reflect.ValueOf(key), reflect.ValueOf(steam_id.Id64)})
			fmt.Fprintln(w, Response{"response": profile})
			fmt.Printf("%s for %d\n", method, steam_id.Id64)
		}

		// worker <- profileParams{steam_id: steam_id, key: key, method: method, w: w}
		// <-worker
	}

}
