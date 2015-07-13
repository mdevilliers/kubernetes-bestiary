package main

import (
	server "github.com/docker/go-redis-server"
	rediscli "github.com/garyburd/redigo/redis"
	"github.com/serialx/hashring"
	"log"
)

type MyHandler struct {
	values map[string][]byte
}

func (h *MyHandler) GET(key string) ([]byte, error) {

	node := locateNodeFromKey(key)
	log.Printf("GET : Key : %s Node : %s", node, key)

	conn, err := rediscli.Dial("tcp", node)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	v, err := rediscli.Bytes(conn.Do("GET", key))

	if err != nil {
		log.Printf(err.Error())
		return []byte{}, err
	}

	return v, nil
}

func (h *MyHandler) SET(key string, value []byte) error {

	node := locateNodeFromKey(key)
	log.Printf("SET : Key : %s Node : %s", node, key)

	conn, err := rediscli.Dial("tcp", node)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	if _, err = conn.Do("SET", key, value); err != nil {
		log.Printf(err.Error())
		return err
	}

	return nil
}

func locateNodeFromKey(key string) string {

	servers := getRedisServers()

	ring := hashring.New(servers)
	server, _ := ring.GetNode(key)
	return server
}

func getRedisServers() []string {
	// TODO call into kubernetess
	return []string{"192.168.0.246:6379",
		"192.168.0.247:6379",
		"192.168.0.249:6379"}
}

func main() {
	handler, _ := server.NewAutoHandler(&MyHandler{values: make(map[string][]byte)})
	server := &server.Server{Handler: handler, Addr: ":6379"}
	server.ListenAndServe()
}
