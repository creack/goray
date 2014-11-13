##
## Makefile for rtv1 in /home/fortin_j//afs/projets/rtv1
##
## Made by julien fortin
## Login   <fortin_j@epitech.net>
##
## Started on  Fri Feb 17 09:24:51 2012 julien fortin
## Last update Sun Mar 11 16:05:48 2012 julien fortin
##

NAME	= rtv1
CC	= gcc
CFLAGS	= -W -Wall -Wextra -I./include/ -Ilib/minilibx -I/opt/X11/include  -O3
DIR	= ./source/
LDFLAGS	= -Llib/libmy/ -Llib/minilibx -lmy -lmlx -L/opt/X11/lib -L/usr/lib64/X11 -lXext -lX11 -lm
SRCS	= $(DIR)main.c 					\
	  $(DIR)mlx_ppti.c 				\
	  $(DIR)my_expose.c 				\
	  $(DIR)my_free.c 				\
	  $(DIR)my_init.c 				\
	  $(DIR)my_init_obj.c 				\
	  $(DIR)my_init_sphere.c 			\
	  $(DIR)my_init_cyl.c 				\
	  $(DIR)my_init_cyl_thereturn.c 		\
	  $(DIR)my_init_cone.c 				\
	  $(DIR)my_init_sphere_thereturn.c 	 	\
	  $(DIR)my_init_sphere_thereturn_and_again.c	\
	  $(DIR)my_init_spot.c 				\
	  $(DIR)calc.c 					\
	  $(DIR)calc_sphere.c 				\
	  $(DIR)calc_plan.c 				\
	  $(DIR)calc_cone.c 				\
	  $(DIR)calc_cyl.c
OBJS	= $(SRCS:.c=.o)

all	: $(NAME)

$(NAME)	: $(OBJS)
	$(MAKE) -C lib/libmy/
	$(MAKE) -C lib/minilibx/
	$(CC) -o $(NAME) $(OBJS) $(LDFLAGS)

clean	:
	$(MAKE) -C lib/libmy/ clean
	$(MAKE) -C lib/minilibx/ clean
	$(RM) $(OBJS)

fclean	: clean
	$(RM) $(NAME)

re	: fclean all

