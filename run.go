package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func main() {
	// List of services to start
	apis := []string{"auths", "user", "feed", "publish"}
	var wg sync.WaitGroup
	os.Chdir("./api")
	for _, api := range apis {
		wg.Add(1)
		go func(api string) {
			defer wg.Done()
			cmd := exec.Command("go", "run", api+".go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				fmt.Printf("Error starting %s: %v\n", api, err)
				return
			}
			fmt.Printf("Started %s\n", api)
			//err = cmd.Wait()
			//if err != nil {
			//	fmt.Printf("Error waiting for %s: %v\n", api, err)
			//}
		}(api)
	}
	wg.Wait()
	fmt.Println("test")
}
