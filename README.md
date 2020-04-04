# 何志行毕设项目数据库设计字段
```
# 实验室管理系统
drop database if exists lab_system;
create database lab_system;
use lab_system;
drop table if exists users_table;
create table users_table(
    user_id int not null primary key auto_increment,  #1、用户ID
    user_name varchar(50) ,                           #2、用户名字
    user_account varchar(11) not null unique ,        #3、用户账户
    user_password varchar(11),                        #4、用户密码
    user_tel varchar(11) not null,                    #5、用户电话
    user_type int,                                    #6、用户类型 1超级管理员 2管理员 3教师
    user_context varchar(200)                         #7、备注信息
);

drop table if exists labs_table;
create table labs_table(
    lab_id int not null primary key auto_increment,   #1、实验室ID
    lab_name varchar(50),                             #2、实验室名字
    lab_address varchar(40),                          #3、实验室地址
    lab_capacity int,                                 #4、实验室容量
    lab_context varchar(200)                          #5、备注信息
);

drop table if exists courses_table;
create table courses_table(
    order_id int not null primary key auto_increment,   #1、记录ID
    lab_id int        ,                                 #2、实验室id
    lab_name varchar(50) ,                              #3、实验室名字
    lab_address varchar(40),                            #4、实验室地址
    lab_context varchar(200),                           #5、实验室描述

    date int,                                           #6、时间戳
    datestr varchar(20),                                #7、2020/4/4字符串时间
    week varchar(10),                                   #8、周几

    teach_id int,                                       #9、预约教师ID
    teach_name varchar(20),                             #10、货物描述
    teach_tel int(11),                                  #11、教师电话
    class_name varchar(40),                             #12、班级名字

    course_id int,                                      #12、课程 12、34、56、78表示12节课
    course_idstr varchar(10),                           #13、课程字符串 1、2节课
    course_context varchar(200)                         #14、课程信息
);
```
