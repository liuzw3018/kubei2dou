package k8s

import (
	"flag"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"path/filepath"
)

/**
 * @Author: liuzw
 * @Date: 2023/2/2 10:40
 * @File: client.go
 * @Description: //TODO $
 * @Version:
 */

var (
	ClientSet *kubernetes.Clientset
)

func InitK8SConfig() {
	var kubeConf *string
	if home := homedir.HomeDir(); home != "" {
		kubeConf = flag.String("kubeConfig", filepath.Join(home, ".kube", "config"), "absolute path to the kubeConfig")
	} else {
		kubeConf = flag.String("kubeConfig", "", "absolute path to the kubeConfig")
	}
	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeConf)
	if err != nil {
		klog.Fatalf("读取kubeConfig失败，%s \n", err)
	}

	ClientSet, err = kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("初始化客户端失败, %s \n", err)
	}
}
