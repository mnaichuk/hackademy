#include <unistd.h>

char *ft_strrev(char *str)
{
    int len = 0;
    char c;
    int i;
    i = 0;
    while(str[i] != '\0')
    {
        len++;
    }
    int j;
    j = 0;
    while (j < len)
    {
        c = str[j];
        str[j] = str[len];
        str[len] = c;
        write(1, &str[j], 1);
        j++;
    }
    return str;
}

int main(int argc, char **argv)
{

    ft_strrev(argv[1]);;
    return 0;
}
