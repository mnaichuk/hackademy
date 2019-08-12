#include <stdlib.h>
 
int *ft_range(int start, int end)
{
    int size = 0;
    if (start == end)
    {
        size = 1;
    }
    else
    {
        int count = start;
        while(count <= end)
        {
            size++;
            count++;
        }
    }
     int *arr;
    
    arr = malloc(sizeof(int)*size);

    while (start <= end && arr != NULL)
    {
        *arr = start;
        start++;
        arr++;
    }
    return arr;
}

int main()
{
    ft_range(1, 3);
}
