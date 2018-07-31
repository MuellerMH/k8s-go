package main

import (
	"flag"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	// Required to authenticate against GKE clusters
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	"github.com/openshift-talks/k8s-go/crd-go/pkg/generated/clientset/versioned"
)

func main() {
	policies, err := listPolicies()
	if err != nil {
		panic(err)
	}
	fmt.Println(policies)
}

func listPolicies() ([]string, error) {
	var nodes []string
	kubeconfig := flag.String("kubeconfig", os.Getenv("HOME")+"/.kube/config", "path to the kubeconfig file to use")
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nodes, err
	}
	clientset, err := versioned.NewForConfig(config)
	if err != nil {
		return nodes, err
	}
	healthCheckPolicieslist, err := clientset.PolicyV1alpha1().HealthCheckPolicies("default").List(metav1.ListOptions{})
	if err != nil {
		return nodes, err
	}
	for _, n := range healthCheckPolicieslist.Items {
		nodes = append(nodes, n.GetName())
	}
	return nodes, nil
}
