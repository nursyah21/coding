#include <stdio.h>


int chain(long n){ //must use long
  int k=1;
  while (n != 1){
    if(n%2)n = 3*n+1;
    else n /= 2;
    k++;
  }
  return k;
}

int main(){
  int n=1000000,max_=0,i_=0;
  
  for(int i=2; i<=n;i++){
    int res=chain(i);

    if(res > max_){
       max_=res;
       i_=i;
    }
  }

  printf("chain:%d result:%d\n", i_,max_);
  return 0;
}
