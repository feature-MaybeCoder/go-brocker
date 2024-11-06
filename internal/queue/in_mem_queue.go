package queue

import (
	"encoding/json"
	"fmt"
	"os"

	main_config "github.com/feature-MaybeCoder/go-brocker/internal/config"
	"github.com/feature-MaybeCoder/go-brocker/internal/models"
)

var inMemQueueCapacity int = 10

type inMemQueue struct {
	messages chan models.Message
}
type inMemQueueConfig struct {
	Name string `json:"name"`
}

func (imq *inMemQueue) PushMessage(message models.Message) error {
	imq.messages <- message
	return nil
}

func (imq *inMemQueue) Pop() models.Message {
	message := <-imq.messages
	return message
}

func newInMemQueue(messages chan models.Message) inMemQueue {
	return inMemQueue{
		messages: messages,
	}
}

type inMemQueuesManager struct {
	queues map[string]Queue
}

func (imqm *inMemQueuesManager) CreateQueue(name string) (Queue, error) {
	existing_queue, ok := imqm.queues[name]

	if ok {
		return existing_queue, nil
	}

	messages_chan := make(chan models.Message, inMemQueueCapacity)
	new_queue := newInMemQueue(messages_chan)

	new_queue_config := inMemQueueConfig{
		Name: name,
	}

	new_queue_config_json, err := json.Marshal(new_queue_config)

	if err != nil {
		return nil, err
	}
	new_config_file_path := main_config.MainConfig.QueuesConfigDir.Join(fmt.Sprintf("%s.json", name))
	f, err := os.OpenFile(
		new_config_file_path.String(),
		os.O_RDWR|os.O_CREATE|os.O_TRUNC,
		0755,
	)

	if err != nil {
		return nil, err
	}
	err = f.Truncate(0)
	_, err = f.Seek(0, 0)
	_, err = fmt.Fprintf(f, "%s", new_queue_config_json)

	if err != nil {
		return nil, err
	}

	imqm.queues[name] = &new_queue
	fmt.Println("queue", new_queue_config.Name, "was created")
	return &new_queue, nil
}

func (imqm *inMemQueuesManager) GetQueue(name string) (Queue, bool) {
	existing_queue, ok := imqm.queues[name]

	return existing_queue, ok
}

func (imqm *inMemQueuesManager) RecoverQueues() error {
	queues_config_files, err := os.ReadDir(main_config.MainConfig.QueuesConfigDir.String())
	if err != nil {
		return err
	}
	for _, queue_config_file := range queues_config_files {
		queue_config_path := main_config.MainConfig.QueuesConfigDir.Join(queue_config_file.Name())
		queue_config, err := os.ReadFile(queue_config_path.String())

		if err != nil {
			return err
		}

		var parsed_queue_config inMemQueueConfig
		err = json.Unmarshal(queue_config, &parsed_queue_config)

		if err != nil {
			return err
		}
		imqm.CreateQueue(parsed_queue_config.Name)
	}
	return nil
}

func NewInMemQueuesManager() inMemQueuesManager {
	return inMemQueuesManager{
		make(map[string]Queue),
	}
}
