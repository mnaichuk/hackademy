#include <stdio.h>

int is_prime(int n) 
{
  int i;
  for (i = 2; i*i < n; i++) {
    if ((n % i) == 0) 
    {
      return 0;
    }
  }
  return 1;
}

void fprime(int n) 
{ int i,flag = 0; 
  int temp = n/2;
  for (i = 2; i <= temp; i++)
  {
    if (is_prime(i)) 
    {
      if ( n%i == 0) 
      { 
	if (flag)
       	{
	  printf("*");
	}      
        flag = 1;	      
        n = n / i;
        printf("%d",i);
        i--;	
      } 
    }
  }
  if (!flag) 
  {
    printf("%d*%d",1,n);
  }
  printf("\n");
}

int main() {
  fprime(180);
}
