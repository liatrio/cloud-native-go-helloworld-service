package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestCloudNativeGoHelloWorldService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "cloud-native-go-helloworld-service")
}
