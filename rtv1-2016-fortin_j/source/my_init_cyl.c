/*
** my_init_cyl.c for rtv1 in /home/fortin_j//afs/svn/rtv1-2016-fortin_j/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Mar 11 15:14:00 2012 julien fortin
** Last update Sun Mar 11 16:02:04 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>
#include	<rtv1.h>

t_obj	*cyl_1(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = CYL;
  obj->x = 19000;
  obj->y = 2500;
  obj->z = 10;
  obj->r = 200;
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

t_obj	*cyl_2(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = CYL;
  obj->x = 19000;
  obj->y = 2000;
  obj->z = 10;
  obj->r = 200;
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

t_obj	*cyl_3(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = CYL;
  obj->x = 19000;
  obj->y = 1500;
  obj->z = 10;
  obj->r = 200;
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

t_obj	*cyl_4(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = CYL;
  obj->x = 19000;
  obj->y = 1000;
  obj->z = 10;
  obj->r = 200;
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

t_obj	*cyl_5(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = CYL;
  obj->x = 19000;
  obj->y = 500;
  obj->z = 10;
  obj->r = 200;
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
