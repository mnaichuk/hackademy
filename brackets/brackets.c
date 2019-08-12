#include <unistd.h>

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
    i = 1;
    j = 0;
    if (ac >= 2)
    {
        while (av[i][j])
        {
            j = 0;
            while (av[i][j])
            { 
                if (av[i][j] == '(' || av[i][j] == '[' || av[i][j] == '{')
                {   
                    c = av[i][j];
                    seek = (c == '(') ? (c + 1) : (c + 2);
                    tmp = j;
                }           
                if (av[i][j] == ')' || av[i][j] == ']' || av[i][j] == '}')
                {
                    if (av[i][j] != seek)
                    {
                        ft_putstr("Error\n");
                        return (0);
                    }
                    else
                    {
                        
                        av[i][tmp] = ' ';
                        av[i][j] = ' ';
                        j = -1;
                    }
                }
                j++;
            }
            i++;
        }
    }
    write(1, "\n", 1);
}
