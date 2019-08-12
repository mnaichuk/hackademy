#include <stdio.h>
#include <stdlib.h> 
 int str_length(const char * str)
 {
     int i = 0;
     while (*str++ != '\0')
     {
         i++;
     }
     printf("%d", i);
     return i;
 }
 


char *expand_str(char *str)
{

    while (*str != '\0')
    {
        printf ("%c", *str);
        if (*str == ' ')
            printf("  ");
        str++;
    }
    printf("\n");
    return str;
}

int main(int argc, char **argv)
{
    char * ssstttrrr = "";
    if (argc != 2 || *argv[1] == *ssstttrrr)
    {
        printf("\n");
        return 1;
    }

    expand_str(argv[1]);
        return 0;
}
