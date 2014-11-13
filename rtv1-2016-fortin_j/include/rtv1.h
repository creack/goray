/*
** rtv1.h for rtv1 in /home/fortin_j//afs/projets/rtv1
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 09:37:10 2012 julien fortin
** Last update Sun Mar 11 16:06:13 2012 julien fortin
*/

#ifndef	__RTV1__
#define	__RTV1__

#include	<s_rtv1.h>

int	main();
int     my_key_hook(int kc, void *param);
int     my_expose_hook(t_expose *expose);
int	calc(t_expose *expose, int x, int y);
double	calc_plan(t_scene *scene, t_obj *obj);
double	calc_sphere(t_scene *scene, t_obj *obj);
double	calc_cyl(t_scene *scene, t_obj *obj);
double	calc_cone(t_scene *scene, t_obj *obj);
void    my_pixel_put_to_image(t_expose *expose, int x, int y, int col);
t_mlx   *my_init_mlx(int x, int y, char *name);
t_img   *my_init_img(int x, int y, t_mlx *mlx);
t_obj   *my_init_obj();
t_spot	*my_init_spot();
t_scene	*save_scene(t_scene *scene, t_obj *obj, int i);
t_scene *my_init_scene(int x, int y);
t_expose        *my_init_expose(t_mlx *mlx, t_img *img);
void	my_free_rtvlol(t_expose *expose);
t_obj	*sphere_1(t_obj *prev);
t_obj	*sphere_2(t_obj *prev);
t_obj	*sphere_3(t_obj *prev);
t_obj	*sphere_4(t_obj *prev);
t_obj	*sphere_5(t_obj *prev);
t_obj	*sphere_6(t_obj *prev);
t_obj	*sphere_7(t_obj *prev);
t_obj	*sphere_8(t_obj *prev);
t_obj	*sphere_9(t_obj *prev);
t_obj	*sphere_10(t_obj *prev);
t_obj	*sphere_11(t_obj *prev);
t_obj	*sphere_12(t_obj *prev);
t_obj	*cyl_1(t_obj *prev);
t_obj	*cyl_2(t_obj *prev);
t_obj	*cyl_3(t_obj *prev);
t_obj	*cyl_4(t_obj *prev);
t_obj	*cyl_5(t_obj *prev);
t_obj	*cyl_6(t_obj *prev);
t_obj	*cyl_7(t_obj *prev);
t_obj	*cone_1(t_obj *prev);

#endif
