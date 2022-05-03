
# 窗口函数



## 概述

    窗口函数从mysql8开始支持

窗口函数是将数据按指定字段分成多份（类似group by），对每个窗口执行对应的分析

静态窗口的大小是固定的，滑动窗口的大小根据设置变化

语法：

```
函数名([参数]) over(partition by [分组字段] order by [排序字段] asc/desc rows/range between 起始位置 and 结束位置)
```

- 函数：特定的窗口函数或聚合函数（所有聚合函数都可以），根据函数性质决定是否需要参数
- over：其中有三个参数，都是非必须
  - partition by：分组字段，分割数据的依据，不填的话整个数据当作一个窗口
  - order by：每个窗口的排序
  - rows/range between：针对滑动窗口的字段，当前窗口下分成更小的窗口
    - current row 边界是当前行
    - unbounded preceding 边界是分区中的第一行
    - unbounded following 边界是分区中的最后一行
    - expr preceding 边界是当前行减去expr的值
    - expr following 边界是当前行加上expr的值。rows是基于行数，range是基于值的大小

## 支持的窗口函数

| 函数           | 描述                                                                           |
|--------------|------------------------------------------------------------------------------|
| RANK         | 与DENSE_RANK（）函数类似，只是当两行或更多行具有相同的排名时，排序值序列中存在间隙。                              |
| DENSE_RANK   | 根据ORDER BY子句为其分区中的每一行分配一个排名。 它为具有相同值的行分配相同的排名。 如果两行或更多行具有相同的等级，则排序值序列中将没有间隙。 |
| ROW_NUMBER   | 为其分区中的每一行分配一个连续整数                                                            |
| CUME_DIST    | 计算一组值中值的累积分布。                                                                |
| FIRST_VALUE  | 返回指定表达式相对于窗口框架中第一行的值。                                                        |
| LAG          | 返回分区中当前行之前的第N行的值。 如果不存在前一行，则返回NULL。                                          |
| LAST_VALUE   | 返回指定表达式相对于窗口框架中最后一行的值。                                                       |
| LEAD         | 返回分区中当前行之后的第N行的值。 如果不存在后续行，则返回NULL。                                          |
| NTH_VALUE    | 返回窗口框架第N行的参数值                                                                |
| NTILE        | 将每个窗口分区的行分配到指定数量的已排名组中。                                                      |
| PERCENT_RANK | 计算分区或结果集中行的百分位数                                                              |

## 示例

```
create user local_test
    identified by '123456'
    with max_connections_per_hour 128 max_user_connections 64 max_queries_per_hour 65536;

create schema local_test collate utf8mb4_bin;

use local_test;

# 创建员工表
create table employee
(
    id            int primary key auto_increment,
    department_id int,
    salary        int
);

# 插入数据
truncate employee;
insert into employee(department_id, salary)
values (1, 8600),
       (1, 9500),
       (2, 8900),
       (1, 8300),
       (2, 8600),
       (3, 9200),
       (3, 8600),
       (1, 8600);

# 每个部门工资排名
select *, rank() over (partition by department_id order by salary desc) as ranking
from employee;

# 对比rank()、dense_rank()、row_number()
select *,
       # rank返回排名，重复项排名相同但会暂用名词
       rank() over (partition by department_id order by salary desc)       as ranking,
       # dense_rank返回排名，重复项排名相同且不暂用名次
       dense_rank() over (partition by department_id order by salary desc,id) dense_ranking,
       # row_number返回排序，重复项名称不同，按获取顺序得到排名
       row_number() over (partition by department_id order by salary desc) as row_num
from employee;

# 结合聚合函数
select *,
       sum(salary) over (order by id) as cur_sum,
       avg(salary) over (order by id) as cur_avg,
       count(1111) over (order by id) as cur_cnt,
       max(salary) over (order by id) as cur_max,
       min(salary) over (order by id) as cur_min
from employee;

# Top N 问题
## 取所有员工中薪资排名前四的
select t1.*, t2.ranking
from employee t1
         join (select *, rank() over (order by salary desc) as ranking
               from employee) t2 on t1.id = t2.id
where t2.ranking <= 4;

## 每个部分取薪资排名前2的
select t1.*, t2.ranking
from employee t1
         join (select *, rank() over (partition by department_id order by salary desc) as ranking
               from employee) t2 on t1.id = t2.id
where t2.ranking <= 2;

# 每个部门的平均工资
select distinct department_id, avg(salary) over (partition by department_id)
from employee;

select department_id, avg(salary)
from employee
group by department_id;
```

## leetcode 相关热门题目

#### 176. 第二高的薪水[中等]

```
Employee 表：
+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
id 是这个表的主键。
表的每一行包含员工的工资信息。

编写一个 SQL 查询，获取并返回 Employee 表中第二高的薪水 。如果不存在第二高的薪水，查询应该返回 null 。

查询结果如下例所示。

示例 1：

输入：
Employee 表：
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
输出：
+---------------------+
| SecondHighestSalary |
+---------------------+
| 200                 |
+---------------------+
示例 2：

输入：
Employee 表：
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+
输出：
+---------------------+
| SecondHighestSalary |
+---------------------+
| null                |
+---------------------+
```

常规解法：

```
select (
    select distinct salary from Employee order by salary desc limit 1 offset 1
) as SecondHighestSalary
```

窗口函数：
```
select(
    select distinct salary from
    (
        select salary,dense_rank() over(order by salary desc) as rk from Employee
    ) t where rk = 2
) SecondHighestSalary
```

#### 180. 连续出现的数字[中等]

```
表：Logs

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| num         | varchar |
+-------------+---------+
id 是这个表的主键。

编写一个 SQL 查询，查找所有至少连续出现三次的数字。

返回的结果表中的数据可以按 任意顺序 排列。

查询结果格式如下面的例子所示：

示例 1:

输入：
Logs 表：
+----+-----+
| Id | Num |
+----+-----+
| 1  | 1   |
| 2  | 1   |
| 3  | 1   |
| 4  | 2   |
| 5  | 1   |
| 6  | 2   |
| 7  | 2   |
+----+-----+
输出：
Result 表：
+-----------------+
| ConsecutiveNums |
+-----------------+
| 1               |
+-----------------+
解释：1 是唯一连续出现至少三次的数字。
```

常规解法：
```
SELECT DISTINCT Num AS ConsecutiveNums FROM Logs 
WHERE (Id+1, Num) IN (SELECT * FROM Logs)
AND (Id+2, Num) IN (SELECT * FROM Logs)
```

窗口函数：

```
select distinct t.num ConsecutiveNums from(
    select num,
        lead(num,1) over() l1,
        lead(num,2) over() l2
    from logs
) t where t.num = t.l1 and t.num=t.l2
```

#### 178. 分数排名[中等]

```
表: Scores

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| score       | decimal |
+-------------+---------+
Id是该表的主键。
该表的每一行都包含了一场比赛的分数。Score是一个有两位小数点的浮点值。

编写 SQL 查询对分数进行排序。排名按以下规则计算:

分数应按从高到低排列。
如果两个分数相等，那么两个分数的排名应该相同。
在排名相同的分数后，排名数应该是下一个连续的整数。换句话说，排名之间不应该有空缺的数字。
按 score 降序返回结果表。

查询结果格式如下所示。

示例 1:

输入: 
Scores 表:
+----+-------+
| id | score |
+----+-------+
| 1  | 3.50  |
| 2  | 3.65  |
| 3  | 4.00  |
| 4  | 3.85  |
| 5  | 4.00  |
| 6  | 3.65  |
+----+-------+
输出: 
+-------+------+
| score | rank |
+-------+------+
| 4.00  | 1    |
| 4.00  | 1    |
| 3.85  | 2    |
| 3.65  | 3    |
| 3.65  | 3    |
| 3.50  | 4    |
+-------+------+
```

常规解法：
```
select s.score,a.rank from scores s left join (
select score,cast(@rownum:=@rownum+1 as unsigned) as 'rank' from
(select distinct score from scores order by score desc) t,(SELECT @rownum:=0) a
) a on s.score=a.score order by s.score desc
```

窗口函数：

```
select score, dense_rank() over(order by score desc) `rank`
from scores;
```

#### 185. 部门工资前三高的所有员工[困难]

```
表: Employee

+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| id           | int     |
| name         | varchar |
| salary       | int     |
| departmentId | int     |
+--------------+---------+
Id是该表的主键列。
departmentId是Department表中ID的外键。
该表的每一行都表示员工的ID、姓名和工资。它还包含了他们部门的ID。

表: Department

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| name        | varchar |
+-------------+---------+
Id是该表的主键列。
该表的每一行表示部门ID和部门名。

公司的主管们感兴趣的是公司每个部门中谁赚的钱最多。一个部门的 高收入者 是指一个员工的工资在该部门的 不同 工资中 排名前三 。

编写一个SQL查询，找出每个部门中 收入高的员工 。

以 任意顺序 返回结果表。

查询结果格式如下所示。

示例 1:

输入: 
Employee 表:
+----+-------+--------+--------------+
| id | name  | salary | departmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 85000  | 1            |
| 2  | Henry | 80000  | 2            |
| 3  | Sam   | 60000  | 2            |
| 4  | Max   | 90000  | 1            |
| 5  | Janet | 69000  | 1            |
| 6  | Randy | 85000  | 1            |
| 7  | Will  | 70000  | 1            |
+----+-------+--------+--------------+
Department  表:
+----+-------+
| id | name  |
+----+-------+
| 1  | IT    |
| 2  | Sales |
+----+-------+
输出: 
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Max      | 90000  |
| IT         | Joe      | 85000  |
| IT         | Randy    | 85000  |
| IT         | Will     | 70000  |
| Sales      | Henry    | 80000  |
| Sales      | Sam      | 60000  |
+------------+----------+--------+
解释:
在IT部门:
- Max的工资最高
- 兰迪和乔都赚取第二高的独特的薪水
- 威尔的薪水是第三高的

在销售部:
- 亨利的工资最高
- 山姆的薪水第二高
- 没有第三高的工资，因为只有两名员工
```

常规解法：
```
select d.name Department,e.name Employee, e.Salary
from Employee e left join Department d on e.departmentId = d.id
where e.id in
(
    select t1.id
    from Employee t1 left join Employee t2
    on t1.departmentId = t2.departmentId and t1.salary<t2.salary
    group by t1.id
    having count(distinct t2.salary) <= 2
)
```

窗口函数：

```
select t2.name Department,t1.name Employee,salary from 
(
    select departmentId,name,salary,dense_rank() over(partition by departmentId order by salary desc) rk
    from employee
) t1 join department t2 on t1.departmentId = t2.id and t1.rk<=3
```

#### 197. 上升的温度[简单]

```
表： Weather

+---------------+---------+
| Column Name   | Type    |
+---------------+---------+
| id            | int     |
| recordDate    | date    |
| temperature   | int     |
+---------------+---------+
id 是这个表的主键
该表包含特定日期的温度信息

编写一个 SQL 查询，来查找与之前（昨天的）日期相比温度更高的所有日期的 id 。

返回结果 不要求顺序 。

查询结果格式如下例。

示例 1：

输入：
Weather 表：
+----+------------+-------------+
| id | recordDate | Temperature |
+----+------------+-------------+
| 1  | 2015-01-01 | 10          |
| 2  | 2015-01-02 | 25          |
| 3  | 2015-01-03 | 20          |
| 4  | 2015-01-04 | 30          |
+----+------------+-------------+
输出：
+----+
| id |
+----+
| 2  |
| 4  |
+----+
解释：
2015-01-02 的温度比前一天高（10 -> 25）
2015-01-04 的温度比前一天高（20 -> 30）
```

常规解法：
```
select t1.id 
from Weather t1 join Weather t2 on dateDiff(t1.recordDate,t2.recordDate)=1 and t1.temperature>t2.temperature
```

窗口函数：

```
select id from
(
    select id,
        temperature nowT,lag(temperature,1) over(order by recordDate) preT,
        recordDate nowD,lag(recordDate,1) over(order by recordDate) preD
    from Weather
) t where t.preT<t.nowT and dateDiff(t.nowD,t.preD)=1;
```

#### 177. 第N高的薪水[中等]

```
表: Employee

+-------------+------+
| Column Name | Type |
+-------------+------+
| id          | int  |
| salary      | int  |
+-------------+------+
Id是该表的主键列。
该表的每一行都包含有关员工工资的信息。

编写一个SQL查询来报告 Employee 表中第 n 高的工资。如果没有第 n 个最高工资，查询应该报告为 null 。

查询结果格式如下所示。

示例 1:

输入: 
Employee table:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
| 2  | 200    |
| 3  | 300    |
+----+--------+
n = 2
输出: 
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| 200                    |
+------------------------+
示例 2:

输入: 
Employee 表:
+----+--------+
| id | salary |
+----+--------+
| 1  | 100    |
+----+--------+
n = 2
输出: 
+------------------------+
| getNthHighestSalary(2) |
+------------------------+
| null                   |
+------------------------+
```

常规解法：
```
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  SET N = N-1;
  RETURN (
      # Write your MySQL query statement below.
      SELECT DISTINCT salary FROM employee ORDER BY salary DESC LIMIT N, 1
  );
END
```

窗口函数：

```
CREATE FUNCTION getNthHighestSalary(N INT) RETURNS INT
BEGIN
  RETURN (
      # Write your MySQL query statement below.
        select distinct salary from
        (
            select salary, dense_rank() over(order by salary desc) rk
            from Employee
        ) t where rk=N
  );
END
```

#### 184. 部门工资最高的员工[中等]

```
表： Employee

+--------------+---------+
| 列名          | 类型    |
+--------------+---------+
| id           | int     |
| name         | varchar |
| salary       | int     |
| departmentId | int     |
+--------------+---------+
id是此表的主键列。
departmentId是Department表中ID的外键。
此表的每一行都表示员工的ID、姓名和工资。它还包含他们所在部门的ID。

表： Department

+-------------+---------+
| 列名         | 类型    |
+-------------+---------+
| id          | int     |
| name        | varchar |
+-------------+---------+
id是此表的主键列。
此表的每一行都表示一个部门的ID及其名称。

编写SQL查询以查找每个部门中薪资最高的员工。
按 任意顺序 返回结果表。
查询结果格式如下例所示。

示例 1:

输入：
Employee 表:
+----+-------+--------+--------------+
| id | name  | salary | departmentId |
+----+-------+--------+--------------+
| 1  | Joe   | 70000  | 1            |
| 2  | Jim   | 90000  | 1            |
| 3  | Henry | 80000  | 2            |
| 4  | Sam   | 60000  | 2            |
| 5  | Max   | 90000  | 1            |
+----+-------+--------+--------------+
Department 表:
+----+-------+
| id | name  |
+----+-------+
| 1  | IT    |
| 2  | Sales |
+----+-------+
输出：
+------------+----------+--------+
| Department | Employee | Salary |
+------------+----------+--------+
| IT         | Jim      | 90000  |
| Sales      | Henry    | 80000  |
| IT         | Max      | 90000  |
+------------+----------+--------+
解释：Max 和 Jim 在 IT 部门的工资都是最高的，Henry 在销售部的工资最高。
```

常规解法：
```
SELECT
    Department.name AS 'Department',
    Employee.name AS 'Employee',
    Salary
FROM
    Employee
        JOIN
    Department ON Employee.DepartmentId = Department.Id
WHERE
    (Employee.DepartmentId , Salary) IN
    (   SELECT
            DepartmentId, MAX(Salary)
        FROM
            Employee
        GROUP BY DepartmentId
	);
```

窗口函数：

```
select t2.name Department,t1.name Employee, t1.Salary
from
(
    select name,departmentId,salary,rank() over(partition by departmentId order by salary desc) rk 
    from employee
) t1 join Department t2 on t1.departmentId = t2.id and t1.rk = 1
```

#### 196. 删除重复的电子邮箱[简单]

```
表: Person

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| id          | int     |
| email       | varchar |
+-------------+---------+
id是该表的主键列。
该表的每一行包含一封电子邮件。电子邮件将不包含大写字母。
 

编写一个 SQL 删除语句来 删除 所有重复的电子邮件，只保留一个id最小的唯一电子邮件。

以 任意顺序 返回结果表。 （注意： 仅需要写删除语句，将自动对剩余结果进行查询）

查询结果格式如下所示。

 

 

示例 1:

输入: 
Person 表:
+----+------------------+
| id | email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
| 3  | john@example.com |
+----+------------------+
输出: 
+----+------------------+
| id | email            |
+----+------------------+
| 1  | john@example.com |
| 2  | bob@example.com  |
+----+------------------+
解释: john@example.com重复两次。我们保留最小的Id = 1。
```

常规解法：
```
DELETE from Person 
Where Id not in (
    Select Id 
    From(
    Select MIN(Id) as id
    From Person 
    Group by Email
   ) t
)
```

窗口函数：

```
delete t1 from Person t1 join 
(
    select id,row_number() over(partition by email order by id) rn from Person
) t2 on t1.id = t2.id
where t2.rn>1
```

#### 601. 体育馆的人流量[困难]

```
表：Stadium
+---------------+---------+
| Column Name   | Type    |
+---------------+---------+
| id            | int     |
| visit_date    | date    |
| people        | int     |
+---------------+---------+
visit_date 是表的主键
每日人流量信息被记录在这三列信息中：序号 (id)、日期 (visit_date)、 人流量 (people)
每天只有一行记录，日期随着 id 的增加而增加

编写一个 SQL 查询以找出每行的人数大于或等于 100 且 id 连续的三行或更多行记录。

返回按 visit_date 升序排列 的结果表。

查询结果格式如下所示。

示例 1:

输入：
Stadium 表:
+------+------------+-----------+
| id   | visit_date | people    |
+------+------------+-----------+
| 1    | 2017-01-01 | 10        |
| 2    | 2017-01-02 | 109       |
| 3    | 2017-01-03 | 150       |
| 4    | 2017-01-04 | 99        |
| 5    | 2017-01-05 | 145       |
| 6    | 2017-01-06 | 1455      |
| 7    | 2017-01-07 | 199       |
| 8    | 2017-01-09 | 188       |
+------+------------+-----------+
输出：
+------+------------+-----------+
| id   | visit_date | people    |
+------+------------+-----------+
| 5    | 2017-01-05 | 145       |
| 6    | 2017-01-06 | 1455      |
| 7    | 2017-01-07 | 199       |
| 8    | 2017-01-09 | 188       |
+------+------------+-----------+
解释：
id 为 5、6、7、8 的四行 id 连续，并且每行都有 >= 100 的人数记录。
请注意，即使第 7 行和第 8 行的 visit_date 不是连续的，输出也应当包含第 8 行，因为我们只需要考虑 id 连续的记录。
不输出 id 为 2 和 3 的行，因为至少需要三条 id 连续的记录。
```

常规解法：
```
select distinct a.* from stadium a,stadium b,stadium c
where a.people>=100 and b.people>=100 and c.people>=100
and (
     (a.id = b.id-1 and b.id = c.id -1) or
     (a.id = b.id-1 and a.id = c.id +1) or
     (a.id = b.id+1 and b.id = c.id +1)
) order by a.id
```

窗口函数：

```
select id,visit_date,people
from(
    select id,visit_date,people,
        lag(people,1) over(order by id) pre1,
        lag(people,2) over(order by id) pre2,
        lead(people,1) over(order by id) next1,
        lead(people,2) over(order by id) next2
    from stadium
) t
where people>=100 and (
    next1>=100 and next2>=100
    or
    pre1>=100 and pre2>=100
    or
    pre1>=100 and next1>=100
)
order by id
```

#### 182. 查找重复的电子邮箱[简单]

```
编写一个 SQL 查询，查找 Person 表中所有重复的电子邮箱。

示例：

+----+---------+
| Id | Email   |
+----+---------+
| 1  | a@b.com |
| 2  | c@d.com |
| 3  | a@b.com |
+----+---------+
根据以上输入，你的查询应返回以下结果：

+---------+
| Email   |
+---------+
| a@b.com |
+---------+
说明：所有电子邮箱都是小写字母。
```

常规解法：
```
select email from person group by email having count(1)>1
```

窗口函数：

```
select distinct email from(
    select email, count(1) over(partition by email) cnt from person
) t where t.cnt>1
```

#### 596. 超过5名学生的课[简单]

```
表: Courses

+-------------+---------+
| Column Name | Type    |
+-------------+---------+
| student     | varchar |
| class       | varchar |
+-------------+---------+
(student, class)是该表的主键列。
该表的每一行表示学生的名字和他们注册的班级。

编写一个SQL查询来报告 至少有5个学生 的所有类。

以 任意顺序 返回结果表。

查询结果格式如下所示。

示例 1:

输入: 
Courses table:
+---------+----------+
| student | class    |
+---------+----------+
| A       | Math     |
| B       | English  |
| C       | Math     |
| D       | Biology  |
| E       | Math     |
| F       | Computer |
| G       | Math     |
| H       | Math     |
| I       | Math     |
+---------+----------+
输出: 
+---------+ 
| class   | 
+---------+ 
| Math    | 
+---------+
```

常规解法：
```
select class from courses group by class having count(1)>=5;
```

窗口函数：

```
select distinct class from(
    select class,count(1) over(partition by class) cnt from courses
) t where cnt>=5;
```

#### 511.游戏玩法分析 I[简单]

活动表Activity：

```
+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| player_id    | int     |
| device_id    | int     |
| event_date   | date    |
| games_played | int     |
+--------------+---------+
表的主键是 (player_id, event_date)。
这张表展示了一些游戏玩家在游戏平台上的行为活动。
每行数据记录了一名玩家在退出平台之前，当天使用同一台设备登录平台后打开的游戏的数目（可能是 0 个）。
```

Q: 写一条 SQL 查询语句获取每位玩家 第一次登陆平台的日期。

```
Activity 表：
+-----------+-----------+------------+--------------+
| player_id | device_id | event_date | games_played |
+-----------+-----------+------------+--------------+
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 2         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |
+-----------+-----------+------------+--------------+

Result 表：
+-----------+-------------+
| player_id | first_login |
+-----------+-------------+
| 1         | 2016-03-01  |
| 2         | 2017-06-25  |
| 3         | 2016-03-02  |
+-----------+-------------+
```

A:

```
# 常规解法
select player_id,min(event_date) as first_login from Activity group by player_id
# 窗口函数
select distinct player_id, min(event_date) over(partition by player_id) as first_login from Activity;
```

#### 512.游戏玩法分析 II[简单]

```
Table: Activity

+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| player_id    | int     |
| device_id    | int     |
| event_date   | date    |
| games_played | int     |
+--------------+---------+
(player_id, event_date) 是这个表的两个主键
这个表显示的是某些游戏玩家的游戏活动情况
每一行是在某天使用某个设备登出之前登录并玩多个游戏（可能为0）的玩家的记录

请编写一个 SQL 查询，描述每一个玩家首次登陆的设备名称

查询结果格式在以下示例中：

Activity table:
+-----------+-----------+------------+--------------+
| player_id | device_id | event_date | games_played |
+-----------+-----------+------------+--------------+
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 2         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |
+-----------+-----------+------------+--------------+

Result table:
+-----------+-----------+
| player_id | device_id |
+-----------+-----------+
| 1         | 2         |
| 2         | 3         |
| 3         | 1         |
+-----------+-----------+
```

常规解法：

```
select t1.player_id, t1.device_id
from Activity t1 join
(
    select player_id,min(event_date) min_date from Activity group by player_id 
) t2 on t1.player_id = t2.player_id and t1.event_date = t2.min_date;
```

窗口函数：

```
select player_id,device_id 
from 
(   select player_id,device_id,row_number() over(partition by player_id order by event_date) as rk 
    from Activity
) t
where rk=1
```

#### 534. 游戏玩法分析 III[中等]

```
Table: Activity

+--------------+---------+
| Column Name  | Type    |
+--------------+---------+
| player_id    | int     |
| device_id    | int     |
| event_date   | date    |
| games_played | int     |
+--------------+---------+
（player_id，event_date）是此表的主键。
这张表显示了某些游戏的玩家的活动情况。
每一行是一个玩家的记录，他在某一天使用某个设备注销之前登录并玩了很多游戏（可能是 0 ）。

编写一个 SQL 查询，同时报告每组玩家和日期，以及玩家到目前为止玩了多少游戏。也就是说，在此日期之前玩家所玩的游戏总数。详细情况请查看示例。

查询结果格式如下所示：

Activity table:
+-----------+-----------+------------+--------------+
| player_id | device_id | event_date | games_played |
+-----------+-----------+------------+--------------+
| 1         | 2         | 2016-03-01 | 5            |
| 1         | 2         | 2016-05-02 | 6            |
| 1         | 3         | 2017-06-25 | 1            |
| 3         | 1         | 2016-03-02 | 0            |
| 3         | 4         | 2018-07-03 | 5            |
+-----------+-----------+------------+--------------+

Result table:
+-----------+------------+---------------------+
| player_id | event_date | games_played_so_far |
+-----------+------------+---------------------+
| 1         | 2016-03-01 | 5                   |
| 1         | 2016-05-02 | 11                  |
| 1         | 2017-06-25 | 12                  |
| 3         | 2016-03-02 | 0                   |
| 3         | 2018-07-03 | 5                   |
+-----------+------------+---------------------+
对于 ID 为 1 的玩家，2016-05-02 共玩了 5+6=11 个游戏，2017-06-25 共玩了 5+6+1=12 个游戏。
对于 ID 为 3 的玩家，2018-07-03 共玩了 0+5=5 个游戏。
请注意，对于每个玩家，我们只关心玩家的登录日期。
```

常规解法：

```
select player_id,event_date,
    (select sum(games_played) from Activity where event_date <= t.event_date and player_id = t.player_id) games_played_so_far
from Activity t
```

窗口函数：

```
select player_id,event_date,
    sum(games_played) over(partition by player_id order by event_date) games_played_so_far 
from Activity
```