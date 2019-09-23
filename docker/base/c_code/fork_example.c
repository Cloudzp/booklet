#include <unistd.h>
#include <stdio.h>
int main(){
    pid_t fpid;
    int count=0;
    fpid=fork();
    if (fpid <0 ){
        printf("fork error ! ");
    }else if (){
        printf("i am child. process id is %d\n",getpid());
    }else{
        printf("i am child. process id is %d\n",getpid());
    }
    return 0;
}

