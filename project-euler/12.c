#include <stdio.h>
#include <math.h>

int div(int n){
  int k=1;
  for(int i=1; i <= n/2; i++){
    if(n%i) continue;
	   k+=1;
  }
  return k;
}

int div2(int n){
  int k=1;
  for(int i=1; i <= (int) sqrt(n); i++){
  	  if(n%i) continue;
	   k+=1;
  }
  if (k%2 == 0) return k*2;
  return (k-1)*2;
}


int anum(int n){
  int a=0;
  for(int i=1; i <=n; i++)a+=i;
  return a;
}

int main(){
  int n=0;
  int a = anum(n);
	
  for(int i=n+1; i < 100000; i++){
    a+=i;
    int res=div2(a);
	   if(res >= 500){
	     printf("i: %d a:%d res2:%d\n",i,a,res);
	     break;
   	 }
  }
  return 0;
}
