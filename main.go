package main

import (
	"ffmpegCompress/execit"
	"fmt"
)

func main() {
	out, err := execit.RunCmd(4, []string{
		"Start-Sleep -Seconds 1; Write-Output 'Finished 1'",
		"Start-Sleep -Seconds 2; Write-Output 'Finished 2'",
		"Start-Sleep -Seconds 3; Write-Output 'Finished 3'",
		"Start-Sleep -Seconds 4; Write-Output 'Finished 4'",
	})
	if err != nil {
		fmt.Println(err)
	}
	for _, op := range out {
		fmt.Println(string(op.Stdout))
		fmt.Println(string(op.Stderr))
	}
}
