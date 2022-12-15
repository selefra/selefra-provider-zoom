# Table: zoom_user

## Columns 

|  Column Name   |  Data Type  | Uniq | Nullable | Description | 
|  ----  | ----  | ----  | ----  | ---- | 
| cms_user_id | string | X | √ | CMS ID of user, only enabled for Kaltura integration. | 
| group_ids | json | X | √ | IDs of groups where the user is a member. | 
| plan_united_type | string | X | √ | This field is returned if the user is enrolled in the Zoom United plan. | 
| personal_meeting_url | string | X | √ | User's personal meeting url. | 
| pmi | big_int | X | √ | Personal meeting ID of the user. | 
| vanity_url | string | X | √ | Personal meeting room URL, if the user has one. | 
| status | string | X | √ | User's status: pending, active or inactive. | 
| created_at | timestamp | X | √ | The time when user's account was created. | 
| jid | string | X | √ | User's JID. | 
| job_title | string | X | √ | User's job title. | 
| login_type | int | X | √ | Login type. 0 - Facebook, 1 - Google, 99 - API, 100 - ZOOM, 101 - SSO | 
| host_key | string | X | √ | The host key of the user. | 
| last_client_version | string | X | √ | The last client version that user used to login. | 
| phone_numbers | json | X | √ | User phone number, including verification status. | 
| language | string | X | √ | Default language for the Zoom Web Portal. | 
| pic_url | string | X | √ | The URL for user's profile picture. | 
| type | int | X | √ | User's plan type: 1 - Basic, 2 - Licensed or 3 - On-prem. | 
| id | string | X | √ | User ID. | 
| last_name | string | X | √ | User's last name. | 
| company | string | X | √ | User's company. | 
| custom_attributes | json | X | √ | Custom attributes, if any are assigned. | 
| dept | string | X | √ | Department, if provided by the user. | 
| last_login_time | timestamp | X | √ | User's last login time. There is a three-days buffer period for this field. For example, if user first logged in on 2020-01-01 and then logged out and logged in on 2020-01-02, the value of this field will still reflect the login time of 2020-01-01. However, if the user logs in on 2020-01-04, the value of this field will reflect the corresponding login time since it exceeds the three-day buffer period. | 
| role_id | string | X | √ | Unique identifier of the role assigned to the user. | 
| verified | int | X | √ | Display whether the user's email address for the Zoom account is verified or not. 1 - Verified user email. 0 - User's email not verified. | 
| account_id | string | X | √ | Zoom account ID. | 
| first_name | string | X | √ | User's first name. | 
| email | string | X | √ | User's email address. | 
| location | string | X | √ | User's location. | 
| use_pmi | bool | X | √ | Use Personal Meeting ID for instant meetings. | 
| timezone | string | X | √ | The time zone of the user. | 
| im_group_ids | json | X | √ | IDs of IM directory groups where the user is a member. | 


