package kw

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var (
	_defaultKubectlPath = "/usr/local/bin/kubectl"
	_defaultKubeConfigPath = os.Getenv("KUBECONFIG")
)

type Kubectl struct {
	BinaryPath string
	KubeConfigPath string
}

func New(binaryPath, kubeConfigPath string) (*Kubectl, error) {
	// kubeConfigPath Arg is NoValue -> use $KUBECONFIG
	bPath := binaryPath
	cPath := kubeConfigPath
	if bPath == EmptyString {
		bPath = _defaultKubectlPath
	}
	log.Printf("info: kubectl binary path is %v\n", bPath)
	if cPath == EmptyString {
		cPath = _defaultKubeConfigPath
	}
	log.Printf("info: kubeconfig path is %v\n", cPath)
	kubectl := &Kubectl{
		BinaryPath:     bPath,
		KubeConfigPath: cPath,
	}
	return kubectl, nil
}

func (k *Kubectl) Apply(yaml string) (string, bool) {
	fileUUID, _ := uuid.NewUUID()
	log.Printf("info: generate uuid %v\n", fileUUID.String())

	tmpFileName := fileUUID.String()
	log.Printf("info: file name is %v\n", tmpFileName)
	tmpFile, _ := ioutil.TempFile("", tmpFileName)
	defer os.Remove(tmpFile.Name())

	_, err := tmpFile.Write([]byte(yaml))
	if err != nil {
		return err.Error(), false
	}
	// CMD Template: kubectl --kubeconfig ${KubeConfigPath} apply -f ${ManifestPath}
	args := []string{"--kubeconfig", k.KubeConfigPath, "apply", "-f", tmpFile.Name()}
	cmd := exec.Command(k.BinaryPath, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf("error: execute kubectl => \n result:%v\n, err:%v\n", stdout.String(), err)
		return stdout.String(), false
	}
	return stdout.String(), true
}

func (k *Kubectl) Delete(yaml string) (string, bool) {
	fileUUID, _ := uuid.NewUUID()
	log.Printf("info: generate uuid %v\n", fileUUID.String())

	tmpFileName := fileUUID.String()
	log.Printf("info: file name is %v\n", tmpFileName)
	tmpFile, _ := ioutil.TempFile("", tmpFileName)
	defer os.Remove(tmpFile.Name())

	_, err := tmpFile.Write([]byte(yaml))
	if err != nil {
		return err.Error(), false
	}
	args := []string{"--kubeconfig", k.KubeConfigPath, "delete", "-f", tmpFile.Name()}
	cmd := exec.Command(k.BinaryPath, args...)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	err = cmd.Run()
	if err != nil {
		fmt.Printf("error: execute kubectl => \n result:%v\n, err:%v\n", stdout.String(), err)
		return stdout.String(), false
	}
	return stdout.String(), true
}