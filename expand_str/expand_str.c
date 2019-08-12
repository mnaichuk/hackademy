#include <stdio.h>
#include <stdlib.h>
int my_strlen(char *str){
    int i = 0;
    while(*(str + i) != '\0')
        i++;
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

int main(int argc, char *argv[]){
    char *str = argv[1];
    if(argc != 2 || str[0] == '\0'){
        write(1, "\n", 1);
    } else{
        int i=0;
        char *str = argv[1];
        while(str[i]==32 || str[i] != '\0'){
            str[i] = str[i + 1];
        }
        str = ft_strrev(str);
        while(str[i]==32 || str[i] != '\0'){
            str[i] = str[i + 1];
        }
        str = ft_strrev(str);
        while(str[i] != '\0'){
            while(str[i] != 32 || str[i] != '\0'){
                i++; 
            }
            int j = 0;
            while(str[i] == 32){
                j++;
            }
            if(j<3){
                str[i+3-i] = str[i];
                str[i+1] = 32;
                str[i+2] = 32;
            
            }
            if(j>5){
                i+=2;
                while(str[i] == 32){
                    str[i] = str[i+1];
                }
            }    
        }
    return 0;
}
