# bft-demo
A demo of Byzantine Fault Tolerance algorithm

Lead by  http://cs.brown.edu/courses/cs138/s16/lectures/19consen-notes.pdf Page31

It's original algorithm described by paper -- The Byzantine Generals Problem

It's expensive but easy to understand

# build
* add current path to GOPATH
* compile
```
go build bft
go build bfttest
```

# run and example(on windows)
open five consoles

* console 1
```
>bfttest.exe
2016/08/25 16:30:56 main.go:59: total nodes count 4
2016/08/25 16:31:01 main.go:28: getActor handler
2016/08/25 16:31:03 main.go:28: getActor handler
2016/08/25 16:31:05 main.go:28: getActor handler
2016/08/25 16:31:08 main.go:28: getActor handler
2016/08/25 16:31:16 main.go:77: bfttest is started !
2016/08/25 16:31:16 main.go:81: trigger node 3
2016/08/25 16:31:16 main.go:83: trigger url http://localhost:8003/trigger
```
* console 2
```
>bft.exe -id "0"
2016/08/25 16:31:01 main.go:257: total number is 4
2016/08/25 16:31:01 main.go:258: node 0 is started !
2016/08/25 16:31:01 main.go:259: my actor is 0
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {1 [0 1 2 -1] 1 [3] 3}
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8001/trigger
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8002/trigger
2016/08/25 16:31:16 main.go:87: waitresult 0 [-1 1 2 -1] [3] 1
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {0 [0 -1 2 -1] 1 [3 1] 1}
2016/08/25 16:31:16 main.go:54: savamsg 31 1 1
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {0 [0 1 -1 -1] 2 [3 2] 2}
2016/08/25 16:31:16 main.go:54: savamsg 32 2 2
2016/08/25 16:31:17 main.go:114: bft result is 1
```
* console 3
```
>bft.exe -id "1"
2016/08/25 16:31:03 main.go:257: total number is 4
2016/08/25 16:31:03 main.go:258: node 1 is started !
2016/08/25 16:31:03 main.go:259: my actor is 0
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {0 [-1 1 2 -1] 1 [3 0] 0}
2016/08/25 16:31:16 main.go:54: savamsg 30 0 1
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {1 [0 1 2 -1] 1 [3] 3}
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8000/trigger
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8002/trigger
2016/08/25 16:31:16 main.go:87: waitresult 0 [0 -1 2 -1] [3] 1
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {0 [0 1 -1 -1] 2 [3 2] 2}
2016/08/25 16:31:16 main.go:54: savamsg 32 2 2
2016/08/25 16:31:17 main.go:114: bft result is 1
```
* console 4
```
>bft.exe -id "2"
2016/08/25 16:31:05 main.go:257: total number is 4
2016/08/25 16:31:05 main.go:258: node 2 is started !
2016/08/25 16:31:05 main.go:259: my actor is 1
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {0 [-1 1 2 -1] 1 [3 0] 0}
2016/08/25 16:31:16 main.go:54: savamsg 30 0 2
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {0 [0 -1 2 -1] 1 [3 1] 1}
2016/08/25 16:31:16 main.go:54: savamsg 31 1 2
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {1 [0 1 2 -1] 1 [3] 3}
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8000/trigger
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8001/trigger
2016/08/25 16:31:16 main.go:87: waitresult 0 [0 1 -1 -1] [3] 1
2016/08/25 16:31:17 main.go:114: bft result is 2
```
* console 5
```
>bft.exe -id "3"
2016/08/25 16:31:08 main.go:257: total number is 4
2016/08/25 16:31:08 main.go:258: node 3 is started !
2016/08/25 16:31:08 main.go:259: my actor is 0
2016/08/25 16:31:16 main.go:211: trigger handler
2016/08/25 16:31:16 main.go:161: get bft msg {1 [0 1 2 3] 0 [] 3}
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8000/trigger
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8001/trigger
2016/08/25 16:31:16 main.go:134: trigger url http://localhost:8002/trigger
```

#Notes
* possible result is 1 and 2
* actor: 0 is loyal 1 is traitor
* so we can see loyal generals get same bft result
* you can test more nodes with argument -count (for bfttest) -num (for bft)
