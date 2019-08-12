#include <stdio.h>
#include <stdlib.h>

int my_abs(int digit){
    if(digit < 0){
        digit = -digit;
    }

    return digit;
}

int *ft_range(int start, int end) {
    int *res;
    int size = my_abs(start - end);
    res = malloc(size);
    int cur = start;
    if (start < end) {
        cur = start;
        for(int i = 0; i <= size; i++) {
            res[i] = cur;
            cur++;
            printf("%d ", res[i]);
        }
    } else {
        for(int i = 0; i <= size; i++) {
            res[i] = cur;
            cur--;
            printf("%d ", res[i]);
        }
    }
    printf("\n");
    return res;
}

int main(){
    int start = 0, end = -3;
    int size = my_abs(start-end);
    int *res = ft_range(start, end);
    for(int i = 0; i<= size; i++)
        printf("%d ", res[i]);
    return 0;
}
