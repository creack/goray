/*
** main.c for rtv1 in /home/fortin_j//afs/projets/rtv1/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 09:24:36 2012 julien fortin
** Last update Sun Mar 11 00:24:54 2012 julien fortin
*/

#include	<stdlib.h>
#include	<mlx.h>
#include	<libmy.h>
#include	<rtv1.h>

void    my_set_img(t_expose *expose)
{
  int   x;
  int   y;

  x = 0;
  y = 0;
  while (y <= WIN_Y)
    {
      while (x <= WIN_X)
        {
          my_pixel_put_to_image(expose, x, y, calc(expose, x, y));
          x++;
        }
      mlx_put_image_to_window(expose->t_mlx->mlx_ptr, expose->t_mlx->win_ptr,
			      expose->t_img->img_ptr, 0, 0);
      if (!(y % (WIN_Y / 10)))
        my_printf("\b\b\b%d%", y / (WIN_Y / 10) * 10);
      x = 0;
      y++;
    }
}

int     main()
{
  t_mlx *mlx;
  t_img *img;
  t_expose      *expose;

  if ((mlx = my_init_mlx(WIN_X, WIN_Y, "RTv1")) == NULL)
    return (0);
  img = my_init_img(WIN_X, WIN_Y, mlx);
  expose = my_init_expose(mlx, img);
  my_set_img(expose);
  mlx_put_image_to_window(mlx->mlx_ptr, mlx->win_ptr, img->img_ptr, 0, 0);
  mlx_key_hook(mlx->win_ptr, &my_key_hook, expose);
  mlx_expose_hook(mlx->win_ptr, &my_expose_hook, expose);
  mlx_loop(mlx->mlx_ptr);
  return (0);
}
