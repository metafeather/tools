package task

import (
	"fmt"
	"os"

	run "github.com/go-cmd/cmd"
)

func Run(args ...string) error {
	cmdOptions := run.Options{
		Buffered:  false,
		Streaming: true,
	}

	// Run direnv before task to ensure env vars are set
	args = append([]string{"exec", ".", "task", "--exit-code", "--silent"}, args...)
	c := run.NewCmdOptions(cmdOptions, "direnv", args...)

	// Stream output
	// ref: https://github.com/go-cmd/cmd/blob/master/examples/blocking-streaming/main.go
	done := make(chan struct{})
	go func() {
		defer close(done)
		for c.Stdout != nil || c.Stderr != nil {
			select {
			case line, open := <-c.Stdout:
				if !open {
					c.Stdout = nil
					continue
				}
				fmt.Println(line)
			case line, open := <-c.Stderr:
				if !open {
					c.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	// Run and wait for Cmd to return, discard Status
	<-c.Start()
	// Wait for goroutine to print everything
	<-done

	return nil
}
