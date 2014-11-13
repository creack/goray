/*
** my_printf_option_base_next.c for libmy in /home/fortin_j/libmy/
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Wed Nov 16 16:07:58 2011 julien fortin
** Last update Tue Jan 10 02:10:24 2012 julien fortin
*/

#include        <stdarg.h>
#include        <libmy.h>

void	my_printf_option_p(va_list ap)
{
  long	nb;

  nb = (long) va_arg(ap, void *);
  my_putstr("0x");
  my_put_nbrbase_long(nb, "0123456789abcdef");
}
