# Write your MySQL query statement below
select distinct l1.Num as ConsecutiveNums 
    from 
        Logs l1, Logs l2, Logs l3 
    where 
        l2.id - 1 = l1.id and 
        l3.id - 1 = l2.id and
        l1.Num = l2.Num and 
        l2.Num = l3.Num;