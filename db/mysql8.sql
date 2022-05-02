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

