/*************************************************************

给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储一个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

示例 1:

输入: [1,2,3]
输出: [1,2,4]
解释: 输入数组表示数字 123。
示例 2:

输入: [4,3,2,1]
输出: [4,3,2,2]
解释: 输入数组表示数字 4321。

*************************************************************/

#include <vector>

using namespace std;


vector<int> plusOne(vector<int>& digits) {
    vector<int> temp(digits);
    temp.resize(digits.size()+1);
    int add=1;
    for(int i=digits.size()-1;i>=0;i--)
    {
        temp[i+1] = digits[i] + add;
        if (temp[i+1] >= 10)
        {
            add = 1;
            temp[i+1] -= 10;
        }else
        {
            add = 0;
        }
    }
    if (add >= 1)
    {
        temp[0] = 1;
        return temp;
    }else
    {
        vector<int> result;
        result.resize(digits.size());
        for(int i=0;i<result.size();i++)
        {
            result[i] = temp[i+1];
        }
        return result;
    }
}


int main(int argc, char const *argv[])
{
    vector<int> test1;
    plusOne(test1);
    return 0;
}
