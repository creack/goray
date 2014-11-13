/*
** my_printf_option.c for libmy in /home/fortin_j/libmy
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Nov 13 19:32:43 2011 julien fortin
** Last update Tue Jan 10 02:10:38 2012 julien fortin
*/

#include	<stdlib.h>
#include	<stdarg.h>
#include        <libmy.h>

void	my_printf_option_c(va_list ap)
{
  char	c;

  c = va_arg(ap, int);
  my_putchar(c);
}

void	my_printf_option_s(va_list ap)
{
  char	*str;

  str = va_arg(ap, char *);
  if (str != NULL)
    my_putstr(str);
  else
    my_putstr("(null)");
}

void	my_printf_option_i(va_list ap)
{
  int	nb;

  nb = va_arg(ap, int);
  my_put_nbr(nb);
}

void	my_printf_option_ss(va_list ap)
{
  int	i;
  char	*str;

  i = 0;
  str = va_arg(ap, char *);
  if (str != NULL)
    {
      while (str[i])
	{
	  if (str[i] <= 32 || str[i] >= 127)
	    {
	      if (str[i] < 8)
		my_putstr("\\00");
	      else if (str[i] > 99)
		my_putstr("\\");
	      else
		my_putstr("\\0");
	      my_put_nbrbase(str[i], "01234567");
	    }
	  else
	    my_putchar(str[i]);
	  i = i + 1;
	}
    }
}
