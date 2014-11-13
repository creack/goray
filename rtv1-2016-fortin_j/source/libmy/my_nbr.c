/*
** my_nbr.c for libmy in /home/fortin_j/libmy/
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Jan  8 15:44:32 2012 julien fortin
** Last update Mon Jan 16 12:24:40 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>

int     my_put_nbr(int nb)
{
  int	pow;

  if (nb == -2147483648)
    my_putstr("-2147483648");
  else
    {
      pow = 1;
      if (nb < 0)
        {
          my_putchar('-');
          nb = nb * (-1);
        }
      while ((nb / pow) > 9)
	pow = pow * 10;
      while (pow != 0)
        {
          my_putchar('0' + (nb / pow) % 10);
          pow = pow / 10;
        }
    }
  return (0);
}

int     my_getnbr(char *str)
{
  int   i;
  int   nb;
  int   signe;

  i = 0;
  nb = 0;
  signe = 1;
  while (str[i] != '\0' && (str[i] == '+' || str[i] == '-'))
    {
      if (str[i] == '-')
	signe = signe * -1;
      i = i + 1;
    }
  str = str + i;
  i = 0;
  while (str[i] >= '0' && str[i] <= '9')
    {
      nb = nb * 10;
      nb = nb - (str[i] - '0');
      i = i + 1;
    }
  return (nb * signe * -1);
}

int	my_intlen(int nb)
{
  int	i;

  i = 0;
  while (nb != 0)
    {
      nb = nb / 10;
      i = i + 1;
    }
  return (i);
}

void    my_put_nbrbase(unsigned int nb, char *base)
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
  my_putstr(number);
  free(number);
}

void	my_put_nbrbase_long(long nb, char *base)
{
  int   i;
  int   modulo;
  char  *number;
  long	save;

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
  my_putstr(number);
  free(number);
}
