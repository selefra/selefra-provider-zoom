# Table: zoom_role

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| account_id | string | X | √ | Zoom account ID. | 
| id | string | X | √ | Role ID. | 
| name | string | X | √ | Role name. | 
| description | string | X | √ | Role description. | 
| total_members | int | X | √ | Total number of members in the role. | 
| privileges | json | X | √ | Privileges assigned to the role. | 
| sub_account_privileges | json | X | √ | Privileges for management of sub-accounts. | 


