#include <stdarg.h>
#include <unistd.h>

void check_str(char *str) 
{
  int i = 0, br1 = 0, br2 = 0, br3 = 0, res;
  while (str[i] != '\0') 
  {
    br1 = br1 + (str[i] == '[');
    br1 = br1 - (str[i] == ']');
    br2 = br2 + (str[i] == '(');
    br2 = br2 - (str[i] == ')');
    br3 = br3 + (str[i] == '{');
    br3 = br3 - (str[i] == '}');
    if (br1 < 0) 
    {
      write(1,"Error\n",6);
      return;
    }
    if (br2 < 0) 
    {
      write(1,"Error\n",6); 
      return;    
    }
    if (br3 < 0) 
    {
      write(1,"Error\n",6); 
      return;
    }
    i++;
  }  
  if ((br1 == 0)*(br2 == 0)*(br3 == 0)) 
  {
    write(1,"OK\n",3);
  } else 
  {
    write(1,"Error\n",6);
  }
}
/*
void check_brackets(va_list *ptr)  
{	
  char *str;
  int res;
  while (1) 
  {
    if (str == NULL) 
    {
      break;
    }   
    str = va_arg(ptr,char*);
    check_str(str);   
  }
}
*/
int main(char **args)  
{
  //va_list(args);
  //va_start(args,char*);
  check_str("olololo");	
  check_str("()()([{}]){{([])[]}}");	 
  check_str("]");	
  check_str("()[()");	
  check_str("}{");	
  //check_brackets(args);
}
