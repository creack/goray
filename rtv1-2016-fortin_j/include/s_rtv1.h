/*
** s_rtv1.h for rtv1 in /home/fortin_j//afs/projets/rtv1
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Fri Feb 17 09:26:05 2012 julien fortin
** Last update Sun Mar 11 14:04:54 2012 julien fortin
*/

#ifndef	__S_RTV1__
# define __S_RTV1__

# define WIN_X	1000
# define WIN_Y	1000
# define D	-1000

# define SQR(x)	((x) * (x))
# define RAD(x)	(((x) * (3.14159265)) / 180)

enum {
  PLAN,
  SPH,
  CYL,
  CONE,
};

typedef struct	s_mlx
{
  void		*mlx_ptr;
  void		*win_ptr;
}		t_mlx;

typedef struct	s_inter
{
  double	x;
  double	y;
  double	z;
  double	k;
}		t_inter;

typedef struct	s_spot
{
  double	x;
  double	y;
  double	z;
  int		color;
  struct s_spot	*next;
}		t_spot;

typedef struct	s_img
{
  void		*img_ptr;
  char		*data;
  int		sizeline;
  int		endian;
  int		bytes;
  int		x;
  int		y;
}		t_img;

typedef struct		s_obj
{
  int			type;
  int			x;
  int			y;
  int			z;
  int			r;
  int			x_rot;
  int			y_rot;
  int			z_rot;
  int			color;
  double		bright;
  double		angle;
  double		k;
  struct s_obj		*next;
  struct s_inter	*inter;
}			t_obj;

typedef struct	s_scene
{
  double	x_oeil;
  double	y_oeil;
  double	z_oeil;
  double	x1;
  double	y1;
  double	z1;
  double	vx;
  double	vy;
  double	vz;
}		t_scene;

typedef struct	s_expose
{
  t_mlx		*t_mlx;
  t_img		*t_img;
  t_obj		*t_obj;
  t_spot	*t_spot;
}		t_expose;

typedef union	u_color
{
  unsigned int	color;
  unsigned char	rgb[4];
}		t_color;

#endif /* !__S_RTV1__ */
