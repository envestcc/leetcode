/*
验证回文字符串
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
*/

#include<string>

using namespace std;

bool isPalindrome(string s) 
{
    string format;
    for(int i=0;i<s.length();i++)
    {
        if (s[i]>='a' && s[i]<='z')
        {
            format.push_back(s[i]);
        }else if (s[i]>='A' && s[i]<='Z')
        {
            format.push_back(s[i]-'A'+'a');
        }else if (s[i]>='0' && s[i]<='9')
        {
            format.push_back(s[i]);
        }
    }
    string revert;
    for(int i=format.length()-1;i>=0;i--)
    {
        revert.push_back(format[i]);
    }
    return format==revert;
}