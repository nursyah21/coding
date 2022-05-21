#include <stdio.h>
#include <math.h>
#include <string.h>

int main(){
   int res =0;
   for(int i=10; i < 6*pow(9,5);i++){
      int a=0;char num[10];

      sprintf(num,"%d",i);
      for(int i=0; i < strlen(num); i++)
         a+= pow((int)num[i]-48,5);

      if(i == a)res+=a; 
   }
   printf("%d\n",res);
}
