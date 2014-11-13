/*
** my_init_sphere_thereturn_and_again.c for rtv1 in /home/fortin_j//afs/svn/rtv1-2016-fortin_j/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Mar 11 14:59:29 2012 julien fortin
** Last update Sun Mar 11 15:01:51 2012 julien fortin
*/

#include	<stdlib.h>
#include	<libmy.h>
#include	<rtv1.h>

t_obj	*sphere_11(t_obj *prev)
{
  t_obj*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 1000;
  obj->y = 230;
  obj->z = 150;
  obj->r = 20;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00FFFFFF;
  obj->bright = 0.9;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

t_obj	*sphere_12(t_obj *prev)
{
  t_obj*obj;

  obj = xmalloc(sizeof(*obj));
  obj->type = SPH;
  obj->x = 1420;
  obj->y = 295;
  obj->z = 48;
  obj->r = 108;
  obj->x_rot = 0.0;
  obj->y_rot = 0.0;
  obj->z_rot = 0.0;
  obj->angle = 0;
  obj->color = 0x00DD985C;
  obj->bright = 0.9;
  obj->k = -1.0;
  obj->inter = NULL;
  obj->next = prev;
  return (obj);
}

