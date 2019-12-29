package kw

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestKubectl_Apply(t *testing.T) {
	kubectl, err := New("", "")
	if err != nil {
		t.Fatal(err)
	}
	manifest, _ := NewManifest()
	obj, _ := manifest.NewDeploymentObject()
	obj.Name = "test"
	obj.Spec.Selector.MatchLabels = map[string]string{"app": "test"}
	obj.Spec.Template = v1.PodTemplateSpec{
		ObjectMeta: v12.ObjectMeta{
			Labels: map[string]string{"app": "test"},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{Name: "test", Image: "nginx:latest"},
			},
		},
	}
	yaml, _ := manifest.TransformObjectToYaml(obj)
	result, ok := kubectl.Apply(yaml)
	if !ok {
		t.Fatal(result)
	}
	fmt.Print(result)
}

func TestKubectl_Delete(t *testing.T) {
	kubectl, err := New("", "")
	if err != nil {
		t.Fatal(err)
	}
	manifest, _ := NewManifest()
	obj, _ := manifest.NewDeploymentObject()
	obj.Name = "test"
	obj.Spec.Selector.MatchLabels = map[string]string{"app": "test"}
	obj.Spec.Template = v1.PodTemplateSpec{
		ObjectMeta: v12.ObjectMeta{
			Labels: map[string]string{"app": "test"},
		},
		Spec: v1.PodSpec{
			Containers: []v1.Container{
				{Name: "test", Image: "nginx:latest"},
			},
		},
	}
	yaml, _ := manifest.TransformObjectToYaml(obj)
	result, ok := kubectl.Apply(yaml)
	if !ok {
		t.Fatal(result)
	}
	result, ok = kubectl.Delete(yaml)
	if !ok {
		t.Fatal(result)
	}
	fmt.Print(result)
}