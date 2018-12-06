/****************
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

示例:

输入: [0,1,0,3,12]
输出: [1,3,12,0,0]
说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。
*******************/

#include <vector>

using namespace std;

void moveZeroes(vector<int>& nums) 
{
    int zeros=0;
    for(int i=0;i<nums.size();i++)
    {
        if (nums[i] == 0)
        {
            zeros ++;
        }else if (zeros)
        {
            nums[i-zeros] = nums[i];
        }
    }
    for(int i=0;i<zeros;i++)
    {
        nums[nums.size()-1-i] = 0;
    }
}