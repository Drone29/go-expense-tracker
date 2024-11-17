# Expense Tracker

## Description

Allows users to add, delete, and view their expenses as well as view a summary of the expenses

## Usage

### Add Expense
```sh
./expense-tracker add -description "some descr" -amount 10 -category "category"
# Output: Expense added successfully (ID: 1)
```
### Update Expense
```sh
./expense-tracker update 1 -description "New description"
# Output: Expense updated successfully (ID: 1)
./expense-tracker update 1 -category "New category"
# Output: Expense updated successfully (ID: 1)
```
### Delete Expense
```sh
./expense-tracker delete 1
# Output: Expense deleted successfully (ID: 1)
```

### List Expenses
#### List All
```sh
./expense-tracker list
# Output:
# ID         Date         Description          Amount     Category            
# 1          2024-11-17   Lunch                10                             
# 2          2024-11-17   Dinner               10                             
# 3          2024-11-17   Train                5                              
# 4          2024-11-17   some descr           10                             
# 5          2024-11-17   some descr           10         cat                 
# 6          2024-11-17   some descr2          1          cat                 
# 7          2024-11-17   some descr3          12         cat                 
# 8          2024-11-17   some descr4          3          cat                 
# 9          2024-11-17   some descr5          3          cat2                
# 10         2024-11-17   some descr6          2          cat2                
# 11         2024-11-17   some descr7          1          cat2 
```
#### List With Filter By Month/Category
```sh
./expense-tracker list -month 10
# Output:
# ID         Date         Description          Amount     Category            
# 1          2024-10-17   Lunch                10                             
# 2          2024-10-17   Dinner               10                             
# 3          2024-10-17   Train                5    
```
```sh
./expense-tracker list -category "cat2"
# Output:
# 9          2024-11-17   some descr5          3          cat2                
# 10         2024-11-17   some descr6          2          cat2                
# 11         2024-11-17   some descr7          1          cat2 
```

### Export To CSV
```sh
./expense-tracker export-csv
# Output:
# Exported expenses  2024  successfully to expenses-2024.csv
```

## Build And Install

To build and install, use `go build` and `go install` respectively, from the project's root directory 
```sh
go build
```
```sh
go install
```

## Testing

```sh
go test ./...
```
# Roadmap link
https://roadmap.sh/projects/expense-tracker
