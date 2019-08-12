#include <unistd.h>

void ft_putstr(char *str)
{
    int i = 0;
    while (str[i])
    {
        if (str[i] != -1)
        {
            write (1, &str[i], 1);
            i++;
        }
    }
    write (1, "\n", 1);
}

int main()
{
    char str[] = "I KNOW ALL YOUR SECRETS!!!";
    ft_putstr(str);
    return 0;
}
