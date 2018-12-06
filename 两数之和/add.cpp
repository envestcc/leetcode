/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

#include<vector>
#include<map>

using namespace std;

vector<int> twoSum(vector<int>& nums, int target)
{
    vector<int> res;
    int n = nums.size();
    map<int ,vector<int>> valueToIndex;
    for(int i=0, value=0;i<n;i++)
    {
        value = nums[i];
        if (valueToIndex.count(value)<1)
        {
            vector<int> tmp;
            tmp.push_back(i);
            valueToIndex[value] = tmp;
        }else
        {
            valueToIndex[value].push_back(i);
        }
    }
    for(int i=0, value=0;i<n;i++)
    {
        value = target - nums[i];
        if (valueToIndex.count(value) > 0)
        {
            vector<int> *tmp=&(valueToIndex[value]);
            for(int j=0;j<tmp->size();j++)
            {
                if (tmp->at(j) != i)
                {
                    res.push_back(i);
                    res.push_back(tmp->at(j));
                    return res;
                }
            }
        }
    }
    return res;
}