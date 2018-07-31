package main

import (
	"flag"
	"os"
	"path/filepath"
	"time"

	"github.com/golang/glog"
	"k8s.io/apimachinery/pkg/util/wait"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	// Required to authenticate against GKE clusters
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/openshift-talks/k8s-go/controller/cmd/policy-controller/controller"
	policyclientset "github.com/openshift-talks/k8s-go/controller/pkg/generated/clientset/versioned"
	policyinformers "github.com/openshift-talks/k8s-go/controller/pkg/generated/informers/externalversions"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	policyClient, err := policyclientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building clientset: %s", err.Error())
	}
	policyInformerFactory := policyinformers.NewSharedInformerFactory(policyClient, time.Second*30)

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building clientset: %s", err.Error())
	}
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)

	controller := controller.NewController(policyClient, policyInformerFactory, kubeClient, kubeInformerFactory)

	go policyInformerFactory.Start(wait.NeverStop)
	go kubeInformerFactory.Start(wait.NeverStop)

	if err = controller.Run(2, wait.NeverStop); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	defaultKubeConfig := ""
	if home := homeDir(); home != "" {
		defaultKubeConfig = filepath.Join(home, ".kube", "config")
	}
	flag.StringVar(&kubeconfig, "kubeconfig", defaultKubeConfig, "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func homeDir() string {
	if h := os.Getenv("HOME"); h != "" {
		return h
	}
	return os.Getenv("USERPROFILE") // windows
}
