package main

import (
	"errors"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/api"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/client"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/fields"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	redisserver "github.com/docker/go-redis-server"
	rediscli "github.com/garyburd/redigo/redis"
	"github.com/serialx/hashring"
	"log"
)

type MyHandler struct {
	redisserver.DefaultHandler
}

func (h *MyHandler) Get(key string) ([]byte, error) {

	node, err := locateNodeFromKey(key)

	if err != nil {
		log.Fatal(err)
	}

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

func (h *MyHandler) Set(key string, value []byte) error {

	node, err := locateNodeFromKey(key)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("SET : Key : %s Node : %s", node, key)

	conn, err := rediscli.Dial("tcp", node)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Close()

	if _, err = conn.Do("SET", key, value); err != nil {
		log.Print(err.Error())
		return err
	}

	return nil
}

func locateNodeFromKey(key string) (string, error) {

	servers, err := getRedisServers()

	if err != nil {
		return "", err
	}

	ring := hashring.New(servers)
	server, ok := ring.GetNode(key)

	if !ok {
		return "", errors.New("Cannot locate node.")
	}

	return server, nil
}

func getRedisServers() ([]string, error) {

	config, err := client.InClusterConfig()
	nodes := []string{}

	if err != nil {
		return nodes, err
	}

	client, err := client.New(config)
	if err != nil {
		return nodes, err
	}

	labels, err := labels.Parse("name=redis-node")

	if err != nil {
		return nodes, err
	}

	podList, err := client.Pods(api.NamespaceDefault).List(labels, fields.Everything())

	if err != nil {
		return nodes, err
	}

	for _, pod := range podList.Items {

		address := pod.Status.PodIP
		log.Printf(address)
		nodes = append(nodes, address+":6379")
	}

	return nodes, nil
}

func main() {

	srv, err := redisserver.NewServer(redisserver.DefaultConfig().Port(6379).Handler(&MyHandler{}))
	if err != nil {
		log.Panic(err)
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}
}
