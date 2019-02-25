package temp

import (
	"batman-server/lib"
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

var globalMathRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateCommand() Command {
	pref := PerfInfo{
		TS:     time.Now().UnixNano() / int64(time.Millisecond), // 毫秒
		CPU:    globalMathRand.Float32() * float32(100),
		Memory: globalMathRand.Intn(4096),
		NetIn:  globalMathRand.Intn(2048),
		NetOut: globalMathRand.Intn(2048),
	}

	var d []interface{}
	d = append(d, pref)
	return Command{Action: "push-perf", Content: map[string]interface{}{"data": d}}
}

func Service() {
	c, err := lib.ConnectWS("127.0.0.1:9009", "/ws/pc")
	if err != nil {
		log.Println("connect err:", err)
		return
	}
	for {
		time.Sleep(time.Duration(globalMathRand.Int63n(200)) * time.Millisecond)

		cmd := generateCommand()
		b, err1 := json.Marshal(&cmd)
		if err1 != nil {
			log.Println(err1)
			return
		}

		err = c.WriteMessage(1, b)
		if err != nil {
			log.Println(err)
			return
		}

		_, b, err = c.ReadMessage()
		if err != nil {
			log.Println("read err:", err)
			return
		}

		log.Println(string(b))
	}
}
