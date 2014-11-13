##
## Makefile for libmy in /home/fortin_j/libmy/
##
## Made by julien fortin
## Login   <fortin_j@epitech.net>
##
## Started on  Thu Oct 27 09:23:21 2011 julien fortin
## Last update Fri Feb 17 10:01:07 2012 julien fortin
##

RM	= rm -f

NAME	= libmy

CC	= gcc

CFLAGS	= -W -Wall -Wextra -I.

SRC	= my_convert_base.c \
	  my_errors.c \
	  my_free.c \
	  my_nbr.c \
	  my_size.c \
	  my_printf.c \
	  my_printf_option_base.c \
	  my_printf_option_base_next.c \
	  my_printf_option.c \
	  my_string.c \
	  my_string_dup.c \
	  my_strcmp.c

OBJ	= $(SRC:.c=.o)

all: $(NAME)

$(NAME): $(OBJ)
	ar rc $(NAME).a $(OBJ)
	ranlib $(NAME).a

clean:
	$(RM) $(OBJ)

fclean: clean
	$(RM) $(NAME)

re:	fclean all