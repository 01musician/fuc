package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	// Connect to the etcd cluster.
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"http://localhost:2379"}, // Use the address of your etcd cluster
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println("Failed to connect to etcd:", err)
		return
	}
	defer etcdClient.Close()

	// The key and value to store in etcd.
	key := "exampleKey"
	value := "exampleValue"

	// Put the key-value pair into etcd.
	_, putErr := etcdClient.Put(context.TODO(), key, value)
	if putErr != nil {
		fmt.Println("Failed to put key-value pair into etcd:", putErr)
		return
	}

	fmt.Printf("Key '%s' with value '%s' has been successfully stored in etcd.\n", key, value)
}

