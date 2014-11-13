/*
** my_errors.c for libmy in /home/fortin_j/libmy/
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Jan  8 16:05:31 2012 julien fortin
** Last update Sun Jan  8 17:31:09 2012 julien fortin
*/

#include	<stdlib.h>
#include	<unistd.h>

void	my_putchar_error(char c)
{
  write(2, &c, 1);
}

void	my_putstr_error(char *str)
{
  int	i;

  i = 0;
  while (str[i])
    {
      my_putchar_error(str[i]);
      i = i + 1;
    }
}

void	xexit(char *str)
{
  my_putstr_error(str);
  exit(EXIT_FAILURE);
}
