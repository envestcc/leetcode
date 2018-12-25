/*

字符串中的第一个唯一字符
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.
 

注意事项：您可以假定该字符串只包含小写字母。

*/

#include<string>
#include<iostream>
using namespace std;

static int xx = []() {
    ios::sync_with_stdio(false);
    cin.tie(NULL);
    return 0;
}();


int firstUniqChar(string s)
{
    int hash[26]={0};
    int res = -1;
    int len = s.size();
    for(int i=0;i<len;i++)
    {
        if (hash[s[i]-'a'] == 1)
        {
            hash[s[i]-'a'] = 2;
        }else if(hash[s[i]-'a'] == 0)
        {
            hash[s[i]-'a'] = 1;
        }
    }
    for(int i=0;i<len;i++)
    {
        if (hash[s[i]-'a']==1)
        {
            res = i;
            break;
        }
        
    }
    return res;
}
