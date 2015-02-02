// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Unknwon/com"
	"github.com/Unknwon/macaron"
	"github.com/gorilla/websocket"
)

type person struct {
	name, info string
}

var people = make([]person, 0, 100)

func init() {
	// Load data.
	if com.IsExist("data.txt") {
		data, err := ioutil.ReadFile("data.txt")
		if err != nil {
			log.Printf("Fail to load data file: %v", err)
			os.Exit(2)
		}
		if len(data) == 0 {
			log.Println("Data file cannot be empty")
			os.Exit(2)
		}
		for _, line := range strings.Split(string(data), "\n") {
			line := strings.TrimSpace(line)
			if len(line) == 0 {
				continue
			}

			infos := strings.Split(line, "\t")
			if len(infos) < 2 {
				continue
			}
			people = append(people, person{infos[0], infos[1]})
		}
	} else {
		// Generate fake data.
		s := rand.NewSource(int64(time.Now().Nanosecond()))
		r := rand.New(s)
		for i := 0; i < 100; i++ {
			info := com.ToStr(r.Int())
			people = append(people, person{info, info})
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func sendRandomInfo(ws *websocket.Conn, stop <-chan bool) {
	var err error
	s := rand.NewSource(int64(time.Now().Nanosecond()))
	r := rand.New(s)

	for {
		select {
		case <-stop:
			log.Println("Exit")
			return
		default:
			if err = ws.WriteMessage(websocket.TextMessage, []byte(people[r.Intn(len(people))].info)); err != nil {
				log.Printf("Fail to send info: %v", err)
				return
			}
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func main() {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Get("/", func(ctx *macaron.Context) {
		ctx.HTML(200, "home")
	})
	m.Get("/ws", func(ctx *macaron.Context) {
		conn, err := upgrader.Upgrade(ctx.Resp, ctx.Req.Request, nil)
		if err != nil {
			log.Printf("Fail to upgrade connection: %v", err)
			return
		}

		stop := make(chan bool)
		for {
			_, data, err := conn.ReadMessage()
			if err != nil {
				if err != io.EOF {
					log.Printf("Fail to read message: %v", err)
				}
				return
			}
			msg := string(data)

			switch msg {
			case "START":
				if len(people) == 0 {
					conn.WriteMessage(websocket.TextMessage, []byte("没人了我会乱说？"))
					conn.Close()
					return
				}
				go sendRandomInfo(conn, stop)
			case "STOP":
				stop <- true
			default:
				// Find corresponding name to display.
				for i, p := range people {
					if p.info == msg {
						if err = conn.WriteMessage(websocket.TextMessage, []byte(p.name)); err != nil {
							log.Printf("Fail to send message: %v", err)
							return
						}
						people = append(people[:i], people[i+1:]...)
					}
				}
			}
		}
	})
	m.Run()
}
