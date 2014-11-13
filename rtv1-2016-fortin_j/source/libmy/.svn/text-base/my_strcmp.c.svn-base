/*
** my_strcmp.c for libmy in /home/fortin_j/libmy/
**
** Made by julien fortin
** Login   <fortin_j@epitech.net>
**
** Started on  Sun Jan  8 17:30:00 2012 julien fortin
** Last update Tue Jan 17 18:28:43 2012 julien fortin
*/

int	my_strcmp(char *s1, char *s2)
{
  int   i;
  int   sum;

  i = 0;
  if (s1[i] == s2[i])
    {
      while (s1[i] && s2[i] && s1[i] == s2[i] && s1[i] != '\0')
	i = i + 1;
      if (i != 0)
	{
	  if (s1[i - 1] == s2[i - 1])
	    sum = s1[i] - s2[i];
	}
      else
	return (0);
    }
  else
    sum = s1[i] - s2[i];
  return (sum);
}

int	my_strncmp(char *s1, char *s2, int n)
{
  int   i;
  int   sum;

  i = 0;
  if ((s1[i] == s2[i]) && (i + 1 != n))
    {
      while ((s1[i] == s2[i]) && (i + 1 != n))
        i = i + 1;
      sum = s1[i] - s2[i];
    }
  else
    sum = s1[i] - s2[i];
  return (sum);
}
