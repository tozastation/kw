package kw

import (
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type Kubectl struct {
	BinaryPath string
	KubeConfigPath string
}

func New(binaryPath, kubeConfigPath string) (*Kubectl, error) {
	kubectl := &Kubectl{
		BinaryPath:     binaryPath,
		KubeConfigPath: kubeConfigPath,
	}
	return kubectl, nil
}

func (k *Kubectl) Apply(yaml string) (string, bool) {
	fileUUID, _ := uuid.NewUUID()
	tmpFileName := fileUUID.String() + ".yaml"
	tmpFile, _ := ioutil.TempFile("", tmpFileName)
	defer os.Remove(tmpFile.Name())
	_, err := tmpFile.Write([]byte(yaml))
	if err != nil {
		return err.Error(), false
	}
	// kubectl --kubeconfig ${KubeConfigPath} apply -f ${ManifestPath}
	args := []string{"--kubeconfig", k.KubeConfigPath, "apply", "-f", tmpFileName}
	result, err := exec.Command(k.BinaryPath, args...).Output()
	if err != nil {
		log.Printf("error: execute kubectl %v", err)
		return  string(result), false
	}
	return string(result), true
}

func (k *Kubectl) Delete(yaml string) (string, bool) {
	fileUUID, _ := uuid.NewUUID()
	tmpFileName := fileUUID.String() + ".yaml"
	tmpFile, _ := ioutil.TempFile("",tmpFileName)
	defer os.Remove(tmpFile.Name())
	_, err := tmpFile.Write([]byte(yaml))
	if err != nil {
		return err.Error(), false
	}
	// kubectl --kubeconfig ${KubeConfigPath} delete -f ${ManifestPath}
	args := []string{"--kubeconfig", k.KubeConfigPath, "delete", "-f", tmpFileName}
	result, err := exec.Command(k.BinaryPath, args...).Output()
	if err != nil {
		log.Printf("error: execute kubectl %v", err)
		return  string(result), false
	}
	return string(result), true
}