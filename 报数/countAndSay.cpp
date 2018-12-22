/*
报数
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

 

示例 1:

输入: 1
输出: "1"
示例 2:

输入: 4
输出: "1211"
*/

#include<string>

using namespace std;

string countAndSay(int n) 
{
    if (n == 1) return string("1");
    string last = countAndSay(n-1);
    char at=last[0];
    int num=1;
    string result;
    for(int i=1;i<last.length();i++)
    {
        if (last[i] == at){
            num++;
        }else{
            result.push_back(num+'0');
            result.push_back(at);
            at = last[i];
            num = 1;
        }
    }
    result.push_back(num+'0');
    result.push_back(at);
    return result;
}