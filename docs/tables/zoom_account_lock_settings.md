# Table: zoom_account_lock_settings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| schedule_meeting | json | X | √ | Schedule meeting lock settings. | 
| email_notification | json | X | √ | Email notification lock settings. | 
| meeting_security | json | X | √ | Meeting security lock settings. | 
| account_id | string | X | √ | Zoom account ID. | 
| id | string | X | √ | Account ID. Set to 'me' for the master account. | 
| in_meeting | json | X | √ | In meeting lock settings. | 
| recording | json | X | √ | Recording lock settings. | 
| telephony | json | X | √ | Telephony lock settings. | 
| tsp | json | X | √ | TSP lock settings. | 


