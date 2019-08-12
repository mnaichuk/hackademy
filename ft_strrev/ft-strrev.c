#include <stdio.h>
#include <stdlib.h>

int my_strlen(char *str){
    int i = 0;
    while (str[i] != '\0'){
        i++;
    }

    return i;
}

char *ft_strrev(char *str){
    int len = my_strlen(str);
    char cur;
    for(int i = 0; i < len/2; i++ ){
        cur = str[i];
        str[i] = str[len - i - 1];
        str[len - i - 1] = cur;
    }

    return str;
}


int main(){
    
    char str[] = "asdfghjkl";
    ft_strrev(str);
    printf("%s \n", str);

    return 0;
}
