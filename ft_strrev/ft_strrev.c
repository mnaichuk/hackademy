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

//#include <stdio.h>

int main()
{
    char str[] = "1234567";
    printf("%s\n", ft_strrev(str));
    return (0);
}
