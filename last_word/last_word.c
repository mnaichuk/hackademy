#include <unistd.h>

void        ft_putchar(char c)
{
    write (1, &c, 1);
}

int         ft_strlen(char *str)
{
    int     i;

    i = 0;
    while (str[i])
        i++;
    return (i);
}

char        *ft_strrev(char *str)
{
    int     len;
    int     i;
    char    tmp;
    int     j;

    if (!str)
        return (0);
    i = 0;
    len = ft_strlen(str) - 1;
    j = (len + 1) / 2;
    while (i < j)
    {
        tmp = str[i];
        str[i] = str[len - i];
        str[len - i] = tmp;
        i++;
    }
    return (str);
}


int         main(int ac, char **av)
{
    int     i;
    char    *str;

    i = 0;
    str = 0;
    if (ac == 2)
    {
        str = ft_strrev(av[1]);
        while (str[i] && (str[i] == ' ' || str[i] == '\t'))
            i++;
        while (str[i] && str[i] != ' ' && str[i] != '\t')
            i++;
        i--;
        while (i >= 0 && str[i] != ' ' && str[i] != '\t')
        {
            ft_putchar(str[i]);
            i--;
        }
    }
    write(1, "\n", 1);
    return (0);
}
