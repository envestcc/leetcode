/*
整数反转
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

示例 1:

输入: 123
输出: 321
 示例 2:

输入: -123
输出: -321
示例 3:

输入: 120
输出: 21
注意:

假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−2^31,  2^31 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。
*/

#include<vector>

using namespace std;

int reverse(int x) 
{
   vector<char> parts;
   int flag = x > 0 ? 1 : -1;
   int max=2147483647, min=-2147483648;   
   for(int i=x*flag;i>0;)
   {
      parts.push_back(i % 10);
      i = i / 10;
   }
   int res = 0;
   for(int i=0;i<parts.size();i++)
   {
      if ((max-parts[i])/10 < res)
      {
         return 0;
      }
      res = res * 10 + parts[i];
   }
   res *= flag;
   return res;
}