package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"io"
	"path/filepath"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)


func DownloadLogs(clientset *kubernetes.Clientset, pod v1.Pod, containerName string, previous bool) { 

	filename := "./" + pod.Name + ".log"
	if len(pod.Spec.Containers) > 1 {
		filename = "./" + pod.Name + "-" + containerName + ".log"
	}
	fmt.Printf("Downloading logs: %s\n", filename)
	outfile, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer outfile.Close()

	req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &v1.PodLogOptions{Previous: previous, Container: containerName})
	podLogs, err := req.Stream()
	if err != nil {
		fmt.Println("error in opening stream")
		return
	}
	defer podLogs.Close()

	out := bufio.NewWriter(outfile)
	defer out.Flush()

	r := bufio.NewReader(podLogs)
	for {
		bytes, err := r.ReadBytes('\n')
		if _, err := out.Write(bytes); err != nil {
			return
		}

		if err != nil {
			if err != io.EOF {
				return
			}
			return
		}
	}
}

func main() {
	var ns string
	flag.StringVar(&ns, "namespace", "", "namespace")
	flag.Parse()

	kubeconfig := filepath.Join(os.Getenv("HOME"), ".kube", "config")
	fmt.Println("Using kubeconfig file: ", kubeconfig)
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}


	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	pods, err := clientset.CoreV1().Pods(ns).List(metav1.ListOptions{})
	if err != nil {
		log.Fatalln("failed to get pods:", err)
	}


	// print pods
	for i, pod := range pods.Items {
		for _, c := range pod.Spec.Containers {
			DownloadLogs(clientset, pod, c.Name, true)
			//clientset.CoreV1().Pods(ns).GetLogs(pod.GetName(),&v1.PodLogOptions{Previous: true,Container: c.Name})
			fmt.Printf("[%d] %s %s %s\n", i, pod.GetName(), pod.Status.Phase,c.Name)
	}
}
}

// Sample output
//Run as : go run pod.go --namespace <ns-name>
//[0] flink-jobmanager-5897bd55 Running flink-jobmanager
//[1] flink-taskmanager-59c844 Running flink-taskmanager

