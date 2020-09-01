# Write your MySQL query statement below
select Name as Customers 
    from Customers left join Orders 
    on Customers.Id = Orders.CustomerId 
    where Orders.CustomerId <=> null;

# 注意左链接，链接条件用on关键字