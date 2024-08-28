create user 'application'@'%' identified by 'app_passwd';
grant select, insert, update on v3.* to 'application'@'%';

create table if not exists devices (
    deviceID varchar(255) not null primary key
    );

create table if not exists gps (
    deviceID varchar(255) not null,
    latitude float(24) not null,
    longitude float(24) not null,
    time bigint not null,
    constraint ID primary key (deviceID, time),
    constraint gps_device foreign key (deviceID) references devices(deviceID)
    );

create table if not exists gyroscope (
    deviceID varchar(255) not null,
    x float(24) not null,
    y float(24) not null,
    z float(24) not null,
    time bigint not null,
    constraint ID primary key (deviceID, time),
    constraint gyroscope_device foreign key (deviceID) references devices(deviceID)
    );

create table if not exists photos (
    deviceID varchar(255) not null,
    photo varchar(255) not null,
    time bigint not null,
    constraint ID primary key (deviceID, time),
    constraint photo_device foreign key (deviceID) references devices(deviceID)
    );

insert into devices (deviceID) 
    values ('Device1'), ('Device2'), ('Device3');
