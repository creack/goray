/*
** my_free.c for rtv1 in /home/fortin_j//afs/svn/rtv1-2016-fortin_j/source
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Mar 11 03:56:49 2012 julien fortin
** Last update Sun Mar 11 05:11:25 2012 julien fortin
*/

#include	<stdlib.h>
#include	<rtv1.h>

void	free_obj(t_obj *obj)
{
  t_obj	*tmp;

  tmp = NULL;
  while (obj != NULL)
    {
      if (tmp != NULL)
	free(tmp);
      tmp = obj;
      obj = obj->next;
    }
}

void	my_free_rtvlol(t_expose *expose)
{
  free(expose->t_mlx->mlx_ptr);
  free(expose->t_mlx->win_ptr);
  free(expose->t_mlx);
  free(expose->t_img);
  free_obj(expose->t_obj);
  free(expose->t_spot);
}
