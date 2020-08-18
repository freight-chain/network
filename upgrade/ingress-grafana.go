package main

import (
	"fmt"
	networkingv1beta1 "k8s.io/api/networking/v1beta1"
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
	kubeclient := clientset.ExtensionsV1beta1().Ingresses("monitoring")

	// Create resource object
	object := &networkingv1beta1.Ingress{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Ingress",
			APIVersion: "extensions/v1beta1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "ingress-rules",
			Namespace: "monitoring",
			Labels: map[string]string{
				"app":       "nginx-ingress",
				"namespace": "monitoring",
				"release":   "grafana-ingress",
			},
			Annotations: map[string]string{
				"kubernetes.io/ingress.class":                "nginx",
				"nginx.ingress.kubernetes.io/rewrite-target": "/$1",
				"nginx.ingress.kubernetes.io/ssl-redirect":   "true",
			},
		},
		Spec: networkingv1beta1.IngressSpec{
			Rules: []networkingv1beta1.IngressRule{
				networkingv1beta1.IngressRule{
					IngressRuleValue: networkingv1beta1.IngressRuleValue{
						HTTP: &networkingv1beta1.HTTPIngressRuleValue{
							Paths: []networkingv1beta1.HTTPIngressPath{
								networkingv1beta1.HTTPIngressPath{
									Path: "/(.*)",
									Backend: networkingv1beta1.IngressBackend{
										ServiceName: "besu-grafana",
										ServicePort: intstr.IntOrString{
											Type:   intstr.Type(0),
											IntVal: 3000,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Manage resource
	_, err = kubeclient.Create(object)
	if err != nil {
		panic(err)
	}
	fmt.Println("Ingress Created successfully!")
}
