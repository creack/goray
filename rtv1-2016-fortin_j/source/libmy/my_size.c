/*
** my_size.c for libmy in /home/fortin_j//libmy
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Mon Jan 30 12:35:35 2012 julien fortin
** Last update Mon Jan 30 12:40:52 2012 julien fortin
*/

#include	<stdlib.h>

int     my_strlen(char *str)
{
  int   i;

  i = 0;
  if (str != NULL)
    {
      while (str[i])
        i = i + 1;
      return (i);
    }
  return (0);
}

int     my_wtlen(char **wt)
{
  int   i;

  i = 0;
  if (wt != NULL)
    {
      while (wt[i])
        i = i + 1;
      return (i);
    }
  return (0);
}
