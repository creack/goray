/*
** my_printf.c for my_printf in /home/fortin_j//afs/projets/my_printf
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Tue Nov  8 12:02:38 2011 julien fortin
** Last update Tue Jan 10 02:09:58 2012 julien fortin
*/

#include        <stdarg.h>
#include	<libmy.h>

void	my_printf_set_option(void (**my_printf_option)(va_list ap))
{
  my_printf_option[0] = &my_printf_option_s;
  my_printf_option[1] = &my_printf_option_ss;
  my_printf_option[2] = &my_printf_option_i;
  my_printf_option[3] = &my_printf_option_i;
  my_printf_option[4] = &my_printf_option_c;
  my_printf_option[5] = &my_printf_option_x;
  my_printf_option[6] = &my_printf_option_xx;
  my_printf_option[7] = &my_printf_option_o;
  my_printf_option[8] = &my_printf_option_u;
  my_printf_option[9] = &my_printf_option_b;
  my_printf_option[10] = &my_printf_option_p;
  my_printf_option[11] = '\0';
}

void	my_printf_option(va_list ap, char option)
{
  void	(*my_printf_option[12])(va_list s_ap);
  char	*tab_option;
  int	i;

  i = 0;
  tab_option = "sSidcxXoubp";
  if (option == '%')
    my_putchar('%');
  else
    {
      my_printf_set_option(my_printf_option);
      while ((tab_option[i] - option) != 0 && tab_option[i] != '\0')
	i = i + 1;
      if (tab_option[i] == '\0')
	{
	  my_putchar('%');
	  my_putchar(option);
	}
      else
	my_printf_option[i](ap);
    }
}

void	my_printf(char *format, ...)
{
  va_list	ap;
  int	len;
  int	i;

  i = 0;
  va_start(ap, format);
  len = my_strlen(format);
  while (i <= len)
    {
      if (format[i] == '%')
	{
	  i = i + 1;
	  my_printf_option(ap, format[i]);
	}
      else
	my_putchar(format[i]);
      i = i + 1;
    }
}
