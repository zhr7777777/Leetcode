# Write your MySQL query statement below
select Score as score, dense_rank() over (order by Score desc) as 'Rank' from Scores

# dense_rank()的用法
# rank()是跳跃排序，有两个第一名时接下来是第三名，113
# dense_rank()是连续排序，有两个第一名时接下来是第二名，112
# row_number()，生成序号，123
# Rank 是保留字需要用引号