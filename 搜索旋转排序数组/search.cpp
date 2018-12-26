/*
搜索旋转排序数组
假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/

#include<vector>

using namespace std;

int search(vector<int>& nums, int target) {
  //find sec start
  if (nums.size() <= 0) return -1;
  if (nums.size() == 1) {
    if (nums[0] == target) return 0;
    else return -1;
  }
  int secStart=0, secEnd=nums.size()-1;
  int mid=0;
  while(secStart<=secEnd)
  {
    mid = (secStart + secEnd) / 2;
    if (nums[mid] >= nums[0]) secStart = mid+1;
    else secEnd = mid-1;
  }
  if (secStart >= nums.size()) mid = nums.size();
  else if (secStart == secEnd) mid = secStart;
  else if (nums[secStart] > nums[secEnd]) mid = secEnd;
  else mid = secStart;

  // find from [0, mid-1], [mid, size-1] two sorted list
  int find_start, find_end;
  if (target < nums[0]) {
    find_start = mid;
    find_end = nums.size()-1;
  }else {
    find_start = 0;
    find_end = mid-1;
  }
  while(find_start<=find_end)
  {
    int mm = (find_start+find_end)/2;
    if (nums[mm] == target) {
      return mm;
    }else if (nums[mm] > target) {
      find_end = mm-1;
    }else {
      find_start = mm+1;
    }
  }
  return -1;
}