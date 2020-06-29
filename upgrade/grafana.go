package main

import (
	"fmt"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	kubeclient := clientset.AppsV1().Deployments("monitoring")

	// Create resource object
	object := &appsv1.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Deployment",
			APIVersion: "apps/v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      "grafana",
			Namespace: "monitoring",
			Labels: map[string]string{
				"app": "grafana",
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: ptrint32(1),
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{
					"app": "grafana",
				},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{
						"app": "grafana",
					},
				},
				Spec: corev1.PodSpec{
					Volumes: []corev1.Volume{
						corev1.Volume{
							Name: "grafana-configmap-datasources",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "besu-grafana-configmap-datasources",
									},
								},
							},
						},
						corev1.Volume{
							Name: "grafana-configmap-dashboards-dashboard",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "besu-grafana-configmap-dashboards-dashboard",
									},
								},
							},
						},
						corev1.Volume{
							Name: "grafana-configmap-dashboards-besu",
							VolumeSource: corev1.VolumeSource{
								ConfigMap: &corev1.ConfigMapVolumeSource{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "besu-grafana-configmap-dashboards-besu",
									},
								},
							},
						},
					},
					Containers: []corev1.Container{
						corev1.Container{
							Name:  "grafana",
							Image: "grafana/grafana:6.2.5",
							Ports: []corev1.ContainerPort{
								corev1.ContainerPort{
									Name:          "http",
									HostPort:      0,
									ContainerPort: 3000,
									Protocol:      corev1.Protocol("TCP"),
								},
							},
							Env: []corev1.EnvVar{
								corev1.EnvVar{
									Name: "POD_IP",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "status.podIP",
										},
									},
								},
								corev1.EnvVar{
									Name: "POD_NAME",
									ValueFrom: &corev1.EnvVarSource{
										FieldRef: &corev1.ObjectFieldSelector{
											FieldPath: "metadata.name",
										},
									},
								},
								corev1.EnvVar{
									Name:  "GF_SECURITY_ADMIN_USER",
									Value: "admin",
								},
								corev1.EnvVar{
									Name:  "GF_SECURITY_ADMIN_PASSWORD",
									Value: "password",
								},
							},
							Resources: corev1.ResourceRequirements{
								Limits: corev1.ResourceList{
									"cpu":    *resource.NewQuantity(500, resource.DecimalSI),
									"memory": *resource.NewQuantity(536870912, resource.BinarySI),
								},
								Requests: corev1.ResourceList{
									"cpu":    *resource.NewQuantity(100, resource.DecimalSI),
									"memory": *resource.NewQuantity(268435456, resource.BinarySI),
								},
							},
							VolumeMounts: []corev1.VolumeMount{
								corev1.VolumeMount{
									Name:      "grafana-configmap-datasources",
									ReadOnly:  true,
									MountPath: "/etc/grafana/provisioning/datasources/prometheus.yml",
									SubPath:   "prometheus.yml",
								},
								corev1.VolumeMount{
									Name:      "grafana-configmap-dashboards-dashboard",
									ReadOnly:  true,
									MountPath: "/etc/grafana/provisioning/dashboards/dashboard.yml",
									SubPath:   "dashboard.yml",
								},
								corev1.VolumeMount{
									Name:      "grafana-configmap-dashboards-besu",
									ReadOnly:  true,
									MountPath: "/etc/grafana/provisioning/dashboards/besu.json",
									SubPath:   "besu.json",
								},
							},
							ImagePullPolicy: corev1.PullPolicy("IfNotPresent"),
						},
					},
				},
			},
			Strategy:        appsv1.DeploymentStrategy{},
			MinReadySeconds: 0,
		},
	}

	// Manage resource
	_, err = kubeclient.Create(object)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deployment Created successfully!")
}

func ptrint32(p int32) *int32 {
	return &p
}
