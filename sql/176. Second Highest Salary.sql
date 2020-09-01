# Write your MySQL query statement below
SELECT
    (SELECT DISTINCT
            Salary
        FROM
            Employee
        ORDER BY Salary DESC
        LIMIT 1 OFFSET 1) AS SecondHighestSalary
;

SELECT
    IFNULL(
      (SELECT DISTINCT Salary
       FROM Employee
       ORDER BY Salary DESC
        LIMIT 1 OFFSET 1),
    NULL) AS SecondHighestSalary
;

# select max(Salary) as SecondHighestSalary from Employee where Salary < (select max(Salary) from Employee)

# 用order by排序，用distinct对结果去重，取第二条元组。考虑只有一条记录的情况，没有第二高工资，需要返回一条null的记录，不能返回空
# 临时表
# IFNULL关键字
# select NULL as SecondHighestSalary