/*
** calc_cone.c for rtv1 in /home/fortin_j//afs/projets/rtv1/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Thu Mar  8 20:25:34 2012 julien fortin
** Last update Sun Mar 11 02:50:56 2012 julien fortin
*/

#include	<math.h>
#include	<libmy.h>
#include	<rtv1.h>

t_inter	*add_inter_cone(t_scene *scene, t_obj *obj, double lambda)
{
  t_inter	*inter;

  inter = xmalloc(sizeof(*inter));
  inter->x = (lambda * scene->vx) + scene->x_oeil;
  inter->y = (lambda * scene->vy) + scene->y_oeil;
  inter->z = -obj->angle * ((lambda * scene->vz) + scene->z_oeil);
  scene = save_scene(scene, obj, 1);
  return (inter);
}

double	calc_cone(t_scene *scene, t_obj *obj)
{
  double	a;
  double	b;
  double	c;
  double	delta;
  double	lambda;

  scene = save_scene(scene, obj, 0);
  a = SQR(scene->vx) + SQR(scene->vy) - (SQR(scene->vz) /
					 SQR(tan(obj->angle)));
  b = 2 * ((scene->vx * scene->x_oeil) + (scene->vy * scene->y_oeil) -
	   ((scene->vz * scene->z_oeil) / SQR(tan(obj->angle))));
  c = SQR(scene->x_oeil) + SQR(scene->y_oeil) - ((1 / SQR(tan(obj->angle)))
						 * SQR(scene->z_oeil));
  delta = SQR(b) - (4 * (a * c));
  if (delta < 0)
    {
      scene = save_scene(scene, obj, 1);
      return (-1.0);
    }
  else
    {
      lambda = (-b - sqrt(delta)) / (2 * a);
      obj->inter = add_inter_cone(scene, obj, lambda);
      return (lambda);
    }
}
