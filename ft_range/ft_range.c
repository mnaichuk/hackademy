#include <stdio.h>
#include <stdlib.h>
int *ft_range(int start, int end)
{
  int *p;
  int size;
  int flag=0;
  if (start>end)
  {
    flag=1;
  }
  if (flag)
  {
    size=start-end;
  }
  else
  {
    size=end-start;
  }
  if(size<0)
  {
    size=-size;
  }
  size=size+1;
  p=malloc(sizeof(int)*size);
  int i=0;
  if (flag)
  {
    for(;i<size;i++)
    {
      p[i]=start-i;
      printf("%d\n",p[i]);
    }
    return p;
  }
  else
  {
    for(;i<size;i++)
    {
      p[i]=start+i;
      printf("%d\n",p[i]);
    }
    return p;
  }
}

int main()
{
  int *px;
  px=ft_range(10,4);
}