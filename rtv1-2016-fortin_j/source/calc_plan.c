/*
** calc_plan.c for rtv1 in /home/fortin_j//afs/projets/rtv1/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 11:19:13 2012 julien fortin
** Last update Sun Mar 11 03:48:37 2012 julien fortin
*/

#include	<libmy.h>
#include	<rtv1.h>

double	calc_plan(t_scene *scene, t_obj *obj)
{
  t_inter	*inter;
  double	k;

  k = -1.0;
  if (scene->vz != 0)
    k = (-scene->z_oeil) / scene->vz;
  obj->inter = xmalloc(sizeof(*inter));
  obj->inter->x = 0;
  obj->inter->y = 0;
  obj->inter->z = scene->z_oeil;
  return (k);
}
