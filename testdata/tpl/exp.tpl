{{.}}
-----
{{"\"output\""}}
	A string constant.
{{`"output"`}}
	A raw string constant.
// 函数格式输出
{{printf "%q" "output"}}
	A function call.
// 管道
{{"output" | printf "%q"}}
	A function call whose final argument comes from the previous
	command.
// 分组
{{printf "%q" (print "out" "," "put")}}
	A parenthesized argument.
// 管道链式操作
{{"put" | printf "%s_%s" "out" | printf "%q" }}
	A more elaborate call.
{{"output" | printf "%s" | printf "%q"}}
	A longer chain.
// with操作
{{with "output"}}{{printf "%q" .}}{{end}}
	A with action using dot.
// with 变量
{{with $x := "output" | printf "%q"}}{{$x}}{{end}}
	A with action that creates and uses a variable.
{{with $x := "output"}}{{printf "%q" $x}}{{end}}
	A with action that uses the variable in another action.
{{with $x := "output"}}{{$x | printf "%q"}}{{end}}
	The same, but pipelined.

----
// 嵌套模板定义
{{template "T3"}}
{{define "T1"}}TPL ONE,{{end}}
{{define "T2"}}TPL TWO{{end}}
{{define "T3"}}{{template "T1"}} {{template "T2"}}{{end}}
----
// 代码块
{{block "name" .}} {{- "THIS IS BLOCK TMPL... "}}{{end}}
