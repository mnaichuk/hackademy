#include <unistd.h>

void ft_putstr(char *s) 
{
  while (*s != '\0') 
  {
    write(1,s,1);
    s++;
  }
}

int main()
{
    ft_putstr("12345454564565\n");
    ft_putstr("ahjhjhjhjh");
}
