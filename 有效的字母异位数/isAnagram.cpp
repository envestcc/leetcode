/*
有效的字母异位词
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的一个字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false
说明:
你可以假设字符串只包含小写字母。

进阶:
如果输入字符串包含 unicode 字符怎么办？你能否调整你的解法来应对这种情况？
*/

#include<string>
#include<map>

using namespace std;

bool isAnagram(string s, string t) 
{
    map<string, int> s_map, t_map;
    for(int i=0;i<s.size();i++)
    {
        string tmp;
        tmp.push_back(s[i]);
        if (s_map.count(tmp)==0)
        {
            s_map[tmp] = 1;
        }else 
        {
            s_map[tmp] ++;
        }
    }
    for(int i=0;i<t.size();i++)
    {
        string tmp;
        tmp.push_back(t[i]);
        if (t_map.count(tmp)==0)
        {
            t_map[tmp] = 1;
        }else 
        {
            t_map[tmp] ++;
        }
    }
    if (s_map.size() != t_map.size()) return false;
    map<string, int>::iterator iter, search;
    for(iter=s_map.begin();iter != s_map.end();iter++)
    {
        search = t_map.find(iter->first);
        if (search == t_map.end())
        {
            return false;
        }
        else if (search->second != iter->second)
        {
            return false;
        }
    }
    return true;
}