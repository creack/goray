/*
** calc_sphere.c for rtv1 in /home/fortin_j//afs/projets/rtv1/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 10:47:31 2012 julien fortin
** Last update Sun Mar 11 03:07:10 2012 julien fortin
*/

#include	<stdlib.h>
#include	<math.h>
#include	<libmy.h>
#include	<rtv1.h>

void	change_scene(t_scene *scene, t_obj *obj)
{
  if (obj->x != 0)
    scene->x_oeil -= obj->x;
  else
    scene->x_oeil += obj->x;
  if (obj->y != 0)
    scene->y_oeil -= obj->y;
  if (obj->z > 0)
    scene->z_oeil -= obj->z;
  else
    scene->z_oeil += obj->z;
}

t_scene	*save_scene(t_scene *scene, t_obj *obj, int i)
{
  static double	x_oeil;
  static double	y_oeil;
  static double	z_oeil;

  if (i == 0)
    {
      x_oeil = scene->x_oeil;
      y_oeil = scene->y_oeil;
      z_oeil = scene->z_oeil;
      change_scene(scene, obj);
    }
  else
    {
      scene->x_oeil = x_oeil;
      scene->y_oeil = y_oeil;
      scene->z_oeil = z_oeil;
    }
  return (scene);
}

t_inter	*add_inter_sphere(t_scene *scene, double lambda)
{
  t_inter	*inter;

  inter = xmalloc(sizeof(*inter));
  inter->x = (lambda * scene->vx) + scene->x_oeil;
  inter->y = (lambda * scene->vy) + scene->y_oeil;
  inter->z = (lambda * scene->vz) + scene->z_oeil;
  return (inter);
}

double	calc_sphere(t_scene *scene, t_obj *obj)
{
  double	a;
  double	b;
  double	c;
  double	delta;
  double	lambda;

  scene = save_scene(scene, obj, 0);
  a = SQR(scene->vx) + SQR(scene->vy) + SQR(scene->vz);
  b = 2 * ((scene->x_oeil * scene->vx) + (scene->y_oeil * scene->vy) + \
	   (scene->z_oeil * scene->vz));
  c = SQR(scene->x_oeil) + SQR(scene->y_oeil) + SQR(scene->z_oeil)
    - SQR(obj->r);
  delta = SQR(b) - (4 * (a * c));
  if (delta < 0)
    {
      scene = save_scene(scene, obj, 1);
      return (-1.0);
    }
  else
    {
      lambda = (-b - sqrt(delta)) / (2 * a);
      obj->inter = add_inter_sphere(scene, lambda);
      scene = save_scene(scene, obj, 1);
      return (lambda);
    }
}
