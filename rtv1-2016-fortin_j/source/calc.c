/*
** calc.c for rtv1 in /home/fortin_j//afs/projets/rtv1
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 10:16:01 2012 julien fortin
** Last update Sun Mar 11 15:57:40 2012 julien fortin
*/

#include	<math.h>
#include	<stdlib.h>
#include	<rtv1.h>

#define	STOP	10000000.0

int	brightness(int color, int spot_color, double coef, double a)
{
  t_color	spot;
  t_color	obj;

  obj.color = color;
  spot.color = spot_color;
  if ((obj.rgb[0] + (coef * spot.rgb[0] * a)) > 255)
    obj.rgb[0] = 255;
  else
    obj.rgb[0] = (obj.rgb[0] + (coef * spot.rgb[0] * a));
  if ((obj.rgb[1] + (coef * spot.rgb[1] * a)) > 255)
    obj.rgb[1] = 255;
  else
    obj.rgb[1] = (obj.rgb[1] + (coef * spot.rgb[1] * a));
  if ((obj.rgb[2] + (coef * spot.rgb[2] * a)) > 255)
    obj.rgb[2] = 255;
  else
    obj.rgb[2] = (obj.rgb[2] + (coef * spot.rgb[2] * a));
  if ((obj.rgb[3] + (coef * spot.rgb[3] * a)) > 255)
    obj.rgb[3] = 255;
  else
    obj.rgb[3] = (obj.rgb[3] + (coef * spot.rgb[3] * a));
  return (obj.color);
}

int		my_light(t_obj *obj, t_spot *spot)
{
  t_color	color;
  double	lx;
  double	ly;
  double	lz;
  double	n;
  double	a;

  lx = obj->inter->x - spot->x;
  ly = obj->inter->y - spot->y;
  lz = obj->inter->z - spot->z;
  n = sqrt(SQR(obj->inter->x) + SQR(obj->inter->y) + SQR(obj->inter->z));
  a = ((obj->inter->x * lx) + (obj->inter->y * ly) + (obj->inter->z * lz)) \
    / (n * (sqrt(SQR(lx) + SQR(ly) + SQR(lz))));
  free(obj->inter);
  if (a < 0)
    return (0);
  color.color = obj->color;
  color.rgb[0] *= a;
  color.rgb[1] *= a;
  color.rgb[2] *= a;
  color.rgb[3] *= a;
  return (brightness(color.color, spot->color, obj->bright, a));
}

int		get_color(t_obj *obj, t_spot *spot)
{
  t_obj		*tmp;
  double	k;

  k = STOP;
  tmp = NULL;
  while (obj != NULL)
    {
      if (obj->k <= k && obj->k != -1.0 && obj->k >= 0.0)
	{
	  tmp = obj;
	  k = obj->k;
	}
      obj = obj->next;
    }
  if (k == STOP)
    return (0);
  return (my_light(tmp, spot));
}

int		calc(t_expose *expose, int x, int y)
{
  t_scene	*scene;
  t_obj		*obj;

  obj = expose->t_obj;
  scene = my_init_scene(x, y);
  while (obj != NULL)
    {
      if (obj->type == 0)
	obj->k = calc_plan(scene, obj);
      else if (obj->type == 1)
	obj->k = calc_sphere(scene, obj);
      else if (obj->type == 2)
       	obj->k = calc_cyl(scene, obj);
      else if (obj->type == 3)
	obj->k = calc_cone(scene, obj);
      obj = obj->next;
    }
  free(scene);
  return (get_color(expose->t_obj, expose->t_spot));
}
