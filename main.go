package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

var (
	host     string
	port     string
	username string
	password string
)

func validateFlags() error {
	if host == "" {
		return errors.New("host cannot be blank")
	}

	if port == "" {
		return errors.New("port cannot be blank")
	}

	if username == "" {
		return errors.New("username cannot be blank")
	}

	if password == "" {
		return errors.New("password cannot be blank")
	}

	return nil
}

type Queue struct {
	Name  string `json:"name"`
	Vhost string `json:"vhost"`
}

func init() {
	flag.StringVar(&host, "host", "", "host passed to rabbitmqadmin")
	flag.StringVar(&port, "port", "", "port passed to rabbitmqadmin")
	flag.StringVar(&username, "username", "", "username passed to rabbitmqadmin")
	flag.StringVar(&password, "password", "", "password passed to rabbitmqadmin")
	flag.Parse()

	err := validateFlags()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GetQueues() []Queue {
	cmd := exec.Command(
		"rabbitmqadmin",
		fmt.Sprintf("--host=%s", host),
		fmt.Sprintf("--port=%s", port),
		fmt.Sprintf("--username=%s", username),
		fmt.Sprintf("--password=%s", password),
		"--format=raw_json",
		"list",
		"queues",
	)

	file, _ := ioutil.TempFile("", "amqp-purge")
	defer file.Close()
	cmd.Stdout = file
	cmd.Stderr = os.Stderr

	cmd.Run()

	fmt.Println(file.Name())

	data, _ := ioutil.ReadFile(file.Name())

	var queues []Queue
	json.Unmarshal(data, &queues)

	return queues
}

func PurgeQueue(queue Queue) {
	fmt.Println(queue)

	cmd := exec.Command(
		"rabbitmqadmin",
		fmt.Sprintf("--host=%s", host),
		fmt.Sprintf("--port=%s", port),
		fmt.Sprintf("--username=%s", username),
		fmt.Sprintf("--password=%s", password),
		fmt.Sprintf("--vhost=%s", queue.Vhost),
		"purge",
		"queue",
		fmt.Sprintf("name=%s", queue.Name),
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.Run()
}

func main() {
	queues := GetQueues()

	for _, queue := range queues {
		PurgeQueue(queue)
	}
}
