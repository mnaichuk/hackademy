#include <unistd.h>

void    ft_putchar(char c)
{
    write (1, &c, 1);
}

void	ft_putstr(char *str)
{
    int i;

    i = 0;
    if (str == 0)
        return(0);
    while (str[i])
    {
        ft_putchar(str[i]);
        i++;
    }
}

int main()
{
    char *str;

    str = "AaBbcC HEY-HOP Lala-Ley";
    ft_putstr(str);
    return (0);
}
