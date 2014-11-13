/*
** my_init_cyl_threreturn.c for rtv1 in /home/fortin_j//afs/svn/rtv1-2016-fortin_j/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Mar 11 16:02:47 2012 julien fortin
** Last update Sun Mar 11 16:10:34 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>
#include	<rtv1.h>

t_obj	*cyl_6(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = CYL;
  obj->x = 2400;
  obj->y = 50;
  obj->z = 10;
  obj->r = 50;
  obj->x_rot= 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00F4661B;
  obj->bright = 0.0;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*cyl_7(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = CYL;
  obj->x = 2400;
  obj->y = -150;
  obj->z = 10;
  obj->r = 50;
  obj->x_rot= 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00F4661B;
  obj->bright = 0.0;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}
