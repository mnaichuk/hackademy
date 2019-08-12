#include<stdio.h>
#include<stdlib.h>
#include<unistd.h>

void rot_13(char *s)
{
  char *tmp;
  tmp=s;
  int i=0;
  while(*s)
  {
    if(*s==' ')
    {
      write(1," ",1);
      s++;
      tmp++;
      continue;
    }
    *tmp=(*s-13);
    write(1,tmp,1);
    tmp++;
    s++;
  }
  write(1,"\n",1);
}
int main()
{
  char arr[]="My horse is Amazing.";
  rot_13(arr);
}