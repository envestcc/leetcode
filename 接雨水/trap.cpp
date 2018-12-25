/*
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。 感谢 Marcos 贡献此图。

示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6

*/

#include<vector>

using namespace std;

int trap(vector<int>& height) {
  int left=-1, right=-1, result=0, status=0;
  for(int i=1;i<height.size();i++) {
    if (status == 0) {//find left
      if (height[i] < height[i-1]) {
        left = i-1;
        status = 1;
      }
    }else if (status == 1) {//find right
      if (height[i] > height[i-1]) {
        right = i;
        status = 2;
      }
    }else if (status == 2) {//find heighest right
      if (height[i] >= height[i-1]) {
        right = i;
      }else {//get
        int h = height[left]>height[right] ? height[right] : height[left];
        for(int j=left+1;j<right;j++) {
          if (height[j] < h) result += h-height[j];
        }
        left = i-1;
        right = -1;
        status = 1;
      }
    }
  }
  if (status == 2 && left>=0 && right>=0){
    int h = height[left]>height[right] ? height[right] : height[left];
    for(int j=left+1;j<right;j++) {
      if (height[j] < h) result += h-height[j];
    }
  }
  return result;
}