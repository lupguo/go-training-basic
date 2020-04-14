
### Part1 初始化项目
```
// 利用cobra命令初始化一个mini-corba命令
https://github.com/spf13/cobra/blob/master/cobra/README.md

// 进入到目录下
mkdir -p mini-corba && cd mini-corba

// 初始化项目包
cobra init --pkg-name github.com/lupguo/mini-corba

```

### Part2 新增命令

```
// 我们需要相关命令:
cobra add serve     
cobra add config

// 加入create/update/delete 到config命令下，-p 代表 --parent
cobra add create -p 'configCmd'
cobra add update -p 'configCmd'
cobra add delete -p 'configCmd'
```

### Part3 配置控制
```
// 配置文件信息

// 持久FLAG
rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c","", "config file (default is $HOME/.mini-corba.yaml)")
rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

// 局部 flag，在局部定命令中使用
```