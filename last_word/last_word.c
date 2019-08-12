#include <unistd.h>

void last_word( char *str )
{
    char *l_w;
    int i = 0;
    while (str[i] != '\0')
    {
        if (str[i] <= 32 && str[i + 1] > 32)
            l_w = &str[1 + i];
        i++;
    }
    i = 0;
    while (l_w && l_w[i] > 32 )
    {
        write(1, &l_w[i], 1);
        i++;
    }
}


int main(int argc, char **argv)
{

    char str[] = "";
    if(argc != 2 || argv[1] == str )
    {
        write(1, "\n", 1);
        return 1;
    }
    last_word(*argv);
}

