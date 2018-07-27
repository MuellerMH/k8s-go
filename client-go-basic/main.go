package main

import (
	"flag"
	"fmt"
	"os"
	// the following line is needed to make dep work, just ignore it:
	_ "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	nodes, _ := listnodes()
	fmt.Println(nodes)
}

func listnodes() ([]string, error) {
	var nodes []string
	kubeconfig := flag.String("kubeconfig", os.Getenv("HOME")+"/.kube/config", "path to the kubeconfig file to use")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nodes, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nodes, err
	}
	nodelist, err := clientset.CoreV1().Nodes().List(metav1.ListOptions{})
	if err != nil {
		return nodes, err
	}
	for _, n := range nodelist.Items {
		nodes = append(nodes, n.GetName())
	}
	return nodes, nil
}
