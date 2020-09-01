# Write your MySQL query statement below
select w2.id from Weather w1, Weather w2 where datediff(w2.recordDate, w1.recordDate) = 1 and w2.Temperature > w1.Temperature;

# 使用datediff函数比较Date类型的数据， eg：datediff(w2.recordDate, w1.recordDate) = 1