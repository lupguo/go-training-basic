package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"text/tabwriter"
)

func main() {
	cmd := exec.Command("sleep", "1")
	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
	stat := cmd.ProcessState

	kv := map[string]interface{}{
		"SysUsage":           stat.SysUsage(),
		"stat.String()":      stat.String(),
		"stat.Sys(),":        stat.Sys(),
		"stat.ExitCode()":    stat.ExitCode(),
		"stat.Pid(),":        stat.Pid(),
		"stat.Success(),":    stat.Success(),
		"stat.SystemTime(),": stat.SystemTime(),
		"stat.UserTime()":    stat.UserTime(),
	}

	tw := tabwriter.NewWriter(os.Stdout, 0,0,0,' ', 0)
	fmt.Fprintln(tw, "Key\tVal")
	fmt.Fprintln(tw, "---\t---")

	for k, v := range kv {
		fmt.Fprintf(tw, "%s\t%+v\n", k, v)
	}
	fmt.Fprintln(tw, "---\t---")

	tw.Flush()

	fmt.Printf("%+v\n", stat.SysUsage())

	log.Println("Over execute!")
}
