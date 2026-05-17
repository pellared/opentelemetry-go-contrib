package main

import (
	"fmt"
	"os/exec"
)

func main() {
	count := 50
	failed := 0
	for i := 0; i < count; i++ {
		cmd := exec.Command("go", "test", "-race", "-v", "-run", "^TestRemotelyControlledSampler$")
		cmd.Dir = "samplers/jaegerremote"
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Iteration %d failed:\n%s\n", i, string(out))
			failed++
			break
		}
		if (i+1)%10 == 0 {
			fmt.Printf("Completed %d iterations...\n", i+1)
		}
	}
	fmt.Printf("Failed %d out of %d iterations\n", failed, count)
}
