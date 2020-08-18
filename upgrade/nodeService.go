package main

import (
	"fmt"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"os"
	"path/filepath"
)

func main() {
	// Create client
	var kubeconfig string
	kubeconfig, ok := os.LookupEnv("KUBECONFIG")
	if !ok {
		kubeconfig = filepath.Join(homedir.HomeDir(), ".kube", "config")
	}

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	kubeclient := clientset.CoreV1().Services("besu")

	// Create resource object
	object := &corev1.Service{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Service",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "besu-node",
			Namespace: "besu",
			Labels: map[string]string{
				"app": "node",
			},
		},
		Spec: corev1.ServiceSpec{
			Ports: []corev1.ServicePort{
				corev1.ServicePort{
					Name:     "discovery",
					Protocol: corev1.Protocol("UDP"),
					Port:     30303,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(0),
						IntVal: 30303,
					},
					NodePort: 0,
				},
				corev1.ServicePort{
					Name:     "rlpx",
					Protocol: corev1.Protocol("TCP"),
					Port:     30303,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(0),
						IntVal: 30303,
					},
					NodePort: 0,
				},
				corev1.ServicePort{
					Name:     "json-rpc",
					Protocol: corev1.Protocol("TCP"),
					Port:     8545,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(0),
						IntVal: 8545,
					},
					NodePort: 0,
				},
				corev1.ServicePort{
					Name:     "ws",
					Protocol: corev1.Protocol("TCP"),
					Port:     8546,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(0),
						IntVal: 8546,
					},
					NodePort: 0,
				},
				corev1.ServicePort{
					Name:     "graphql",
					Protocol: corev1.Protocol("TCP"),
					Port:     8547,
					TargetPort: intstr.IntOrString{
						Type:   intstr.Type(0),
						IntVal: 8547,
					},
					NodePort: 0,
				},
			},
			Selector: map[string]string{
				"app": "node",
			},
			Type:                corev1.ServiceType("ClusterIP"),
			HealthCheckNodePort: 0,
		},
	}

	// Manage resource
	_, err = kubeclient.Create(object)
	if err != nil {
		panic(err)
	}
	fmt.Println("Service Created successfully!")
}
