package execit

// Process execution, proc sync

func RunCmd(maxConcurrentProc int, commands []string) ([]IOBuff, error) {

	inCmds := make(chan string)
	concurrentProcRunner := distributeProcess(maxConcurrentProc, inCmds)
	go func() {
		defer close(inCmds)
		for _, command := range commands {
			inCmds <- command
		}
	}()
	mergedProcRunner := mergeProcess(concurrentProcRunner)
	var combinedOutput []IOBuff
	for combined := range mergedProcRunner {
		combinedOutput = append(combinedOutput, combined)
	}
	return combinedOutput, nil
}
