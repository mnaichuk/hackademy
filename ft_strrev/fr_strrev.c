#include<stdlib.h>
#include<stdio.h>
char *ft_strrev(char *str)
{
  char temp;
  int k=0;
  int i=0;
  int length=0;
  while(str[length])
    length++;
  i=length-1;
  while(k<i)
  {
    temp=str[k];
    str[k]=str[i];
    str[i]=temp;
    k++;
    i--;
  }
  return str;
}

int main()
{
  char str[] = "12123123123";
  ft_strrev(str);
}