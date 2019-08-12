#include <unistd.h>
 

char *rot_13 (char *str)
{
    char c;
    int i;

    i = 0;
    while (str[i] != '\0')
    {
        str[i] = str[i] + 13;
        c = str[i];
        write(1, &c, 1);
        i++;
    }
    return str;
}

int main(int argc, char **argv)
{

    if (argc != 2 )
    {
        write(1, "\n", 1);
        return 1;
    }
    rot_13(argv[1]);
      write(1, "\n", 1);
}
