/*
** my_init_obj.c for rtv1 in /home/fortin_j//afs/projets/rtv1/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Mar 11 03:00:58 2012 julien fortin
** Last update Sun Mar 11 16:11:32 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>
#include	<rtv1.h>

t_obj	*my_init_plan(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = PLAN;
  obj->x = 0;
  obj->y = 0;
  obj->z = 0;
  obj->r = 0;
  obj->x_rot = 0;
  obj->y_rot = 0;
  obj->z_rot = 0;
  obj->angle = 0;
  obj->color = 0x000000FF;
  obj->bright = 0;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj   *my_init_obj()
{
  t_obj *obj;

  obj = NULL;
  obj = my_init_plan(obj);
  obj = sphere_1(obj);
  obj = sphere_2(obj);
  obj = sphere_3(obj);
  obj = sphere_4(obj);
  obj = sphere_5(obj);
  obj = sphere_6(obj);
  obj = sphere_7(obj);
  obj = sphere_8(obj);
  obj = sphere_9(obj);
  obj = sphere_10(obj);
  obj = sphere_11(obj);
  obj = sphere_12(obj);
  obj = cyl_1(obj);
  obj = cyl_2(obj);
  obj = cyl_3(obj);
  obj = cyl_4(obj);
  obj = cyl_6(obj);
  obj = cyl_7(obj);
  obj = cone_1(obj);
  return (obj);
}
