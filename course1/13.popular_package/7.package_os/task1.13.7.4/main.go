package main

import (
	"fmt"
	"os/exec"
)

/*
func ExecBin(binPath string, args ...string) string {

		args = append([]string{binPath}, args...)

		//getting path
		path := os.Getenv("PATH")
		paths := strings.Split(path, string(os.PathListSeparator))
		for _, dir := range paths {
			fullPath := dir + string(os.PathSeparator) + binPath
			_, err := os.Stat(fullPath)
			if err == nil {
				binPath = fullPath
				break
			}
		}
		//execute
		attr := &os.ProcAttr{
			Env: os.Environ(),
			Files: []*os.File{
				os.Stdin,
				os.Stdout,
				os.Stderr,
			}}
		pid, err := os.StartProcess(binPath, args, attr)
		if err != nil {
			return "Error executing binary: " + err.Error()
		}
		_, err = pid.Wait()
		if err != nil {
			return "Error waiting binary: " + err.Error()
		} else {
			return "Hello, World!"
		}
	}
*/
func ExecBin(binPath string, args ...string) string {
	command := exec.Command(binPath, args...)
	output, err := command.CombinedOutput()
	if err != nil {
		return "Error waiting binary: " + err.Error()
	}
	return string(output)
}

func main() {
	fmt.Println(ExecBin("ls", "-la"))
	fmt.Println(ExecBin("nonexistent-binary"))
}
