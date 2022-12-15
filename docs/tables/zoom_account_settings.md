# Table: zoom_account_settings

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| email_notification | json | X | √ | Email notification settings. | 
| security | json | X | √ | Security settings. | 
| integration | json | X | √ | Integration settings. | 
| trusted_domains | json | X | √ | Associated domains allow all users with that email domain to be prompted to join the account. | 
| id | string | X | √ | Account ID. Set to 'me' for the master account. | 
| tsp | json | X | √ | TSP settings. | 
| managed_domains | json | X | √ | Associated domains allow all users with that email domain to be prompted to join the account. | 
| in_meeting | json | X | √ | In meeting settings. | 
| meeting_authentication | json | X | √ | Meeting authentication options applied to the account. | 
| account_id | string | X | √ | Zoom account ID. | 
| schedule_meeting | json | X | √ | Schedule meeting settings. | 
| telephony | json | X | √ | Telephony settings. | 
| feature | json | X | √ | Feature settings. | 
| recording_authentication | json | X | √ | Recording authentication options applied to the account. | 
| meeting_security | json | X | √ | Meeting security settings applied to the account. | 
| recording | json | X | √ | Recording settings. | 


