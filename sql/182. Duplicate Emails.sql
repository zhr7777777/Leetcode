# Write your MySQL query statement below
select Email from (
    select Email, count(Email) as Number from Person group by Email
) as Temp where Number > 1

# error: Every derived table must have its own alias, 临时表要用as命名