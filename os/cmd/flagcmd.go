package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

var inta int
var config map[string]interface{}
var open bool
var ip *int
var flagVal myFlag

func init() {
	flag.IntVar(&inta, "inta", 1988, "help message for inta")
	flag.BoolVar(&open, "open", false, "help message for open")

}

func init()  {
	ip = flag.Int("ip", 1234, "help message for ip")
	flag.Var(&flagVal, "id", "help message for my flag id")
	//flag.Parse()
}

type myFlag struct {
	id int
}

func (f *myFlag) String() string {
	return fmt.Sprintf("flagId:%d", f.id)
}

func (f *myFlag) Set(s string) error {
	var err error
	f.id, err = strconv.Atoi(s)
	return err
}

func main() {
	//flag.Parse()
	//
	//fmt.Println(flag.Parsed())
	////flag.Parse()
	//fmt.Println(flag.Parsed())
	//
	////lookup flag
	//mf := flag.Lookup("open")
	//fmt.Println(mf.Name, mf.DefValue, mf.Usage, mf.Value)
	//
	//fmt.Println(inta)
	//fmt.Println(&flagVal)
	//fmt.Println(*ip)
	//fmt.Println(open)

	// flag set
	fss := ""
	fset := flag.NewFlagSet("ExampleValue", flag.ContinueOnError)
	fset.StringVar(&fss, "hv", "default_hv", "help for ffset")
	fset.Usage = func() {
		fmt.Println("Big Big Error: fss", )
	}
	fmt.Println("fset.Parsed()",fset.Parsed())
	err := fset.Parse([]string{"hv"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("fset.Parsed()",fset.Parsed())

	fmt.Println("hv::", fss)

	err = fset.Set("hv", fss)
	if err != nil {
		fmt.Println(err)
	}

	fset.PrintDefaults()
	fset.SetOutput(os.Stdout)
	fsflag := fset.Lookup("hv")
	fmt.Println("fsflag:", fsflag)

	fmt.Println(fset.Name(), fset.Args(), fset.NArg(), fset.NFlag(), fss, )
}
