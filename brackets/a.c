#include <unistd.h>
#include <stdio.h>

int        ft_strlen(char *str)
{
    int     i;

    i = 0;
    while (str[i])
        i++;
    return (i);
}

void    ft_putchar (char c)
{
    write(1, &c, 1);
}

void    ft_putstr(char *str)
{
    int i;

    i = 0;
    while(str[i])
    {
        ft_putchar(str[i]);
        i++;
    }
}

int         main(int ac, char **av)
{
    int     i;
    int     j;
    int     tmp;
    char    c;
    char    seek;
    
    tmp = 0;
    printf("here\n");
    i = 1;
    j = 0;
    if (ac >= 2)
    {
        printf ("here\n");
        while (av[i][j])
        {
            printf ("here\n");
            j = 0;
            while (av[i][j])
