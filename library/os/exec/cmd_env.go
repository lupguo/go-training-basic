package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"
)

func main() {
	// set command ls environment
	cmd := exec.Command("ls")
	elems := []string{
		"FOO=duplicate_value", // ignored
		"FOO=actual_value",    // this value is used
		"FoO=duplicate_value", // different
	}
	cmd.Env = append(os.Environ(), elems...)
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("cmd env:\n%+v\n\n", cmd.Env)

	// set system env (Case sensitive)
	os.Setenv("foo", "foo")
	os.Setenv("foO", "foO")
	os.Setenv("Foo", "Foo")
	os.Setenv("FOO", "FOO")

	// get system env FOO
	fmt.Printf("Env FOO=%s,Foo=%s\n\n", os.Getenv("FOO"), os.Getenv("Foo"))

	// output
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0 )
	fmt.Fprintln(tw, "Variable\tValue")
	fmt.Fprintln(tw, "----\t----")

	// output env
	limit := 120
	for _, val := range os.Environ() {
		kv := strings.Split(val, "=")
		if len(kv[1]) > limit {
			kv[1] = kv[1][:limit] + "..."
		}
		_, _ = tw.Write([]byte(kv[0]+"\t"+kv[1]+"\n"))
	}
	_ = tw.Flush()
}
