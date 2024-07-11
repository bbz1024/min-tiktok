package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
)

type ServiceManager struct {
	rootPath string
}

func NewServiceManager(rootPath string) *ServiceManager {
	return &ServiceManager{rootPath: rootPath}
}

var BlockService = []string{""}
var BlockApi = []string{""}

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

	sm := NewServiceManager(root)
	if err := sm.startServices("services", "go"); err != nil {
		panic(err)
	}
	if err := sm.startServices("api", "go"); err != nil {
		panic(err)
	}

	sm.handleSignals()
}
