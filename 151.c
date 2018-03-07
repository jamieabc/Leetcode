/* Given an input string, reverse the string word by word. */

/* For example, */
/* Given s = "the sky is blue", */
/* return "blue is sky the". */

/* Update (2015-02-12): */
/* For C programmers: Try to solve it in-place in O(1) space. */

#include <stdio.h>
#include <string.h>

/* copy str from src to desc */
void copy_str(char *src, char *desc) {
  int idx = 0;
  while (desc[idx] != '\0') {
    *(desc + idx) = *(src + idx);
    idx++;
  }
  desc[idx] = '\0';
}

void reverseWords(char *s) {

  /* i records current location of space, p precords previous space location */
  /* both i and p set to last non-space location of string */
  /* go backward from string until it meets space */
  /* copy i-p string to new str */
  /* set p to i, loop until p reaches start of string */

  int length = strlen(s);

  /* return if null character */
  if (length == 0) {
    return;
  }

  int end = length - 1;
  int begin = 0;

  /* trim spaces at end */
  while(end >= 0 && s[end] == ' ') {
    end--;
  }

  /* trim spaces at beginning */
  while(begin < end && s[begin] == ' ') {
    begin++;
  }

  /* return null character if all spaces */
  if (end < 0) {
    *s = '\0';
    return;
  }

  s[end + 1] = '\0';

  int i = end;
  int p = end;
  int new = 0;

  char new_str[end - begin + 2];
  *(new_str + end - begin + 1) = '\0';

  while (i >= begin) {
    /* find next space backward */
    while (s[i] != ' ' && i >= begin) {
      i--;
    }

    /* copy string */
    for (int k = i + 1; k <= p && s[k] != ' '; ++k) {
      new_str[new++] = s[k];
    }

    /* find next non-space character backward */
    while(s[i] == ' ' && i > begin) {
      i--;
    }

    /* add space */
    if (i >= begin) {
      new_str[new++] = ' ';
    }

    p = i + 1;
  }

  new_str[new] = '\0';

  copy_str(new_str, s);
  return;
}

int main() {
  char s[] = "  a  b ";

  reverseWords(s);
}
