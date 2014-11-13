/*
** my_convert_base.c for libmy in /home/fortin_j/libmy
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Mon Jan 16 12:22:07 2012 julien fortin
** Last update Mon Jan 16 12:24:22 2012 julien fortin
*/

#include	<libmy.h>

char	*my_convert_base(unsigned int nb, char *base)
{
  int   i;
  int   modulo;
  char  *number;
  unsigned int	save;

  i = 0;
  save = nb;
  while (nb != 0)
    {
      modulo = nb % my_strlen(base);
      nb = nb / my_strlen(base);
      i = i + 1;
    }
  nb = save;
  number = xmalloc(sizeof(*number) * i + 1);
  number[i] = '\0';
  while (nb != 0)
    {
      modulo = nb % my_strlen(base);
      nb = nb / my_strlen(base);
      number[i - 1] = base[modulo];
      i = i - 1;
    }
  return (number);
}
