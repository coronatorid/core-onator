# Coronator

[![Coverage Status](https://coveralls.io/repos/github/coronatorid/core-onator/badge.svg?branch=master)](https://coveralls.io/github/coronatorid/core-onator?branch=master)

Coronator is corona detection application that aim perfect contact tracing between it's user anonymously.

[Presentation](https://docs.google.com/presentation/d/1B1wPEZKtG-sUSVK--z16QpEKHrt8e94ErVX_Xv3sHKI/edit#slide=id.ga50e610f2f_0_15).

## Background

### Problem

The rise of coronavirus cases in Indonesia and the lack of awareness of ourselves which is we are contagious or not.

### Solution

Create the perfect contact tracing app that lets users know if they've been exposed to people who have been confirmed positive in the last 7 days
anonymously

## How to Contribute

### Prerequisite

You need to have all of this to run core-onator

1. Go with a version of 1.14.6 or higher
2. MYSQL
3. Redis
4. Kafka & Zookeeper
5. Altair, used for API gateway and authentication service. You can install it by clone it in [here](https://github.com/coronatorid/altair)

### Installation

Here is step by step if you want to contribute to core-onator.

1. clone this repo `git clone git@github.com:coronatorid/core-onator.git`
2. Run the migration (make sure mysql is on and the schema is already created) `go run coronator.go db:migrate`
3. Don't forget to run Altair.
4. Run coronator `go run coronator.go run:api`


### Contribution

1. Make sure you follow the standard of golang coding style using `gofmt`
2. All PR will be reviewed by main contributor of this repo
3. PR which inactive more than 1 month will be closed. (You can reopen it)


### Feature Idea Contribution

This one seems like a good idea, we will have on in near future.

## Diagram

All diagram related to core-onator.

### Business Flow Diagram

![business-flow-diagram](https://user-images.githubusercontent.com/20650401/97368363-54887480-18dd-11eb-9a4c-afa1dd58e563.jpg)

### Architecture Diagram

![architecture-diagram](https://user-images.githubusercontent.com/20650401/103154120-136eec80-47c8-11eb-8935-a22e5a94ecfd.png)

### Database Diagram

![coronator](https://user-images.githubusercontent.com/20650401/103976861-78541a80-51aa-11eb-9c70-bd408e3f8b6e.png)

```
Table users {
  id int [pk, increment] // auto-increment
  phone varchar(255) [not null]
  state tinyint [not null,note: '1 -> active, 0 -> inactive']
  role tinyint [not null,note: 'null -> admin, 1 -> superadmin']
  created_at datetime [not null]
  updated_at datetime [not null]

  indexes {
    phone [unique]
    (phone,state) [name:'phone_state']
  }
}

Table locations {
  id int [pk, increment]
  user_id int [not null]
  lat double [not null]
  long double [not null]
  created_at datetime [not null]
  updated_at datetime [not null]

  indexes {
    (user_id,lat) [name:'user_id_lat']
    (user_id,long) [name:'user_id_long']
  }
}

Table reported_cases {
  id int [pk, increment]
  user_id int  [not null]
  status int  [not null, note:'1 -> confirmed, 0 -> not_confirmed']
  telegram_message_id varchar(255) [not null]
  telegram_image_url varchar(255) [not null]
  image_path varchar(255) [not null]
  image_deleted boolean
  created_at datetime  [not null]
  updated_at datetime  [not null]

  indexes {
    (user_id) [name:'user_id',unique]
  }

}

Table confirmed_cases {
  id int [pk, increment]
  user_id int  [not null]
  status int  [not null, note:'1 -> positive, 2 -> suspek, 3 -> probable, 4 -> kontak erat']
  created_at datetime  [not null]
  updated_at datetime  [not null]

  indexes {
    (user_id) [name:'user_id',unique]
  }
}

Table exposed_users {
  id int [pk, increment]
  user_id int  [not null]
  confirmed_cases_id int  [not null]
  lat double [not null]
  long double [not null]
  created_at datetime  [not null]
  updated_at datetime  [not null]

  indexes {
    (user_id) [name:'user_id']
  }
}

Ref: exposed_users.user_id > users.id
Ref: locations.user_id > users.id
Ref: users.id - confirmed_cases.user_id
Ref: confirmed_cases.id < exposed_users.confirmed_cases_id
Ref: users.id - reported_cases.user_id
```
