// 基于viper读写yaml配置文件
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"io"
	"os"
	"text/tabwriter"
)

func main() {
	// set config file, read and parse
	configFile := `./testdata/config/config.yml`
	viper.SetConfigFile(configFile)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// tab writer print
	tw := tabwriter.NewWriter(os.Stdout, 8, 4, 0, ' ', 0)
	MustFprintln(tw, "company\thomepage")
	MustFprintln(tw, "----\t----")
	for k, v := range viper.GetStringMap("urls") {
		MustFprintf(tw, "%s\t%s\n", k, v)
	}
	_ = tw.Flush()
}

func MustFprintln(w io.Writer, a ...interface{})  {
	if _, err := fmt.Fprintln(w, a...); err != nil {
		panic(err)
	}
}

func MustFprintf(w io.Writer, format string, a ...interface{}) {
	if _, err := fmt.Fprintf(w, format, a...); err != nil {
		panic(err)
	}
}
