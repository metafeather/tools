package task

import (
	"fmt"
	"os"

	run "github.com/go-cmd/cmd"
)

func Run(args ...string) error {
	// TODO: use pseudoterminal to support color and prompt in go-task
	// ref: https://www.dolthub.com/blog/2022-11-28-go-os-exec-patterns/
	// ref: https://ishuah.com/2021/03/10/build-a-terminal-emulator-in-100-lines-of-go/
	// ref: https://www.kelche.co/blog/go/exec/
	// ref: https://petersouter.xyz/testing-and-mocking-stdin-in-golang/
	// ref: https://github.com/Netflix/go-expect

	// ref: https://blog.kowalczyk.info/article/wOYk/advanced-command-execution-in-go-with-osexec.html
	// path, err := exec.LookPath("ls")

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
