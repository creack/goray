/*
** my_printf_option_base.c for libmy in /home/fortin_j/libmy
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Tue Nov 15 19:59:28 2011 julien fortin
** Last update Tue Jan 10 02:10:14 2012 julien fortin
*/

#include	<stdarg.h>
#include        <libmy.h>

void	my_printf_option_x(va_list ap)
{
  unsigned int	nb;

  nb = va_arg(ap, unsigned int);
  my_put_nbrbase(nb, "0123456789abcdef");
}

void    my_printf_option_xx(va_list ap)
{
  unsigned int	nb;

  nb = va_arg(ap, unsigned int);
  my_put_nbrbase(nb, "0123456789ABCDEF");
}

void    my_printf_option_o(va_list ap)
{
  unsigned int	nb;

  nb = va_arg(ap, unsigned int);
  my_put_nbrbase(nb, "01234567");
}

void    my_printf_option_u(va_list ap)
{
  unsigned int	nb;

  nb = va_arg(ap, unsigned int);
  my_put_nbrbase(nb, "0123456789");
}

void    my_printf_option_b(va_list ap)
{
  unsigned int	nb;

  nb = va_arg(ap, unsigned int);
  my_put_nbrbase(nb, "01");
}
