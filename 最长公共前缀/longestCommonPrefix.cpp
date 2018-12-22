/*
最长公共前缀
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
说明:

所有输入只包含小写字母 a-z 。
*/

#include<string>
#include<vector>

using namespace std;

string longestCommonPrefix(string str1, string str2) 
{
    string common="";
    for(int i=0;i<str1.length() && i<str2.length();i++)
    {
        if (str1[i]==str2[i]) common.push_back(str1[i]);
        else break;
    }
    return common;
}

string longestCommonPrefix(vector<string>& strs) 
{
    int size = strs.size();
    if (size == 0) return "";
    if (size == 1) return strs[0];
    string common = longestCommonPrefix(strs[0], strs[1]);
    for(int i=2;i<size;i++)
    {
        if (common == "") return common;
        common = longestCommonPrefix(common, strs[i]);
    }
    return common;
}