# Golang Echo Showcase

This application is a small Golang showcase using SQLite and sqlc (with future Docker support for running the software anywhere).

## Data Sources

### 1. User Database
- A SQLite database storing user information.
- Stores first name and last name of a user.

### 2. KPIs Map
- A map storing KPIs in the format: `name string: kpi int`.
- Each KPI can have a type:
  - **`count`**:  
    The initial integer value is added to the current hour number each full hour, then divided by the total number of hours the entry has existed.  
    *Example*:  
    If you create a `count` KPI at 12:00 with a value of `7`, by 15:00 the value becomes:  
    `(7 + 13 + 14 + 15) / 3`
  - **`value`**:  
    The stored integer remains unchanged without any calculations.


## Features

- Perform standard CRUD operations on the User Database.
- All available endpoints are listed in `routes.go`.

## Code explonation

- In this code, I create a small example in Go using Echo and a local SQLite database.
- I follow a modular structure similar to Nest.js. This has the advantage that multiple people can work on different endpoints without overwriting each other’s work.


## Start Code
you can start the code with  `docker compose up` (docker must be installed).