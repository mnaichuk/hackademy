#include <stdio.h>
#include <stdlib.h>
#include <assert.h>
#include <unistd.h>

int my_strlen(char *str){
    int i = 0;
    while(*(str + i) != '\0')
        i++;
    return i;
}


void ft_puts(char *str){
    int i = 0;

    write(1, str, my_strlen(str));
    write(1, '\n', 1);
}

int main(){
    

    ft_puts("hello world");
    return 0;
}
