/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

*/

#include<vector>
#include<cstdio>

using namespace std;

int fixedLeftSide(int start, int end, vector<int>& height) {
  if (start >= end) return 0;
  int max=start+1;
  for(int i=start+2;i<=end;i++) {
    if (height[i] > height[max]) max=i;
  }
  int result=0;
  for(int i=start+1;i<max;i++) {
    result += height[max] - height[i];
  }
  result += fixedLeftSide(max, end, height);
  return result;
}

int fixedRightSide(int start, int end, vector<int>& height) {
  if (start >= end) return 0;
  int max=start;
  for(int i=start+1;i<end;i++) {
    if (height[i] > height[max]) max=i;
  }
  int result=0;
  for(int i=max+1;i<end;i++) {
    result += height[max] - height[i];
  }
  result += fixedRightSide(start, max, height);
  return result;
}


int trap(vector<int>& height) {
  int result=0;
  int max=0, i=0;
  while(i<height.size()) {
    if (height[i] > height[max]) max = i;
    i++;
  }
  result = fixedRightSide(0, max, height) +fixedLeftSide(max, height.size()-1, height);
  return result;
}

int main(int argc, char const *argv[])
{
  vector<int> height = {6,4,2,0,3,2,0,3,1,4,5,3,2,7,5,3,0,1,2,1,3,4,6,8,1,3};
  int result = trap(height);
  printf("%d\n", result);
  return 0;
}
