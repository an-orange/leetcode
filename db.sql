-- 176
select IFNULL( (select e.Salary from Employee e group by e.Salary order by e.Salary desc limit 1, 1), NULL) SecondHighestSalary;
SELECT MAX(Salary) as SecondHighestSalary FROM Employee WHERE Salary NOT IN (SELECT MAX(Salary) FROM Employee);

-- 627
update salary set  sex = (case when sex='f' then 'm' else 'f' end)
update salary  set sex = IF(sex='m','f','m')

--182 duplicate emails
select Email from Person group by Email having count(*) >1