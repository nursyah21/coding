#include <stdio.h>
#include <math.h>
#define uint unsigned int

unsigned int clen(char *c){
  unsigned int count=0,n=0;
  while(c[n++])
    count++;
  
  return count;
}

int main(){
  //int n=0;
	 char *name;
  name="";
  printf("%s %d\n",name, clen(name));
  
  return 0;
}
