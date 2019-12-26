package kw

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"testing"
)

func TestKubectl_Apply(t *testing.T) {
	kubectl, err := New("kubectl", "/root/.kube/config")
	if err != nil {
		t.Fatal(err)
	}
	manifest, _ := NewManifest()
	obj, _ := manifest.NewDeploymentObject()
	obj.Name = "test"
	obj.Spec.Template = v1.PodTemplateSpec{Spec:v1.PodSpec{Containers:[]v1.Container{{Name:"test", Image:"nginx:latest"}}}}
	yaml, _ := manifest.TransformObjectToYaml(obj)
	fmt.Printf("info: yaml => %v", yaml)
	result, ok := kubectl.Apply(yaml)
	if !ok {
		t.Fatal(result)
	}
	fmt.Print(result)
}