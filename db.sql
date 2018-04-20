-- 176 second highest salary
SELECT IFNULL( (SELECT e.Salary FROM Employee e GROUP BY e.Salary ORDER BY e.Salary DESC LIMIT 1, 1), NULL) SecondHighestSalary;
SELECT MAX(Salary) as SecondHighestSalary FROM Employee WHERE Salary NOT IN (SELECT MAX(Salary) FROM Employee);

-- 627 swap salary
UPDATE salary SET  sex = (CASE WHEN sex='f' THEN 'm' ELSE 'f' END)
UPDATE salary SET sex = IF(sex='m','f','m')

-- 182 duplicate emails
SELECT Email FROM Person GROUP BY Email having COUNT(*) >1