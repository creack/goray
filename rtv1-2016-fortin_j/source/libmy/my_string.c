/*
** my_string.c for libmy in /home/fortin_j/libmy/
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Jan  8 15:36:05 2012 julien fortin
** Last update Tue Jan 31 17:11:03 2012 julien fortin
*/

#include	<stdlib.h>
#include	<unistd.h>

int     my_char_isnum(char c)
{
  if (c >= 48 && c <= 57)
    return (1);
  else
    return (0);
}

void	my_putchar(char c)
{
  write(1, &c, 1);
}

void	my_putstr(char *str)
{
  int	i;

  i = 0;
  if (str != NULL)
    {
      while (str[i])
	{
	  my_putchar(str[i]);
	  i = i + 1;
	}
    }
}

int	my_tputs(int c)
{
  write(1, &c, 1);
  return (42);
}
