#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

int         main(int ac, char **av)
{
    int     numb;
    int     i;
    int     j;
    int     arr[100];


    i = 0;
    j = 0;
    numb = 0;
    while (i < 100)
    {
        arr[i] = 0;
        i++;
    }
    i = 2;
    arr[0] = 2;
    if (ac == 2)
    {
        numb = atoi(av[1]);
        if (numb == 1)
        {
            printf("1\n");
            return (0);
        }
        while(numb != 1)
        {
            while (numb % i == 0 && numb != 1)
            {
                numb /= i;
                if (numb == 1)
                    printf("%d", i);
                else
                    printf("%d*", i);
            }
            i++;
            j = 0;
            while (arr[j] > 0 && i < (numb / 2))
            {
                if (i % arr[j] == 0)
                {
                    i++;
                    j = 0;
                }
                else
                    j++;
            }
            if (arr[j] == 0)
                arr[j] = i;
        }
    }
    printf("\n");
    return (0);
}
