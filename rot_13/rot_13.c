#include <stdio.h>
#include <stdlib.h>
 
int main(int argc, char *argv[]){
    int i = 0;
    if(argc < 2){
        printf("\n");
    } else {
        char *str = argv[1];
        while(str[i])
        {
            if(str[i] >= 65 && str[i] <= 90){
                for(int j = 0; j < 13; j++){
                    if(str[i] == 65){
                        str[i] = 90;
                    }
                    str[i]--;
                } 
            } else if(str[i] >= 97 && str[i] <= 122){
                for(int j = 0; j < 13; j++){
                    if(str[i] == 97){
                        str[i] = 122;
                    }
                    str[i]--;
                }
            }
            i++;
        }
        printf("%s", str);
    }
    return 0;
}
