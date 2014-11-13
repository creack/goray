/*
** libmy.h for libmy in /home/fortin_j//afs/projets/rtv1/source/libmy
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 10:02:39 2012 julien fortin
** Last update Fri Feb 17 10:02:47 2012 julien fortin
*/

#ifndef __LIBMY_H__
#define __LIBMY_H__

#include	<stdarg.h>

char	*my_convert_base(unsigned int nb, char *base);
char	*my_strcat(char *str1, char *str2);
char    *my_strdup(char *src);
char    *my_strlowcase(char *str);
char    *my_strndup(char *src, int i, int nb_char);
int     my_char_isnum(char c);
int     my_getnbr(char *str);
int	my_intlen(int nb);
int     my_put_nbr(int nb);
int	my_strcmp(char *s1, char *s2);
int     my_strlen(char *str);
int	my_strncmp(char *s1, char *s2, int n);
int	my_tputs(int c);
int     my_wtlen(char **wt);
void	my_free_str(char *s1, char *s2, char *s3, char *s4);
void	my_free_wt(char **wt);
void	my_printf(char *format, ...);
void    my_printf_option_b(va_list ap);
void	my_printf_option_c(va_list ap);
void	my_printf_option_i(va_list ap);
void    my_printf_option_o(va_list ap);
void	my_printf_option_p(va_list ap);
void	my_printf_option_ss(va_list ap);
void	my_printf_option_s(va_list ap);
void    my_printf_option_u(va_list ap);
void	my_printf_option(va_list ap, char option);
void	my_printf_option_x(va_list ap);
void    my_printf_option_xx(va_list ap);
void	my_printf_set_option(void (**my_printf_option)(va_list ap));
void	my_putchar(char c);
void	my_putchar_error(char c);
void	my_put_nbrbase_long(long nb, char *base);
void    my_put_nbrbase(unsigned int nb, char *base);
void	my_putstr(char *str);
void	my_putstr_error(char *str);
void	xexit(char *str);
void    *xmalloc(int size);

#endif
