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
