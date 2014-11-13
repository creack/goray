/*
** my_strndup.c for libmy in /home/fortin_j/libmy/
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Mon Oct 24 18:19:04 2011 julien fortin
** Last update Tue Jan 10 02:56:23 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>

void    *xmalloc(int size)
{
  void	*obj;

  if ((obj = malloc(size)) == 0)
    xexit("Fatal error: cannot allocate memory.\n");
  return (obj);
}

char    *my_strdup(char *src)
{
  int   len;
  int   i;
  char  *str;

  i = 0;
  len = my_strlen(src);
  str = malloc(sizeof(*str) * (len + 1));
  if (str == NULL)
    xexit("Fatal error: cannot allocate memory to copy a string.\n");
  while (src[i])
    {
      str[i] = src[i];
      i = i + 1;
    }
  str[i] = '\0';
  return (str);
}

char    *my_strndup(char *src, int i, int nb_char)
{
  char	*str;
  int	x;
  int	y;

  x = 0;
  y = i + nb_char;
  str = malloc(sizeof(*str) * (nb_char + 1));
  if (str == NULL)
    xexit("Fatal error: cannot allocate memory to copy a string.\n");
  while (i < y)
    {
      str[x] = src[i];
      x = x + 1;
      i = i + 1;
    }
  str[x] = '\0';
  return (str);
}

char	*my_strcat(char *str1, char *str2)
{
  char	*str;
  int	len;
  int	i;
  int	x;

  i = 0;
  x = 0;
  len = my_strlen(str1) + my_strlen(str2);
  str = xmalloc(sizeof(*str) * (len + 1));
  while (str1[i])
    {
      str[x] = str1[i];
      x = x + 1;
      i = i + 1;
    }
  i = 0;
  while (str2[i])
    {
      str[x] = str2[i];
      x = x + 1;
      i = i + 1;
    }
  str[x] = '\0';
  return (str);
}

char    *my_strlowcase(char *str)
{
  int   i;

  i = 0;
  while (str[i])
    {
      if (str[i] >= 65 && str[i] <= 90)
        {
          str[i] = str[i] + 32;
        }
      i = i + 1;
    }
  return (str);
}
