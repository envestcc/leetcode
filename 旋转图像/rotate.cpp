/*
旋转图像
给定一个 n × n 的二维矩阵表示一个图像。

将图像顺时针旋转 90 度。

说明：

你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。

示例 1:

给定 matrix = 
[
  [1,2,3],
  [4,5,6],
  [7,8,9]
],

原地旋转输入矩阵，使其变为:
[
  [7,4,1],
  [8,5,2],
  [9,6,3]
]
示例 2:

给定 matrix =
[
  [ 5, 1, 9,11],
  [ 2, 4, 8,10],
  [13, 3, 6, 7],
  [15,14,12,16]
], 

原地旋转输入矩阵，使其变为:
[
  [15,13, 2, 5],
  [14, 3, 4, 1],
  [12, 6, 8, 9],
  [16, 7,10,11]
]
*/

#include<vector>

using namespace std;

void rotate(vector<vector<int>>& matrix) 
{
  for(int i=0;i<matrix.size();)
  {
    int n = matrix.size();
    if (n <= 1) return;
    
    for(int k=0;k<n-1-i;k++)
    {
      for(int j=0,x=i/2,y=i/2+k,x1=0,y1=0,last=matrix[x][y],tmp=0;j<4;j++)
      {
        x1 = y;
        y1 = n-1-x;
        tmp = matrix[x1][y1];
        matrix[x1][y1] = last;
        last = tmp;
        x = x1;
        y = y1;
      }
    }
    
    i+=2;
  }      
}

int main()
{
  return 0;
}