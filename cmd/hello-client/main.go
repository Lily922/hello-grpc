package main

import (
	//"context"
	"flag"
	"fmt"
	"log"
	"os"

	pf "github.com/LilyFaFa/hello-grpc/pkg/portforwarder"
	pb "github.com/LilyFaFa/hello-grpc/proto"
	//	"google.golang.org/grpc"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	namespace  = "default"
	serverPort = 10000
)

func main() {
	// server address
	configAddr := flag.String("config-addr", "/.kube/config", "remote hello server address")
	// 解析 flag
	flag.Parse()
	// 日志抬头使用 LstdFlags
	// 日志默认输出到 Stderr
	logger := log.New(os.Stderr, "hello:", log.LstdFlags)
	client := NewClient()
	tunnelPort := createTunnel(logger, *configAddr)
	req := &pb.Request{}
	res, err := client.sayHello(tunnelPort, req)
	if err != nil {
		logger.Fatalf("%v", err)
	}
	logger.Printf("remote server say:%s", res.Message)

}

func createTunnel(logger *log.Logger, configAddr string) (tunnelPort string) {
	config, err := NewClusterConfig(configAddr)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	kubeClient, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	podName, err := pf.GetPodName(kubeClient.CoreV1(), namespace)
	if err != nil {
		logger.Fatalf("%v", err)
	}

	tunnel := pf.NewTunnel(kubeClient.Core().RESTClient(), config, namespace, podName, serverPort)
	err = tunnel.ForwardPort()
	if err != nil {
		logger.Fatalf("%v", err)
	}
	podHost := fmt.Sprintf("localhost:%d", tunnel.Local)
	return podHost
}

func NewClusterConfig(kubeConfig string) (*rest.Config, error) {
	var cfg *rest.Config
	var err error

	if kubeConfig != "" {
		cfg, err = clientcmd.BuildConfigFromFlags("", kubeConfig)
	} else {
		cfg, err = rest.InClusterConfig()
	}
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
