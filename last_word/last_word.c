#include <stdio.h>
#include <stdlib.h>

int my_strlen(char *str){
    int i = 0;
    while(*(str + i) != '\0')
        i++;
    return i;
}
int main(int argc, char *argv[]){
    int i = 0;
    int flag = 0;
    if(argc < 2){
        write(1, "\n", 1);
    } else {
        char *str = argv[1];
        char *word;
        while(str[i] != '\0'){
            if(str[i] <=32 && str[i+1] > 32){
                flag = 1;
                word = &str[i+1];
            }
            i++;
        }
        write(1, word, my_strlen(word));
    }
    if(flag = 0){
        write(1, "\n", 1);
    }
    return 0;
}
