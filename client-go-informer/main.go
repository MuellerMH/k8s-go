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
	resyncinterval := time.Second * 30
	informerFactory := informers.NewSharedInformerFactory(clientset, resyncinterval)
	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc:    func(new interface{}) { fmt.Printf("Added new pod: %v", new) },
		UpdateFunc: func(old, new interface{}) { fmt.Printf("Updated pod: old: %v, new: %v,", old, new) },
		DeleteFunc: func(obj interface{}) { fmt.Printf("Deleted pod: %v", obj) },
	})
	go informerFactory.Start(wait.NeverStop)
	for {
		time.Sleep(time.Second * 10)
		pod, err := podInformer.Lister().Pods("default").Get("webserver")
		if err != nil {
			fmt.Println(fmt.Errorf("Can't get pod updates: %v", err))
			continue
		}
		fmt.Printf("%v", pod)
	}

}
