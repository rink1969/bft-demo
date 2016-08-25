package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Bft_msg struct {
	M      int
	Gens   []int
	V      int
	Path   []int
	Sender int
}

var total int
var instance_count int
var actorList []int

func getActor(w http.ResponseWriter, r *http.Request) {
	log.Println("getActor handler")
	instance_count = instance_count - 1
	fmt.Fprintf(w, strconv.Itoa(actorList[instance_count]))
}

func start_server() {
	//install web service
	http.HandleFunc("/getActor", getActor)
	http.ListenAndServe("0.0.0.0:7999", nil)
}

func main() {
	count := flag.String("count", "4", "count of nodes")
	flag.Parse()

	flag_total, err := strconv.Atoi(*count)
	if err != nil {
		panic(err)
	}
	total = flag_total
	instance_count = total
	/*
		f, err1 := os.OpenFile("bfttest.log."+time.Now().Format("20060102150405"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err1 != nil {
			panic(err1)
		}
		defer f.Close()
		log.SetOutput(f)
	*/
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("total nodes count", total)

	//prepare actor list traitors less 1/3
	rand.Seed(time.Now().UTC().UnixNano())
	actorList = make([]int, total)
	for i := 0; i < (total-1)/3; i++ {
		actorList[rand.Intn(total)] = 1
	}

	go start_server()

	for {
		if instance_count == 0 {
			break
		}
		time.Sleep(10 * time.Second)
	}

	log.Println("bfttest is started !")

	//select a node to trigger the bft process
	trigger_id := rand.Intn(total)
	log.Println("trigger node", trigger_id)
	trigger_url := "http://localhost:800" + strconv.Itoa(trigger_id) + "/trigger"
	log.Println("trigger url", trigger_url)
	var msg Bft_msg
	msg.M = (total - 1) / 3
	msg.Gens = make([]int, total)
	for j := 0; j < total; j++ {
		msg.Gens[j] = j
	}
	msg.V = 0
	msg.Path = nil
	msg.Sender = trigger_id

	msg_json, err2 := json.MarshalIndent(msg, "", "")
	if err2 != nil {
		panic(err2)
	}
	ret, err3 := http.Post(trigger_url, "", bytes.NewReader(msg_json))
	if err3 != nil {
		panic(err3)
	}
	defer ret.Body.Close()

	for {
		time.Sleep(10 * time.Second)
	}
}
