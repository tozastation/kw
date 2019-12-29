# kw (kubectl wrapper)
kubernetes sdkを介さず，直接kubectlバイナリを利用するラッパーライブラリです．
## TODO
- [x] : (manifest) Create deployment object
- [x] : (manifest) Create deployment yaml to object
- [x] : (kubectl) kubectl delete
- [x] : (kubectl) kubectl apply

## Installation (設置)
To install kw package, you need to install Go and set your Go workspace first.  
### 1. The first need Go installed.
- `go get -u github.com/tozastation/kw`  
### 2. import it in your code
- `import "github.com/tozastation/kw"`

## Quick Start (早速始めよう)
```
package main

import "github.com/tozastation/kw"

func main() {
    // kubectl client
    kubectl, _ := kw.New("", "")    
    // yaml object converter
    manifest, _ := kw.NewManifest()
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
}
```

## Library Test (ライブラリテスト)
1. `make` : テスト用のイメージをビルドします
2. `make test` : ライブラリのプログラムをテストします．