/*
** my_wolf_expose.c for wolf3d in /home/fortin_j//afs/projets/wolf3d/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sat Dec 24 21:11:44 2011 julien fortin
** Last update Sun Mar 11 03:54:59 2012 julien fortin
*/

#include	<stdlib.h>
#include	<mlx.h>
#include        <libmy.h>
#include	<rtv1.h>

int     my_key_hook(int kc, void *param)
{
  t_expose	*expose;

  expose = param;
  if (kc == 65307)
    {
      my_free_rtvlol(expose);
      exit(EXIT_SUCCESS);
    }
  else if (kc == 65363 || kc == 65361 || kc == 100 || kc == 97)
    {
      mlx_put_image_to_window(expose->t_mlx->mlx_ptr,
			      expose->t_mlx->win_ptr,
			      expose->t_img->img_ptr, 0, 0);
    }
  return (0);
}

int	my_expose_hook(t_expose *expose)
{
  mlx_put_image_to_window(expose->t_mlx->mlx_ptr, \
			  expose->t_mlx->win_ptr, \
			  expose->t_img->img_ptr, 0, 0);
  return (0);
}
