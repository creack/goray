/*
** my_free.c for libmy in /home/fortin_j/libmy/
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Jan  8 18:13:28 2012 julien fortin
** Last update Sun Feb  5 15:14:42 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>

void	my_free_wt(char **wt)
{
  int	i;

  i = 0;
  while (wt[i])
    {
      free(wt[i]);
      i = i + 1;
    }
  free(wt[i]);
  free(wt);
}

void	my_free_str(char *s1, char *s2, char *s3, char *s4)
{
  free(s1);
  free(s2);
  free(s3);
  free(s4);
}
