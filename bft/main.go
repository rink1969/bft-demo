package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//id of this node
var my_id int

//0 is loyal 1 is traitor
var actor int

//total count of nodes
var totoal int

func getActor() int {
	ret, err := http.Get("http://localhost:7999/getActor")
	if err != nil {
		panic(err)
	}
	defer ret.Body.Close()
	body, err := ioutil.ReadAll(ret.Body)
	if err != nil {
		panic(err)
	}
	val, err1 := strconv.Atoi(string(body))
	if err1 != nil {
		panic(err1)
	}
	return val
}

var store map[string]map[int]int
var waitindex int = 0

func path2str(path []int) string {
	var str string
	for _, v := range path {
		str = str + strconv.Itoa(v)
	}
	return str
}

func savamsg(path []int, g int, v int) {
	log.Println("savamsg", path2str(path), g, v)
	array, ok := store[path2str(path)]
	if !ok {
		array = make(map[int]int)
		store[path2str(path)] = array
	}
	array[g] = v
}

func majority(vs []int) int {
	count := len(vs)
	tmp_map := make(map[int]int)
	for _, v := range vs {
		n, ok := tmp_map[v]
		if ok {
			tmp_map[v] = n + 1
		} else {
			tmp_map[v] = 1
		}
	}

	mv := -1
	for k, val := range tmp_map {
		if val > count/2 {
			mv = k
			break
		}
	}
	return mv

}

func waitresult(index int, gens []int, path []int, v int) {
	log.Println("waitresult", index, gens, path, v)
	for {
		time.Sleep(100 * time.Millisecond)
		complete := true

		var vs []int
		for _, gen := range gens {
			if gen == -1 {
				continue
			}
			newpath := append(path, gen)
			array, ok := store[path2str(newpath)]
			if !ok {
				complete = false
				break
			}
			sv, ok1 := array[gen]
			if !ok1 {
				complete = false
				break
			}
			vs = append(vs, sv)
		}
		if complete {
			vs = append(vs, v)
			mv := majority(vs)
			if index == 0 {
				log.Println("bft result is", mv)
			} else {
				savamsg(path, my_id, mv)
			}
			return
		}
	}
}

type Bft_msg struct {
	M      int
	Gens   []int
	V      int
	Path   []int
	Sender int
}

func sendmsg(g int, m int, gens []int, v int, path []int, sender int) {
	path = append(path, my_id)
	trigger_url := "http://localhost:800" + strconv.Itoa(g) + "/trigger"
	log.Println("trigger url", trigger_url)
	var msg Bft_msg
	msg.M = m
	msg.Gens = gens
	msg.V = v
	msg.Path = path
	msg.Sender = sender

	msg_json, err2 := json.MarshalIndent(msg, "", "")
	if err2 != nil {
		panic(err2)
	}
	ret, err3 := http.Post(trigger_url, "", bytes.NewReader(msg_json))
	if err3 != nil {
		panic(err3)
	}
	defer ret.Body.Close()
}

func wrongvlaue(v int) int {
	if v == 1 {
		return 2
	}
	return 1
}

func bft(msg Bft_msg) {
	log.Println("get bft msg", msg)

	//skip myself
	msg.Gens[my_id] = -1

	//commander
	if msg.Path == nil {
		msg.M = msg.M + 1
		for _, g := range msg.Gens {
			if g == -1 {
				continue
			}
			if actor == 0 {
				//loyal genral send 1
				sendmsg(g, msg.M-1, msg.Gens, 1, msg.Path, my_id)
			} else {
				//traitor send rand value(1 or 2)
				sendmsg(g, msg.M-1, msg.Gens, rand.Intn(2)+1, msg.Path, my_id)
			}
		}
		return
	}

	if msg.M > 0 {
		for _, g := range msg.Gens {
			if g == -1 {
				continue
			}
			if actor == 0 {
				//loyal genral
				sendmsg(g, msg.M-1, msg.Gens, msg.V, msg.Path, my_id)
			} else {
				//traitor send wrong value
				sendmsg(g, msg.M-1, msg.Gens, wrongvlaue(msg.V), msg.Path, my_id)
			}
		}
		go waitresult(waitindex, msg.Gens, msg.Path, msg.V)
		waitindex = waitindex + 1
	} else {
		if actor == 0 {
			//loyal genral
			savamsg(msg.Path, msg.Sender, msg.V)
		} else {
			//traitor send wrong value
			savamsg(msg.Path, msg.Sender, wrongvlaue(msg.V))
		}
	}
}

func trigger(w http.ResponseWriter, r *http.Request) {
	log.Println("trigger handler")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var msg Bft_msg
	err = json.Unmarshal(body, &msg)
	if err != nil {
		panic(err)
	}
	bft(msg)
	fmt.Fprintf(w, "OK")
}

func recv(w http.ResponseWriter, r *http.Request) {
	log.Println("recv handler")

	fmt.Fprintf(w, "OK")
}

func main() {
	id := flag.String("id", "0", "id of node")
	num := flag.Int("num", 4, "total number of nodes")
	flag.Parse()

	flag_id, err := strconv.Atoi(*id)
	if err != nil {
		panic(err)
	}
	my_id = flag_id

	totoal = *num

	actor = getActor()
	store = make(map[string]map[int]int)
	/*
		f, err1 := os.OpenFile("bft"+*id+".log."+time.Now().Format("20060102150405"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err1 != nil {
			panic(err1)
		}
		defer f.Close()
		log.SetOutput(f)
	*/
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	log.Println("total number is", *num)
	log.Println("node", *id, "is started !")
	log.Println("my actor is", actor)

	rand.Seed(time.Now().UTC().UnixNano())

	//install web service
	http.HandleFunc("/trigger", trigger)
	http.HandleFunc("/recv", recv)
	http.ListenAndServe("0.0.0.0:800"+*id, nil)
}
