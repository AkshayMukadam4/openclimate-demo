package server

import (
	"log"
	// "encoding/json"
	// "io/ioutil"
	erpc "github.com/Varunram/essentials/rpc"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func frontendFns() {
	getNationStatesId()
	getNationStates()
	getMultiNationalIds()
	getActorIds()
	getEarthStatus()
	getActors()
	postFiles()
	postRegister()
}

func getId(w http.ResponseWriter, r *http.Request) (string, error) {
	var id string
	err := erpc.CheckGet(w, r)
	if err != nil {
		log.Println(err)
		return id, errors.New("request not get")
	}

	urlParams := strings.Split(r.URL.String(), "/")

	if len(urlParams) < 3 {
		return id, errors.New("no id provided, quitting")
	}

	id = urlParams[2]
	return id, nil
}

func getNationStates() {
	http.HandleFunc("/nation-states", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		w.Write([]byte("nation states"))
	})
}

func getNationStatesId() {
	http.HandleFunc("/nation-states/", func(w http.ResponseWriter, r *http.Request) {
		id, err := getId(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		w.Write([]byte("nation state id: " + id))
	})
}

func getMultiNationalIds() {
	http.HandleFunc("/multinationals/", func(w http.ResponseWriter, r *http.Request) {
		id, err := getId(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		w.Write([]byte("multinational id: " + id))
	})
}

func getActorIds() {
	http.HandleFunc("/actors/", func(w http.ResponseWriter, r *http.Request) {
		id, err := getId(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		urlParams := strings.Split(r.URL.String(), "/")
		if len(urlParams) < 4 {
			log.Println("insufficient amount of params")
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}

		choice := urlParams[3]

		switch choice {
		case "dashboard":
			w.Write([]byte("dashboard: " + id))
		case "nation-states":
			w.Write([]byte("nation-states: " + id))
		case "review":
			w.Write([]byte("review: " + id))
		case "manage":
			w.Write([]byte("manage: " + id))
		case "climate-action-asset":
			if len(urlParams) < 5 {
				log.Println("insufficient amount of params")
				erpc.ResponseHandler(w, erpc.StatusBadRequest)
				return
			}
			id2 := urlParams[4]
			w.Write([]byte("climate-action-assets ids: " + id + id2))
		default:
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
			return
		}
	})
}

func getEarthStatus() {
	http.HandleFunc("/earth-status", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		w.Write([]byte("earth status"))
	})
}

func getActors() {
	http.HandleFunc("/actors", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckGet(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		w.Write([]byte("get actors"))
	})
}

func postFiles() {
	http.HandleFunc("/files", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckPost(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		w.Write([]byte("post files"))
	})
}

func postRegister() {
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		err := erpc.CheckPost(w, r)
		if err != nil {
			log.Println(err)
			erpc.ResponseHandler(w, erpc.StatusBadRequest)
		}

		w.Write([]byte("post register"))
	})
}