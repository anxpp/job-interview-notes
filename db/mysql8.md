
# 窗口函数

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

#### 511.游戏玩法分析 I

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

#### 512.游戏玩法分析 II

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

#### 534. 游戏玩法分析 III

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