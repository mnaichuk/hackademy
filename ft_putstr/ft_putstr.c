#include <unistd.h>
void ft_putstr(char *str)
{
  while(*str!='\0')
  {
    write(1,str,1);
    str++;
  }
  write(1,"\n",1);
}
int main()
{
  ft_putstr("hello, my name is");
}