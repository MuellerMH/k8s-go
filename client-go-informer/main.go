package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/tools/wait"
)

func main() {
	watchpods()
}

func watchpods() {
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
	resyncinterval := time.Second * 30
	informerFactory := informers.NewSharedInformerFactory(clientset, resyncinterval)
	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(new interface{}) { fmt.Printf("Added new pod: %v", new) },
		UpdateFunc: func(old, new interface{}) { fmt.Printf("Updated pod: old: %v, new: %v,", old, new) },
		DeleteFunc: func(obj interface{}) { fmt.Printf("Deleted pod: %v", obj) },
	})
	go informerFactory.Start(wait.NeverStop)
	// pod, err := podInformer.Lister().Pods()
}
