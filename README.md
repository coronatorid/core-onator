# Coronator

Coronator is corona detection application that aim perfect contact tracing between it's user anonymously.

## Background

### Problem

The rise of coronavirus cases in Indonesia and the lack of awareness of ourselves which is we are contagious or not.

### Solution

Create the perfect contact tracing app that lets users know if they've been exposed to people who have been confirmed positive in the last 7 days anonymously

## Business Flow Diagram

![business-flow-diagram](https://user-images.githubusercontent.com/20650401/97368363-54887480-18dd-11eb-9a4c-afa1dd58e563.jpg)

## Architecture Diagram

![architecture-diagram](https://user-images.githubusercontent.com/20650401/97368364-55b9a180-18dd-11eb-9ab2-267edfb7d848.jpg)

## Database Diagram

![coronator](https://user-images.githubusercontent.com/20650401/97368361-53efde00-18dd-11eb-84ed-af5b8cddac57.png)

```
Table users {
  id int [pk, increment] // auto-increment
  phone varchar(12) [not null]
  active boolean [default: 1, not null]
  created_at datetime [not null]
  updated_at datetime [not null]

  indexes {
    phone [unique]
    (phone,active) [name:'phone_active']
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

Table confirmed_cases {
  ud int [pk, increment]
  user_id int  [not null]
  status int  [not null, note:'1 -> positive, 2 -> suspek, 3 -> probable, 4 -> kontak erat']
  created_at datetime  [not null]
  updated_at datetime  [not null]
}
```
