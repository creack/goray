/*
** my_init.c for rtv1 in /home/fortin_j//afs/projets/rtv1
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sat Dec 24 21:13:11 2011 julien fortin
** Last update Sun Mar 11 15:09:51 2012 julien fortin
*/

#include	<stdlib.h>
#include	<mlx.h>
#include        <libmy.h>
#include	<rtv1.h>

t_mlx	*my_init_mlx(int x, int y, char *name)
{
  t_mlx	*mlx;
  void	*mlx_ptr;
  void	*win_ptr;

  mlx_ptr = mlx_init();
  if (mlx_ptr == NULL)
    xexit("Fatal error: mlx_init return NULL");
  win_ptr = mlx_new_window(mlx_ptr, x, y, name);
  mlx = malloc(sizeof(*mlx));
  if (mlx == NULL)
    xexit("Fatal error: cannot allocate memory to initialize mlx.\n");
  mlx->mlx_ptr = mlx_ptr;
  mlx->win_ptr = win_ptr;
  return (mlx);
}

t_img	*my_init_img(int x, int y, t_mlx *mlx)
{
  t_img	*img;
  void	*img_ptr;
  char	*data;
  int	bytes;
  int	sizeline;
  int	endian;

  img_ptr = mlx_new_image(mlx->mlx_ptr, x, y);
  data = mlx_get_data_addr(img_ptr, &bytes, &sizeline, &endian);
  img = malloc(sizeof(*img));
  if (img == NULL)
    xexit("Fatal error: cannot allocate memory to initialize image.\n");
  img->img_ptr = img_ptr;
  img->data = data;
  img->sizeline = sizeline;
  img->endian = endian;
  img->bytes = bytes / 8;
  img->x = x;
  img->y = y;
  return (img);
}

t_scene	*my_init_scene(int x, int y)
{
  t_scene	*scene;

  scene = xmalloc(sizeof(*scene));
  scene->x_oeil = -500.0 * 10;
  scene->y_oeil = 0.0 * 10;
  scene->z_oeil = 30.0 * 10;
  scene->x1 = D;
  scene->y1 = (WIN_X / 2.0) - x;
  scene->z1 = (WIN_Y / 2.0) - y;
  scene->vx = D - scene->x_oeil;
  scene->vy = scene->y1;
  scene->vz = scene->z1;
  return (scene);
}

t_expose	*my_init_expose(t_mlx *mlx, t_img *img)
{
  t_expose	*expose;

  expose = malloc(sizeof(*expose));
  if (expose == NULL)
    xexit("Fatal error: cannot allocate memory to initialize expose.\n");
  expose->t_mlx = mlx;
  expose->t_img = img;
  expose->t_obj = my_init_obj();
  expose->t_spot = my_init_spot();
  return (expose);
}
