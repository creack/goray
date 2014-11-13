/*
** my_init_sphere_thereturn.c for rtv1 in /home/fortin_j//afs/svn/rtv1-2016-fortin_j/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Mar 11 14:48:51 2012 julien fortin
** Last update Sun Mar 11 15:08:49 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>
#include	<rtv1.h>

t_obj	*sphere_6(t_obj *prev)
{
  t_obj*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 1400;
  obj->y = 420;
  obj->z = 120;
  obj->r = 50;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00DD985C;
  obj->bright = 0.2;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_7(t_obj *prev)
{
  t_obj*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 1500;
  obj->y = 200;
  obj->z = 100;
  obj->r = 100;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00BBBBBB;
  obj->bright = 0.2;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_8(t_obj *prev)
{
  t_obj*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 1400;
  obj->y = 180;
  obj->z = 120;
  obj->r = 50;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00DD985C;
  obj->bright = 0.5;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_9(t_obj *prev)
{
  t_obj*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 2000;
  obj->y = 710;
  obj->z = 120;
  obj->r = 250;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00156D9A;
  obj->bright = 0.5;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_10(t_obj *prev)
{
  t_obj*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 1000;
  obj->y = 300;
  obj->z = 150;
  obj->r = 20;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00FFFFFF;
  obj->bright = 0.5;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}
