package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
)

type ServiceManager struct {
	rootPath string
}

func NewServiceManager(rootPath string) *ServiceManager {
	return &ServiceManager{rootPath: rootPath}
}

var BlockService []string
var BlockApi []string

func (sm *ServiceManager) startServices(dirName, ext string) error {
	dirPath := filepath.Join(sm.rootPath, dirName)
	dirs, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, entry := range dirs {
		if !entry.IsDir() {
			continue
		}

		dir := filepath.Join(dirPath, entry.Name())
		if err := os.Chdir(dir); err != nil {
			return err
		}

		fileName := fmt.Sprintf("%s.%s", entry.Name(), ext)
		cmd := exec.Command("go", "run", fileName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Start(); err != nil {
			fmt.Printf("Error running %s %s: %v\n", dirName, entry.Name(), err)
		}

		if err := os.Chdir(sm.rootPath); err != nil {
			return err
		}
	}

	return nil
}

func (sm *ServiceManager) handleSignals() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case sig := <-sigCh:
			fmt.Printf("Received signal: %s\n", sig)
			os.Exit(0)
			return
		}
	}
}

func main() {

	root, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// 获取命令行参数
	var apis, services string
	var serviceSet, apiSet []string
	servicesPath := filepath.Join(root, "services")
	apisPath := filepath.Join(root, "api")
	serviceDir, err := os.ReadDir(servicesPath)
	if err != nil {
		panic(err)
	}
	for _, service := range serviceDir {
		serviceSet = append(serviceSet, service.Name())
	}
	apiDir, err := os.ReadDir(apisPath)
	if err != nil {
		panic(err)
	}
	for _, api := range apiDir {
		apiSet = append(apiSet, api.Name())
	}

	flag.StringVar(&apis, "apis", "", fmt.Sprintf("you can choose apis to run apis：%v", strings.Join(apiSet, " ")))
	flag.StringVar(&services, "services", "", fmt.Sprintf("you can choose services to run services：%v", strings.Join(serviceSet, " ")))
	flag.Parse()
	if apis != "" {
		BlockApi = strings.Split(apis, " ")
	}
	if services != "" {
		BlockService = strings.Split(services, " ")
	}
	log.Println("you will run services:", BlockService)
	log.Println("you will run apis:", BlockApi)

	sm := NewServiceManager(root)
	if err := sm.startServices("services", "go"); err != nil {
		panic(err)
	}
	if err := sm.startServices("api", "go"); err != nil {
		panic(err)
	}

	sm.handleSignals()
}
