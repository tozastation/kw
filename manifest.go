package kw

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/rakyll/statik/fs"
	_ "github.com/tozastation/kw/statik"
	"io/ioutil"
	apps "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	"k8s.io/client-go/kubernetes/scheme"
	"net/http"
)

// Kubernetes Manifest (Kubernetesのマニフェスト)
type Manifest struct {
	DeploymentManifest http.File
}

func NewManifest() (*Manifest, error) {
	manifest := &Manifest{}
	statikFS, err := fs.New()
	if err != nil {
		return nil, errors.New("fs new error")
	}
	manifest.DeploymentManifest, err = statikFS.Open("/deployment.yaml")
	if err != nil {
		return nil, err
	}
	return manifest, nil
}

// NewDeploymentObject is create deployment object
func (m *Manifest) NewDeploymentObject() (*apps.Deployment, error) {
	object, err := m.TransformYamlToObject(KindDeployment)
	if err != nil {
		return nil, err
	}
	return object.(*apps.Deployment), nil
}

// TransformYamlToObject is convert yaml manifest to golang object
func (m *Manifest) TransformYamlToObject(kind string) (interface{}, error) {
	fmt.Printf("info: %s to object\n", kind)
	s := runtime.NewScheme()
	codecFactory := serializer.NewCodecFactory(s)
	deserializer := codecFactory.UniversalDeserializer()
	var resultObject runtime.Object
	switch kind {
	case KindDeployment:
		yaml, err := ioutil.ReadAll(m.DeploymentManifest)
		if err != nil {
			return nil, err
		}
		resultObject, _, err = deserializer.Decode(yaml, nil, &apps.Deployment{})
		if err != nil {
			return nil, err
		}
	default:
		return nil, ErrKindCanNotFound
	}
	return resultObject, nil
}

// TransformObjectToYaml is convert golang object to yaml manifest
func (m *Manifest) TransformObjectToYaml(item runtime.Object) (string, error) {
	s := json.NewSerializerWithOptions(json.DefaultMetaFactory, scheme.Scheme, scheme.Scheme, json.SerializerOptions{Yaml: true})
	var buffer bytes.Buffer
	if err := s.Encode(item, &buffer); err != nil {
		return EmptyString, err
	}
	return buffer.String(), nil
}
