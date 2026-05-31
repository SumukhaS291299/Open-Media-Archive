package execit

import (
	"fmt"
	"io"
	"os/exec"
	"sync"
)

type IOBuff struct {
	Stderr []byte
	Stdout []byte
}

// FanIn Process to get combined output
func mergeProcess(procs []chan IOBuff) chan IOBuff {
	out := make(chan IOBuff)
	var wg sync.WaitGroup
	wg.Add(len(procs))

	for _, proc := range procs {
		go func(ps chan IOBuff) {
			defer wg.Done()
			for p := range ps {
				out <- p
			}
		}(proc)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// FanOut Process with n workers
func distributeProcess(nWorkers int, cmdStr chan string) []chan IOBuff {
	outs := make([]chan IOBuff, nWorkers)

	for i := range nWorkers {
		ch := make(chan IOBuff)
		outs[i] = ch

		go func(out chan<- IOBuff) {
			defer close(out)

			for cmd := range cmdStr {
				var wg sync.WaitGroup
				cmdBuilder := exec.Command("pwsh", "-Command", cmd)
				stdout, err := cmdBuilder.StdoutPipe()
				if err != nil {
					return
				}
				stderr, err := cmdBuilder.StderrPipe()
				if err != nil {
					return
				}
				if err := cmdBuilder.Start(); err != nil {
					return
				}
				fmt.Printf("Ran the command %s...\n", cmd)

				wg.Add(2)

				stdoutBytes := make(chan []byte, 1)
				stderrBytes := make(chan []byte, 1)

				go func() {
					stdoutbytes, err := io.ReadAll(stdout)
					if err != nil {
						fmt.Println(err)
					}
					stdoutBytes <- stdoutbytes
					close(stdoutBytes)
					wg.Done()
				}()

				go func() {
					stderrbytes, err := io.ReadAll(stderr)
					if err != nil {
						fmt.Println(err)
					}
					stderrBytes <- stderrbytes
					close(stderrBytes)
					wg.Done()
				}()

				if err := cmdBuilder.Wait(); err != nil {
					fmt.Println("Wait error:", err)
				}
				wg.Wait()

				out <- IOBuff{Stdout: <-stdoutBytes, Stderr: <-stderrBytes}
			}
		}(ch)
	}

	return outs
}
