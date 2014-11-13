/*
** mlx_ppti.c for rtv1 in /home/fortin_j//afs/projets/rtv1/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 09:31:49 2012 julien fortin
** Last update Fri Feb 17 10:07:29 2012 julien fortin
*/

#include	<mlx.h>
#include	<libmy.h>
#include	<rtv1.h>

void    my_pixel_put_to_image(t_expose *expose, int x, int y, int col)
{
  t_color       color;
  unsigned int  rgb;
  int   line;
  int   pixel;
  int   i;

  if ((x < 0 || y < 0) || (x > WIN_X || y > WIN_Y))
    xexit("Fatal error: pixel out of image.\n");
  line = expose->t_img->sizeline * y;
  pixel = expose->t_img->bytes * x;
  i = 0;
  rgb = mlx_get_color_value(expose->t_mlx->mlx_ptr, col);
  color.color = rgb;
  if (expose->t_img->endian == 0)
    {
      expose->t_img->data[line + pixel + i] = color.rgb[0];
      i = i + 1;
      expose->t_img->data[line + pixel + i] = color.rgb[1];
      i = i + 1;
      expose->t_img->data[line + pixel + i] = color.rgb[2];
      i = i + 1;
      expose->t_img->data[line + pixel + i] = color.rgb[3];
    }
}
