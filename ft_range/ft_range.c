#include <stdio.h>
#include <stdlib.h>

int *ft_range(int start, int end) 
{
  int len = end - start;
  if (len < 0)
  {
    len = - len;	  
    int temp;
    temp = end;
    end = start;
    start = temp;
  }
  int *arr = malloc((len + 1)*sizeof(int));
  int i;
  for (i = 0; i <= len; i++) 
  {
    arr[i] = start + i;
  }
  return arr;
}

int main() {
  int a,b;
  a = 4; 
  b = -2;  
  int *arr = ft_range(a,b);
  int i;
  for (i = 0; i <= abs(b - a);i++ )
  {
    printf("%d\n",arr[i]);
  }
}
