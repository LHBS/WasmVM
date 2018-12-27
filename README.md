## wagon
wagon是go语言实现的WebAssembly解释器，可作为区块链虚拟机使用。<br>

## 用法
示例：
```
#include <string.h>
#include <stdio.h>
#include <module.h>
#include <print.h>
#include <runtime.h>

typedef struct{
    int num;
    char *name;
}averger;

int init(int num, char *name){
    averger m;
    m.num = num;
    m.name = name;
    printi(m.num);
    prints(m.name);
    return 0;
}

export int apply(char *method){
    if(strcmp(method, "init") == 0){
        int num = read_param(1);
        char *name = read_param(2);
        init(num, name);
    }
    return 0;
}
```
源码在cmd/wasm/testdata/test.c，使用[C语言编译器](https://github.com/LHBS/C2Webassembly)编译源码生成wasm文件。<br>
在cmd/wasm目录下执行：<br>
`go build -o vm main.go ` <br>
`./vm`
## 添加API
cmd/wasm/wasmservice/service.go中RegisterApi提供了注册API的方法。
