/*
** my_init_cone.c for rtv1 in /home/fortin_j//afs/svn/rtv1-2016-fortin_j/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Mar 11 15:16:32 2012 julien fortin
** Last update Sun Mar 11 15:19:03 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>
#include	<rtv1.h>

t_obj	*cone_1(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = CONE;
  obj->x = 9950;
  obj->y = -1100;
  obj->z = 1100;
  obj->r = 100;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = RAD(50.0);
  obj->color = 0x00568203B;
  obj->bright = 0.5;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}
