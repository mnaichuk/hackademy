#include <unistd.h>

void        ft_putchar(char c)
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
        while (str[i])
        {
            if (str[i] >= 65 && str[i] <= 90)
            {
                if (str[i] + 13 > 90)
                    str[i] = 65 + ( 13 - (90 - str[i]) - 1);
                else
                    str[i] += 13;
            }
            else if (str[i] >= 97 && str[i] <= 122)
            { 
                if (str[i] + 13 > 122)
                    str[i] = 97 + (13 - (122 - str[i]) - 1);
                else
                    str[i] += 13;

            }
            ft_putchar(str[i]);
            i++;
        }
    }
    write(1, "\n", 1);
}
