package main

import (
	"flag"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"

	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/K-Phoen/dark/pkg/generated/clientset/versioned"
	informers "github.com/K-Phoen/dark/pkg/generated/informers/externalversions"
	"github.com/K-Phoen/dark/pkg/signals"
)

type config struct {
	MasterURL  string
	KubeConfig string
}

func main() {
	cfg := config{}

	flag.StringVar(&cfg.KubeConfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&cfg.MasterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")

	klog.InitFlags(nil)
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	restCfg, err := clientcmd.BuildConfigFromFlags(cfg.MasterURL, cfg.KubeConfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	darkClient, err := clientset.NewForConfig(restCfg)
	if err != nil {
		klog.Fatalf("Error building dark clientset: %s", err.Error())
	}

	darkInformerFactory := informers.NewSharedInformerFactory(darkClient, time.Second*30)

	controller := NewController(kubeClient, darkClient, darkInformerFactory.Controller().V1().GrafanaDashboards())

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	darkInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		klog.Fatalf("Error running controller: %s", err.Error())
	}
}