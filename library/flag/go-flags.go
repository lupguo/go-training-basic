package main

import (
	"encoding/json"
	"fmt"
	"github.com/jessevdk/go-flags"
)

var opts struct {
	// Slice of bool will append 'true' each time the option
	// is encountered (can be set multiple times, like -vvv)
	Verbose []bool `short:"v" long:"verbose" description:"Show verbose debug information"`

	// Example of automatic marshalling to desired type (uint)
	Offset uint `long:"offset" description:"Offset"`

	// Example of a callback, called each time the option is found.
	//Call func(string) `short:"c" description:"Call phone number"`

	// Example of a required flag
	Name string `short:"n" long:"name" description:"A name" required:"true"`

	// Example of a flag restricted to a pre-defined set of strings
	Animal string `long:"animal" choice:"cat" choice:"dog" description:"Choose your pet"`

	// Example of a value name
	File string `short:"f" long:"file" description:"A file" value-name:"FILE"`

	// Example of a pointer
	Ptr *int `short:"p" description:"A pointer to an integer"`

	// Example of a slice of strings
	StringSlice []string `short:"s" description:"A slice of strings"`

	// Example of a slice of pointers
	PtrSlice []*string `long:"ptrslice" description:"A slice of pointers to string"`

	// Example of a map
	IntMap map[string]int `long:"intmap" description:"A map from string to int"`

	// New post card
	Card string `short:"d" long:"card" description:"Post card"`
}

func main() {
	// Callback which will invoke callto:<argument> to call a number.
	// Note that this works just on OS X (and probably only with
	// Skype) but it shows the idea.
	//opts.Call = func(num string) {
	//	cmd := exec.Command("open", "echo "+num)
	//	cmd.Start()
	//	cmd.Process.Release()
	//}

	// Make some fake arguments to parse.
	_ = []string{
		"-vv",
		"-n", "Me",
		"--animal", "dog", // anything other than "cat" or "dog" will raise an error
		"-p", "3",
		"-s", "hello",
		"-s", "world",
		"--ptrslice", "hello",
		"--intmap", "a:1",
		"--intmap", "b:5",
		"arg1",
		"arg2",
		"arg3",
	}

	// Write help message
	//par := flags.NewParser(&opts, flags.Default)
	//par.WriteHelp(os.Stdout)
	//args, err := par.ParseArgs(os.Args)

	// Parse flags from `args'. Note that here we use flags.ParseArgs for
	// the sake of making a working example. Normally, you would simply use
	// flags.Parse(&opts) which uses os.Args
	//args, err := flags.ParseArgs(&opts, os.Args)
	//if err != nil {
	//	log.Fatalln(err)
	//}

	args, err := flags.Parse(&opts)
	if err != nil {
		return
	}

	fmt.Println(args)
	fmt.Printf("%#v\n", opts)
	m, err := json.MarshalIndent(opts, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(m))
	//fmt.Printf("Verbosity: %v\n", opts.Verbose)
	//fmt.Printf("Offset: %d\n", opts.Offset)
	//fmt.Printf("Name: %s\n", opts.Name)
	//fmt.Printf("Animal: %s\n", opts.Animal)
	//fmt.Printf("Ptr: %d\n", *opts.Ptr)
	//fmt.Printf("StringSlice: %v\n", opts.StringSlice)
	//fmt.Printf("PtrSlice: [%v %v]\n", *opts.PtrSlice[0], *opts.PtrSlice[1])
	//fmt.Printf("IntMap: [a:%v b:%v]\n", opts.IntMap["a"], opts.IntMap["b"])
	//fmt.Printf("Remaining args: %s\n", strings.Join(args, " "))
}
