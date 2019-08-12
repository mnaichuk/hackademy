#include <stdio.h>

void rot_13(char *str) 
{
  int i = 0;
  char c;
  while (str[i] != '\0') 
  {
    c = str[i];
    if ((c <= 88) && (c >= 65)) 
    {
      c = c - 13;
      if (c < 65) 
      {
        c = c + 24;
      }
    }
    if ((c >= 89) && (c <= 112)) 
    {
      c = c - 13;
      if (c < 89) 
      {
        c = c + 24;
      }
    }
    str[i] = c; 
    i++;  
  } 
}

int main() {
  rot_13("abc");
}
