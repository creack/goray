/*
** my_init_sphere.c for rtv1 in /home/fortin_j//afs/svn/rtv1-2016-fortin_j/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Mar 11 14:10:39 2012 julien fortin
** Last update Mon Mar 12 10:46:35 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>
#include	<rtv1.h>

t_obj	*sphere_1(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 10000;
  obj->y = -1100;
  obj->z = 1100;
  obj->r = 200;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0.0;
  obj->color = 0x00FFFFFF;
  obj->bright = 0.9;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_2(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 1500;
  obj->y = 300;
  obj->z = 0;
  obj->r = 200;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00BBBBBB;
  obj->bright = 0.9;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_3(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 2500;
  obj->y = -50;
  obj->z = 400;
  obj->r = 200;
  obj->x_rot = 0;
  obj->y_rot = 0;
  obj->z_rot = 0;
  obj->angle = 0;
  obj->color = 0x00FF0000;
  obj->bright = 0.9;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_4(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 2500;
  obj->y = -50;
  obj->z = 600;
  obj->r = 0;
  obj->x_rot = 0;
  obj->y_rot = 0;
  obj->z_rot = 0;
  obj->angle = 0.0;
  obj->color = 0x000000FF;
  obj->bright = 0.9;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_5(t_obj *prev)
{
  t_obj	*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 1500;
  obj->y = 400;
  obj->z = 100;
  obj->r = 100;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00BBBBBB;
  obj->bright = 0.9;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}
