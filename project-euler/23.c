#include <stdio.h>
#include <stdbool.h>

int prop(int n){
   int res=1;
   for(int i=2; i<=n/2; i++)if(n%i == 0)res+=i;
   return res;
}

bool isNumInArr(int n, int *arr, size_t size){
   for(size_t i=0; i < size; i++)
      if(arr[i] == n)return true;
   return false;
}

int main(){
   int arr[28123] = {0}, arrlen=0;
   int arr1[28123] = {0}, arr1len=0;

   for(int i=2; i <28123;i++)
      if(prop(i) > i){arr[arrlen++]=i;}
      
   for(int i=0; i < arrlen;i++)
      for(int j=0; j <= i;j++){
         int s =  arr[i]+arr[j];
         arr1[arr1len++] = s;
         if(s > 28123)break;
      }
   

   printf("len:%d len2:%d",arrlen ,arr1len);
   return 0;
}
