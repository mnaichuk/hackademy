#include <stdlib.h>
#include <stdio.h>

int     *ft_range(int start, int end)
{
    int i;
    int length;
    int sign;
    int *arr;

    i = 0;
    sign = (start < end) ? 1 : -1;
    length = (end - start) * sign + 1;
    arr = (int * )malloc(sizeof(int) * length);
    while (i < length)
    {
        arr[i] = start + i * sign;
        i++;
    }
    return (arr);
}
