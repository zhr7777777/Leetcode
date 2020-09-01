select Department.Name as Department, Temp.Name as Employee, Salary 
from (
    select Name, Salary, DepartmentId, dense_rank() over (partition by DepartmentId order by Salary desc) as department_salary_rank
    from Employee
) as Temp, Department 
where DepartmentId = Department.Id and department_salary_rank <= 3;

# 重命名表不要用Rank
# partition by 用法