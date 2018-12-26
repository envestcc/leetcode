/*
三数之和
给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

例如, 给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

#include<vector>
#include<algorithm>

using namespace std;


vector<vector<int>> threeSum(vector<int>& nums) 
{
    vector<vector<int>> result;
    sort(nums.begin(), nums.end());
    int len = nums.size();
    for(int i=0,lasta=-1;i<len;i++)
    {
        if (nums[i] > 0) break;
        if (lasta>=0 && nums[i]==nums[lasta]) continue;
        int j=i+1,k=len-1,last=-1;
        while(j < k)
        {
            int sum = nums[i]+nums[j]+nums[k];
            if (sum == 0)
            {
                if (last>=0 && nums[j]==nums[last])
                {

                }else
                {
                    result.push_back({nums[i], nums[j], nums[k]});
                    last = j;
                    lasta = i;
                }
                k--;
                j++;
                
            }else if (sum > 0)
            {
                k--;
            }else if (sum < 0)
            {
                j++;
            }
        }
    }
    return result;
}