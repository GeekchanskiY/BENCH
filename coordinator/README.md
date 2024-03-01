

### Models
    users:
        id PK
        name
        email
        password
        is_staff
        ip_adress

        created_at
        updated_at
    
    services:
        id PK
        name
        internal_hostname
    
    service_url:
        id PK
        service FK => services.id
        url
        required_params
        requirements (is_auth, etc.)
    
    visits:
        id PK
        user_id FK => users.id
        service_id FK => services.id
        last_visited
        amount

    

