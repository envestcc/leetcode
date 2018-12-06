/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


上图是一个部分填充的有效的数独。

数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true
示例 2:

输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
说明:

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。
*/

#include<vector>

using namespace std;

bool isValidSudoku(vector<vector<char>>& board) 
{
    char flag[10]={0};
    char empty=0, full=1;
    for(int i=0;i<board.size();i++)
    {
        memset(flag, empty, 10);
        for(int j=0, index=0;j<board.size();j++)
        {
            if (board[i][j] == '.')
            {
                index = 0;
            }else
            {
                index = board[i][j] - '0';
            }
            if (index && flag[index] != empty)
                return false;
            flag[index] = full;
        }
    }
    for(int i=0;i<board.size();i++)
    {
        memset(flag, empty, 10);
        for(int j=0, index=0;j<board.size();j++)
        {
            if (board[j][i] == '.')
            {
                index = 0;
            }else
            {
                index = board[j][i] - '0';
            }
            if (index && flag[index] != empty)
                return false;
            flag[index] = full;
        }
    }
    for(int i=0;i<9;i++)
    {
        int x_start = i%3*3,
            y_start = i/3*3;
        memset(flag, empty, 10);
        for(int j=0;j<3;j++)
        {
            for(int k=0,index=0;k<3;k++)
            {
                char value = board[x_start+j][y_start+k];
                if (value == '.')
                {
                    index = 0;
                }else
                {
                    index = value - '0';
                }
                if (index && flag[index]!=empty)
                    return false;
                flag[index] = full;
            }
        }
    }

    return true;
        
}