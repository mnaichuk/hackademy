#include <unistd.h>

void    ft_putchar(char c)
{
    write(1, &c, 1);
}

int         main(int ac, char **av)
{
    int     i;
    char    *str;


    i = 0;
    if (ac == 2)
    {
        str = av[1];
        while (str[i] == ' ' || str[i] == '\t')
            i++;
        while (str[i])
        {
            while (str[i] && str[i] != ' ' && str[i] != '\t')
            {
                ft_putchar(str[i]);
                i++;
            }
            if (str[i] == ' ' || str[i] == '\t')
            {
                while (str[i] == ' ' || str[i] == '\t')
                    i++;
                if (str[i])
                    write(1, "   ", 3);
            }
        }
    }
    write (1, "\n", 1);
}
