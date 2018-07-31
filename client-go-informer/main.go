package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	watchpods()
}

func watchpods() {
	kubeconfig := flag.String("kubeconfig", os.Getenv("HOME")+"/.kube/config", "path to the kubeconfig file to use")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Println(fmt.Errorf("Can't build config from flags: %v", err))
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println(fmt.Errorf("Can't get config: %v", err))
		os.Exit(1)
	}
	resyncinterval := time.Second * 20
	informerFactory := informers.NewSharedInformerFactory(clientset, resyncinterval)
	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(new interface{}) { fmt.Printf("INFORMER: add pod\n") },
		UpdateFunc: func(old, new interface{}) { fmt.Printf("INFORMER: update pod\n") },
		DeleteFunc: func(obj interface{}) { fmt.Printf("INFORMER: delete pod\n") },
	})
	go informerFactory.Start(wait.NeverStop)
	for {
		time.Sleep(time.Second * 5)
		pod, err := podInformer.Lister().Pods("default").Get("shell")
		if err != nil {
			fmt.Println(fmt.Errorf("Can't get updates on pod `shell`: %v", err))
			continue
		}
		fmt.Printf("Labels of pod `shell`: %v\n", pod.GetLabels())
	}

}
