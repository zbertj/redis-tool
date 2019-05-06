package main

import (
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"strings"
	"time"

)

var (
	help		bool
	host_port 	string
	pwd			string
	db			int
	action		string
)



func init() {
	flag.BoolVar(&help, "help", false, "this help")
	flag.StringVar(&host_port,"h", "127.0.0.1:6379", "-h [ip:port]")
	flag.StringVar(&pwd,"a", "", "- [password]")
	flag.IntVar(&db,"db", 0, "-db [number]")
	flag.StringVar(&action,"act", "test", "-action [get_key::key|del_key::key]")
}

func getRedisClient(address string, password string, db int) *redis.Client{
	client := redis.NewClient(&redis.Options{
		Addr:         address,
		Password:     password,
		DialTimeout:  1 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
		PoolSize:     5,
		PoolTimeout:  3 * time.Second,
		DB:           db,
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err.Error())

	}else {
		//fmt.Println(address, "|", password, "|", poolsize, "|", db)
		return client
	}

}

func parserAction(action string) []string {
	action_arr := strings.Split(action, "::")
	return action_arr
}

func main() {
	flag.Parse()
	if help {
		flag.Usage()
		return
	}
	if action == "" {
		fmt.Println("need action ")
		return
	}
	//fmt.Println(action)
	client := getRedisClient(host_port, pwd, db)
	arr := parserAction(action)
	switch arr[0] {
	case "test":
		fmt.Println("host_port:", host_port)
		fmt.Println("pwd:", pwd)
		fmt.Println("db:", db)
		fmt.Println("action:", action)

	case "del_key":
		keys := arr[1]
		keys_arr :=client.Keys(keys).Val()
		for _, key := range keys_arr{
			_, err := client.Del(key).Result()
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println("del", key, "done")
			}
		}

	case "set_key":
		key := arr[1]
		value := arr[2]
		_, err := client.Set(key, value, time.Duration(10*60*time.Second)).Result()
		if err != nil{
			fmt.Println(err.Error())
		}else{
			fmt.Println("set", key, "done")
		}

	case "get_key":
		keys := arr[1]
		keys_arr :=client.Keys(keys).Val()
		for _, key := range keys_arr{
			res, err := client.Get(key).Result()
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Println(res)
			}

		}

	default:
		fmt.Println("not support action")

	}

	client.Close()
}