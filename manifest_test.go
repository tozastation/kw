package kw

import (
	"gopkg.in/yaml.v2"
	app "k8s.io/api/apps/v1"
	"testing"
)

func TestManifest_TransformYamlToObject(t *testing.T) {
	manifest, err := NewManifest()
	if err != nil {
		t.Fatalf(err.Error())
	}
	expect := "name"
	actual, err := manifest.TransformYamlToObject(KindDeployment)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if actual.(*app.Deployment).ObjectMeta.Name != expect {
		t.Error("Actual： ", actual.(*app.Deployment).ObjectMeta.Name, "Expect： ", expect)
	}
	t.Log("Success")
}

func TestManifest_TransformObjectToYaml(t *testing.T) {
	manifest, err := NewManifest()
	if err != nil {
		t.Fatalf("on NewManifest:%v", err)
	}
	expect := "name"
	d := &app.Deployment{}
	d.Name = "name"
	data, err := manifest.TransformObjectToYaml(d)
	if err != nil {
		t.Fatalf("on manifest.TransformObjectToYaml:%v", err)
	}
	m := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(data), &m); err != nil {
		t.Fatalf("on  yaml.Unmarshal:%v", err)
	}
	actual := m["metadata"].(map[interface{}]interface{})["name"]
	if actual != expect {
		t.Fatalf("actual：\n%v \n expect:\n%v", m, expect)
	}
	t.Log("Success")
}