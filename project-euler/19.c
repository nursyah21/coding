#include <stdio.h>

static int days(int y, int m);

int main(){
   int year, wday,sun,mon;
   wday=sun=0;
   
   for(int y=1901; y < 2001; y++)
      for(int m=1; m<13; m++){
         //1901 01 01 is tuesday, so wday == 5 means sunday
         if (wday==5)sun++;
         wday = (wday+days(y,m))%7;
      }
      

//   int res =days(1904,2);
   printf("weekday: %d\n",sun);
   return 0;
}


int days(int y, int m){
   if(m == 2)
      if((!(y%4) && y%100) || !(y%400)) return 29;
      else return 28;

   if(m == 4 || m == 6 || m == 9 || m == 11) return 30;

   return 31;
}

