## 1. Sprendžiamo uždavinio aprašymas

### **1.1. Sistemos paskirtis**

Projekto tikslas – sukurti informacinę sistemą, leidžiančią smulkioms įmonėms ar verslams (kirpykloms, odontologijos kabinetams, grožio salonams ir t.t.), kurių biudžetas neleidžia turėti nuosavos IS, registruoti būsimus vizitus, o jų klientams – rezervuoti pasirinktą vizito laiką.

Šią informacinę sistemą sudarys dvi esminės dalys – aplikacijų programavimo sąsaja (angl. API) bei grafinė naudotojo sąsaja, realizuota kaip WEB aplikacija.

Įmonė ar verslas, norėdamas pradėti naudotis informacine sistema ir suteikti galimybę savo klientams rezervuoti vizito laikus, iš pradžių turės užsiregistruoti – nurodyti veiklos pavadinimą, rūšį, kontaktus. Atlikusi šiuos veiksmus įmonė turės sulaukti administratoriaus patvirtinimo. Gavusi patvirtinimą įmonė galės pridėti savo teikiamas paslaugas (paslaugos iš esmės galėtų būti ir tokios pačios, tiesiog skirtųsi tas paslaugas suteikiantys specialistai) – kiekvienai jų reikės nurodyti aprašymą, tą paslaugą suteikiančio specialisto darbo laiką, specialisto kontaktus. Klientas, norėdamas atlikti laiko rezervaciją tam tikroje įmonėje, tam tikrai paslaugai, turės taip pat užsiregistruoti nurodydamas savo asmeninius duomenis – vardą, pavardę, el. paštą, telefono numerį. Atlikęs pasirinkto laiko rezervaciją klientas turės sulaukti patvirtinimo iš įmonės, o vėliau, esant poreikiui, tiek klientas, tiek įmonė rezervaciją galės atšaukti.

### **1.2. Funkciniai reikalavimai**

Neregistruotas sistemos naudotojas (svečias) galės:
1. Peržiūrėti įmonių sąrašą
2. Peržiūrėti informaciją apie konkrečią įmonę
3. Peržiūrėti konkrečios įmonės teikiamas paslaugas
4. Peržiūrėti konkrečios įmonės ir konkrečios paslaugos sukurtas rezervacijas (laisvus vizitų laikus)
5. Užsiregistruoti kaip klientas
6. Užsiregistruoti kaip paslaugas teikianti įmonė

Registruotas sistemos naudotojas (įmonė) galės:
1. Prisijungti
2. Atsijungti
3. Pridėti teikiamą paslaugą
4. Peržiūrėti paslaugos informaciją
5. Atnaujinti paslaugos informaciją
6. Ištrinti teikiamą paslaugą
7. Peržiūrėti savo įmonės teikiamų paslaugų sąrašą
8. Peržiūrėti informaciją apie konkrečią rezervaciją
9. Peržiūrėti konkrečios paslaugos rezervacijų sąrašą
10. Atnaujinti įmonės informaciją

Registruotas sistemos naudotojas (klientas) galės:
1. Prisijungti
2. Atsijungti
3. Peržiūrėti įmonių sąrašą
4. Peržiūrėti informaciją apie konkrečią įmonę
5. Peržiūrėti konkrečios įmonės teikiamas paslaugas
6. Peržiūrėti konkrečios įmonės ir konkrečios paslaugos sukurtas rezervacijas (laisvus vizitų laikus)
7. Sukurti rezervaciją
8. Atšaukti (ištrinti) konkrečią rezervaciją
9. Atnaujinti rezervacijos duomenis
10. Peržiūrėti savo visų rezervacijų sąrašą

Registruotas sistemos naudotojas (administratorius) galės:
1. Prisijungti
2. Atsijungti
3. Peržiūrėti įmonių sąrašą
4. Peržiūrėti informaciją apie konkrečią įmonę
5. Peržiūrėti konkrečios įmonės teikiamas paslaugas
6. Peržiūrėti konkrečios įmonės ir konkrečios paslaugos sukurtas rezervacijas (laisvus vizitų laikus)
7. Patvirtinti įmonės registraciją
8. Pašalinti įmonę

## 2. Sistemos architektūra

### **2.1. Pasirinktos technologijos**

Sistemą sudarys dvi dalys:
* Serverio pusė (aplikacijų programavimo sąsaja) – ji bus realizuota su Go programavimo kalba. Duomenų bazės valdymo sistema buvo pasirinkta PostgreSQL.
* Kliento pusė – ji bus realizuota su JavaScript biblioteka React.

### **2.2. Diegimo diagrama**

2.1 pav. pavaizduota sistemos diegimo diagrama. Sistemos talpinimui bus panaudotas Amazon Web Services serveris, kuriame sistemos naudotojų užklausas HTTP protokolu apdoros Traefik atvirkštinis tarpinis serveris (angl. reverse proxy) – jis HTTP užklausas persiųs arba į aplikacijų programavimo sąsają (serverio dalį), arba į WEB aplikaciją (kliento dalį), tuomet sulauks atsakymo ir jį persiųs atgal sistemos naudotojui. Sistemos realizacijai prireiks duomenų bazės valdymo serverio, o komunikaciją su juo atliks aplikacijų programavimo sąsaja TCP/IP protokolu.

![](.README_images/deployment.png)

**pav. 2.1 Sistemos diegimo diagrama**

## 3. Naudotojo sąsajos projektas

Žemiau pateikiami projektuojamos sąsajos langų wireframe`ai ir juos atitinkančių realizacijų langų iškarpos.

### **3.1. Pradinis langas**

Wireframe:

![](.README_images/01_Home.png)

Realizacijos langas:

![](.README_images/01_Home_real.png)

### **3.2. Informacijos apie įmonę langas**

Wireframe:

![](.README_images/02_CompanyInformation.png)

Realizacijos langas:

![](.README_images/02_CompanyInformation_real.png)

### **3.3. Prisijungimo langas**

Wireframe:

![](.README_images/03_Login.png)

Realizacijos langas:

![](.README_images/03_Login_real.png)

### **3.4. Kliento registracijos langas**

Wireframe:

![](.README_images/04_ClientRegistration.png)

Realizacijos langas:

![](.README_images/04_ClientRegistration_real.png)

### **3.5. Įmonės registracijos langas**

Wireframe:

![](.README_images/05_CompanyRegistration.png)

Realizacijos langas:

![](.README_images/05_CompanyRegistration_real.png)

### **3.6. Naujos rezervacijos sukūrimo langas**

Wireframe:

![](.README_images/06_CreateReservation.png)

Realizacijos langas:

![](.README_images/06_CreateReservation_real.png)

### **3.7. Kliento rezervacijų sąrašo langas**

Wireframe:

![](.README_images/07_ClientReservations.png)

Realizacijos langas:

![](.README_images/07_ClientReservations_real.png)

### **3.8. Kliento rezervacijos peržiūros langas**

Wireframe:

![](.README_images/08_ClientViewReservation.png)

Realizacijos langas:

![](.README_images/08_ClientViewReservation_real.png)

### **3.9. Kliento rezervacijos atnaujinimo langas**

Wireframe:

![](.README_images/09_ClientEditReservation.png)

Realizacijos langas:

![](.README_images/09_ClientEditReservation_real.png)

### **3.10. Administratoriaus pradinis langas**

Wireframe:

![](.README_images/10_AdminHome.png)

Realizacijos langas:

![](.README_images/10_AdminHome_real.png)

### **3.11. Įmonės informacijos atnaujinimo langas**

Wireframe:

![](.README_images/11_UpdateCompany.png)

Realizacijos langas:

![](.README_images/11_UpdateCompany_real.png)

### **3.12. Įmonės teikiamų paslaugų sąrašo langas**

Wireframe:

![](.README_images/12_MyServices.png)

Realizacijos langas:

![](.README_images/12_MyServices_real.png)

### **3.13. Įmonės paslaugos atnaujinimo langas**

Wireframe:

![](.README_images/13_UpdateService.png)

Realizacijos langas:

![](.README_images/13_UpdateService_real.png)

### **3.14. Įmonės naujos paslaugos sukūrimo langas**

Wireframe:

![](.README_images/14_CreateService.png)

Realizacijos langas:

![](.README_images/14_CreateService_real.png)

### **3.15. Įmonės paslaugų rezervacijų sąrašo langas**

Wireframe:

![](.README_images/15_CompanyReservations.png)

Realizacijos langas:

![](.README_images/15_CompanyReservations_real.png)

## 4. API specifikacija

### General types

#### Error
| Field   | Type   | Description     |
|---------|--------|-----------------|
| name    | string | Error code name |
| message | string | Error message   |

### General errors

#### Not found
| name      | message   | HTTP status code |
|-----------|-----------|------------------|
| not_found | Not found | 404              |

#### Forbidden
| name      | message   | HTTP status code |
|-----------|-----------|------------------|
| forbidden | Forbidden | 403              |

#### Unauthorized
| name         | message      | HTTP status code |
|--------------|--------------|------------------|
| unauthorized | Unauthorized | 401              |

#### Faulty token
| name         | message               | HTTP status code |
|--------------|-----------------------|------------------|
| faulty_token | Faulty token provided | 400              |

#### Invalid token
| name                     | message                     | HTTP status code |
|--------------------------|-----------------------------|------------------|
| token_invalid_or_expired | Token is invalid or expired | 401              |

#### Bad request
| name         | message      | HTTP status code |
|--------------|--------------|------------------|
| bad_request  | Bad request  | 400              |

#### Internal server error
| name                  | message               | HTTP status code |
|-----------------------|-----------------------|------------------|
| internal_server_error | Internal server error | 500              |

#### Invalid input
| name          | message                              | HTTP status code |
|---------------|--------------------------------------|------------------|
| invalid_input | Input could not pass the validations | 400              |

#### Unknown
| name     | message       | HTTP status code |
|----------|---------------|------------------|
| unknown  | Unknown error | 500              |

### General response
| Field | Type            | Null?      | Description                                              |
|-------|-----------------|------------|----------------------------------------------------------|
| error | [Error](#error) | on success | If response is unsuccessful, error code will be non null |
| data  | *any*           | on error   | The data type is specific to each endpoint               |

### POST /api/v1/users/authenticate

Authentication endpoint. Returns JWT token if authentication is successful.

#### Resource url
```
https://reservio.hs.vc/api/v1/users/authenticate
```

#### Resource information
|     Response formats     | JSON |
|:------------------------:|:----:|
| Requires authentication? |  No  |

#### Request parameters
| Name     | Type   | Required? | Description          | Validations                     | Example        |
|----------|--------|-----------|----------------------|---------------------------------|----------------|
| email    | string | yes       | User's email address | less than or equal to 200 chars | user@gmail.com |
| password | string | yes       | User's password      | between 8 and  100 chars        | Secret444!     |

#### Successful response
HTTP status code: `200`

Fields:

| Name         | Type   | Can be null? | Description       |
|--------------|--------|--------------|-------------------|
| accessToken  | string | no           | JWT access token  |
| refreshToken | string | no           | JWT refresh token |

#### Possible errors:

| name                            | message             | HTTP status code |
|---------------------------------|---------------------|------------------|
| [invalid_input](#invalid-input) |                     |                  |
| invalid_credentials             | Invalid credentials | 401              |
| [unknown](#unknown)             |                     |                  |

#### Example request
```
POST /api/v1/users/authenticate HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json

{
    "email": "user@gmail.com",
    "password": "Secret444!"
}
```

#### Example response
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "error": null,
    "data": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njg3MTMyMjcuNzMzMjM1LCJpYXQiOjE2Njg3MTAyMjcuNzMzMjQyLCJpc3MiOiJyZXNlcnZpbyIsInVzZXJJRCI6IjdiNTYyNGMyLTc3MTMtNGJiMy1iOGQ1LWVhMjliNTc5MjExNSIsInJvbGUiOiJjb21wYW55IiwicnRJRCI6IjY3MWVmNTNmLWY1YWYtNDM0NS05ODQzLWJhMzZkOThkNzlkYiJ9.YzCdBuLYFohulk50yF-hWguSPWR41H2pwqgGXPpnE5Q",
        "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njg3MTc0MjcuNzMzMTc0LCJpYXQiOjE2Njg3MTAyMjcuNzMzMTgyLCJpc3MiOiJyZXNlcnZpbyIsInJ0SUQiOiI2NzFlZjUzZi1mNWFmLTQzNDUtOTg0My1iYTM2ZDk4ZDc5ZGIifQ.aJ20LIyi1M5dXm3RbhKx26WK85q07EamUik0QGd8CqA"
    }
}
```

### POST /api/v1/tokens

Refresh token endpoint. Returns new JWT token if refresh token is valid.

#### Resource url
```
https://reservio.hs.vc/api/v1/tokens
```

#### Resource information
|     Response formats     | JSON |
|:------------------------:|:----:|
| Requires authentication? | Yes  |
|      Required role       | any  |

#### Request parameters
| Name         | Type   | Required? | Description       | Validations | Example                                                                                                                                                                                                                                 |
|--------------|--------|-----------|-------------------|-------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| refreshToken | string | yes       | JWT refresh token | -           | eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njg3MTc0MjcuNzMzMTc0LCJpYXQiOjE2Njg3MTAyMjcuNzMzMTgyLCJpc3MiOiJyZXNlcnZpbyIsInJ0SUQiOiI2NzFlZjUzZi1mNWFmLTQzNDUtOTg0My1iYTM2ZDk4ZDc5ZGIifQ.aJ20LIyi1M5dXm3RbhKx26WK85q07EamUik0QGd8CqA |

#### Successful response
HTTP status code: `200`

Fields:

| Name         | Type   | Can be null? | Description       |
|--------------|--------|--------------|-------------------|
| accessToken  | string | no           | JWT access token  |

#### Possible errors:

| name                            | message                     | HTTP status code |
|---------------------------------|-----------------------------|------------------|
| [invalid_input](#invalid-input) |                             |                  |
| [faulty_token](#faulty-token)   |                             |                  |
| [invalid_token](#invalid-token) |                             |                  |
| [unknown](#unknown)             |                             |                  |

#### Example request
```
POST /api/v1/tokens HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json

{
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njg3MTc0MjcuNzMzMTc0LCJpYXQiOjE2Njg3MTAyMjcuNzMzMTgyLCJpc3MiOiJyZXNlcnZpbyIsInJ0SUQiOiI2NzFlZjUzZi1mNWFmLTQzNDUtOTg0My1iYTM2ZDk4ZDc5ZGIifQ.aJ20LIyi1M5dXm3RbhKx26WK85q07EamUik0QGd8CqA"
}
```

#### Example response
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "error": null,
    "data": {
        "accessToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njg3MTc0MzAuNDQ0MjcyLCJpYXQiOjE2Njg3MTQ0MzAuNDQ0Mjc5LCJpc3MiOiJyZXNlcnZpbyIsInVzZXJJRCI6IjdiNTYyNGMyLTc3MTMtNGJiMy1iOGQ1LWVhMjliNTc5MjExNSIsInJvbGUiOiJjb21wYW55IiwicnRJRCI6IjY3MWVmNTNmLWY1YWYtNDM0NS05ODQzLWJhMzZkOThkNzlkYiJ9.Y2BovB5Tuq3_FLiDdjGUj0XYQB_cG0umWpNJSEfpy9Q"
    }
}
```

### POST /api/v1/users/logout

Logout endpoint. Invalidates refresh token.

#### Resource url
```
https://reservio.hs.vc/api/v1/users/logout
```

#### Resource information
|     Response formats     | JSON |
|:------------------------:|:----:|
| Requires authentication? | Yes  |
|      Required role       | any  |

#### Request parameters
*none*

#### Successful response
HTTP status code: `204`

#### Possible errors:

| name                            | message                     | HTTP status code |
|---------------------------------|-----------------------------|------------------|
| [unauthorized](#unauthorized)   |                             |                  |
| [faulty_token](#faulty-token)   |                             |                  |
| [invalid_token](#invalid-token) |                             |                  |
| [unknown](#unknown)             |                             |                  |

#### Example request
```
POST /api/v1/users/logout HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json
```

#### Example response
```
HTTP/1.1 204 No Content
Content-Type: application/json
```

### POST /api/v1/clients

Create client endpoint. Creates new client and returns its data.

#### Resource url
```
https://reservio.hs.vc/api/v1/clients
```

#### Resource information
|     Response formats     | JSON |
|:------------------------:|:----:|
| Requires authentication? |  No  |

#### Request parameters
| Name      | Type   | Required? | Description          | Validations                     | Example        |
|-----------|--------|-----------|----------------------|---------------------------------|----------------|
| email     | string | yes       | Client email address | less than or equal to 200 chars | user@gmail.com |
| password  | string | yes       | Client password      | between 8 and 100 chars         | Secret444!     |
| firstName | string | yes       | Client first name    | between 3 and 50 chars          | John           |
| lastName  | string | yes       | Client last name     | between 3 and 50 chars          | Devo           |
| phone     | string | yes       | Client phone number  | e164 standart                   | +37061354544   |

#### Successful response
HTTP status code: `200`

Fields:

| Name      | Type   | Can be null? | Description          |
|-----------|--------|--------------|----------------------|
| id        | string | no           | Created user ID      |
| email     | string | no           | Client email address |
| firstName | string | no           | Client first name    |
| lastName  | string | no           | Client last name     |
| phone     | string | no           | Client phone number  |

#### Possible errors:

| name                            | message              | HTTP status code |
|---------------------------------|----------------------|------------------|
| [invalid_input](#invalid-input) |                      |                  |
| email_already_exists            | Email already exists | 422              |
| [unknown](#unknown)             |                      |                  |

#### Example request
```
POST /api/v1/clients HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json

{
    "email": "user@gmail.com",
    "password": "Secret444!",
    "firstName": "John",
    "lastName": "Devo",
    "phone": "+37061354544"
}
```

#### Example response
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "error": null,
    "data": {
        "id": "3b218c3f-b443-4c15-8984-33654aa16a82",
        "email": "user@gmail.com"
        "firstName": "John",
        "lastName": "Devo",
        "phone": "+37061354544",
    }
}
```

### GET /api/v1/clients/:clientID/reservations

Returns all client reservations.

#### Resource url
```
https://reservio.hs.vc/api/v1/clients/:clientID/reservations
```

#### Resource information
|     Response formats     |  JSON  |
|:------------------------:|:------:|
| Requires authentication? |  Yes   |
|      Required role       | Client |

#### Request parameters
| Name      | Type   | Required? | Description | Validations | Example                              |
|-----------|--------|-----------|-------------|-------------|--------------------------------------|
| :clientID | string | yes       | Client ID   | -           | d662e339-6fd0-4de3-b3d2-7118ce70ee53 |

#### Successful response
HTTP status code: `200`

Fields:

| Name         | Type          | Can be null? | Description           |
|--------------|---------------|--------------|-----------------------|
| reservations | Reservation[] | no           | Array of reservations |

Reservation type:

| Name    | Type    | Can be null? | Description         |
|---------|---------|--------------|---------------------|
| id      | string  | no           | Reservation ID      |
| service | Service | no           | Reservation service |
| date    | string  | no           | Reservation date    |
| comment | string  | yes          | Reservation comment |

Service type:

| Name            | Type                      | Can be null? | Description              |
|-----------------|---------------------------|--------------|--------------------------|
| id              | string                    | no           | Service ID               |
| company         | Company                   | no           | Service company          |
| title           | string                    | no           | Service title            |
| description     | string                    | no           | Service description      |
| specialistName  | string                    | yes          | Service specialist name  |
| specialistPhone | string                    | yes          | Service specialist phone |
| visitDuration   | number                    | no           | Service visit duration   |
| workSchedule    | Map<Weekday, DaySchedule> | no           | Service work schedule    |

Weekday type:

| Name | Type   | Can be null? | Description                                                            |
|------|--------|--------------|------------------------------------------------------------------------|
|      | string | no           | One of: monday, tuesday, wednesday, thursday, friday, saturday, sunday |

DaySchedule type:

| Name | Type   | Can be null? | Description                |
|------|--------|--------------|----------------------------|
| from | string | no           | Day start time (eg. 10:30) |
| to   | string | no           | Day end time (eg. 18:00)   |

Company type:

| Name        | Type    | Can be null? | Description                   |
|-------------|---------|--------------|-------------------------------|
| id          | string  | no           | Company ID                    |
| email       | string  | no           | Company email address         |
| name        | string  | no           | Company name                  |
| address     | string  | no           | Company address               |
| description | string  | no           | Company description           |
| approved    | boolean | no           | Is company approved by admin? |

#### Possible errors:

| name                            | message                     | HTTP status code |
|---------------------------------|-----------------------------|------------------|
| [unauthorized](#unauthorized)   |                             |                  |
| [faulty_token](#faulty-token)   |                             |                  |
| [invalid_token](#invalid-token) |                             |                  |
| [forbidden](#forbidden)         |                             |                  |
| [invalid_input](#invalid-input) |                             |                  |
| client_not_found                | Client not found            | 404              |
| [unknown](#unknown)             |                             |                  |

#### Example request
```
GET /api/v1/clients/d662e339-6fd0-4de3-b3d2-7118ce70ee53/reservations HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json
```

#### Example response
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "error": null,
    "data": {
        "reservations": [
            {
                "id": "4ee0457e-e7ff-45bc-9d86-9623ce3e9d15",
                "service": {
                    "id": "cba4fef6-490b-4d56-af12-eb153c97893d",
                    "company": {
                        "id": "7b5624c2-7713-4bb3-b8d5-ea29b5792115",
                        "email": "servisas@gmail.com",
                        "name": "Automobilių servisas",
                        "address": "Ozo g. 10, Vilnius",
                        "description": "Automobilių servisas Vilniuje",
                        "approved": true
                    },
                    "title": "Padangų keitimas",
                    "description": "Pigus automobilio padangų keitimas",
                    "specialistName": "Jonas",
                    "specialistPhone": "+37064343332",
                    "visitDuration": 30,
                    "workSchedule": {
                        "monday": {
                            "from": "08:00",
                            "to": "16:00"
                        },
                        "thursday": {
                            "from": "10:00",
                            "to": "14:00"
                        },
                        "tuesday": {
                            "from": "11:30",
                            "to": "17:00"
                        },
                        "wednesday": {
                            "from": "08:00",
                            "to": "17:00"
                        }
                    }
                },
                "date": "2022-11-17 12:30",
                "comment": "Okaya"
            }
        ]
    }
}
```

### POST /api/v1/companies

Create company endpoint. Creates new company and returns its data.

#### Resource url
```
https://reservio.hs.vc/api/v1/companies
```

#### Resource information
|     Response formats     | JSON |
|:------------------------:|:----:|
| Requires authentication? |  No  |

#### Request parameters
| Name        | Type   | Required? | Description           | Validations                        | Example            |
|-------------|--------|-----------|-----------------------|------------------------------------|--------------------|
| email       | string | yes       | Company email address | less than or equal to 200 chars    | user@gmail.com     |
| password    | string | yes       | Company password      | between 8 and 100 chars            | Secret444!         |
| name        | string | yes       | Company name          | between 3 and 100 chars            | My company         |
| address     | string | yes       | Company address       | between 5 and 200 chars            | Vilnius, Lithuania |
| description | string | yes       | Company description   | greater than or equal to 200 chars | Very good company  |

#### Successful response
HTTP status code: `200`

Fields:

| Name        | Type    | Can be null? | Description                                                 |
|-------------|---------|--------------|-------------------------------------------------------------|
| id          | string  | no           | Created company ID                                          |
| email       | string  | no           | Company email address                                       |
| name        | string  | no           | Company name                                                |
| address     | string  | no           | Company address                                             |
| description | string  | no           | Company description                                         |
| approved    | boolean | no           | Is company approved by admin? When created, equals to false |

#### Possible errors:

| name                            | message              | HTTP status code |
|---------------------------------|----------------------|------------------|
| [invalid_input](#invalid-input) |                      |                  |
| email_already_exists            | Email already exists | 422              |
| [unknown](#unknown)             |                      |                  |

#### Example request
```
POST /api/v1/companies HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json

{
    "email": "user@gmail.com",
    "password": "Secret444!",
    "name": "My company",
    "address": "Vilnius, Lithuania",
    "description": "Very good company"
}
```

#### Example response
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "error": null,
    "data": {
        "id": "07c35941-0af3-4011-b42e-40bca152f2b8",
        "email": "user@gmail.com"
        "name": "My company",
        "address": "Vilnius, Lithuania",
        "description": "Very good company",
        "approved": false
    }
}
```

### GET /api/v1/companies

Returns all registered companies. When requested by admin, returns all companies, otherwise returns only approved companies.

#### Resource url
```
https://reservio.hs.vc/api/v1/companies
```

#### Resource information
|     Response formats     | JSON |
|:------------------------:|:----:|
| Requires authentication? |  No  |

#### Request parameters
*none*

#### Successful response
HTTP status code: `200`

Fields:

| Name      | Type      | Can be null? | Description        |
|-----------|-----------|--------------|--------------------|
| companies | Company[] | no           | Array of companies |

Company type:

| Name        | Type    | Can be null? | Description                                                                          |
|-------------|---------|--------------|--------------------------------------------------------------------------------------|
| id          | string  | no           | Company ID                                                                           |
| email       | string  | no           | Company email address                                                                |
| name        | string  | no           | Company name                                                                         |
| address     | string  | no           | Company address                                                                      |
| description | string  | no           | Company description                                                                  |
| approved    | boolean | no           | Is company approved by admin? If request is made not by admin, always equals to true |

#### Possible errors:

| name                            | message                     | HTTP status code |
|---------------------------------|-----------------------------|------------------|
| [unknown](#unknown)             |                             |                  |

#### Example request
```
GET /api/v1/companies HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json
```

#### Example response
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "error": null,
    "data": {
        "companies": [
            {
                "id": "7b5624c2-7713-4bb3-b8d5-ea29b5792115",
                "email": "servisas@gmail.com",
                "name": "Automobilių servisas",
                "address": "Ozo g. 10, Vilnius",
                "description": "Automobilių servisas Vilniuje",
                "approved": true
            },
            {
                "id": "e91f4c92-1371-48a1-a745-7d66d2178e15",
                "email": "plovykla@gmail.com",
                "name": "Plovykla",
                "address": "Vilniaus g. 17, Kaunas",
                "description": "Automobilių plovykla Kaune",
                "approved": true
            }
        ]
    }
}
```

### GET /api/v1/companies/:companyID

Returns company data by its ID.

#### Resource url
```
https://reservio.hs.vc/api/v1/companies/:companyID
```

#### Resource information
|     Response formats     | JSON |
|:------------------------:|:----:|
| Requires authentication? |  No  |

#### Request parameters
| Name       | Type   | Required? | Description | Validations | Example                              |
|------------|--------|-----------|-------------|-------------|--------------------------------------|
| :companyID | string | yes       | Company ID  | -           | 7b5624c2-7713-4bb3-b8d5-ea29b5792115 |


#### Successful response
HTTP status code: `200`

Fields:

| Name        | Type    | Can be null? | Description                   |
|-------------|---------|--------------|-------------------------------|
| id          | string  | no           | Company ID                    |
| email       | string  | no           | Company email address         |
| name        | string  | no           | Company name                  |
| address     | string  | no           | Company address               |
| description | string  | no           | Company description           |
| approved    | boolean | no           | Is company approved by admin? |

#### Possible errors:

| name                | message           | HTTP status code |
|---------------------|-------------------|------------------|
| company_not_found   | Company not found | 404              |
| [unknown](#unknown) |                   |                  |

#### Example request
```
GET /api/v1/companies/7b5624c2-7713-4bb3-b8d5-ea29b5792115 HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json
```

#### Example response
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "error": null,
    "data": {
        "id": "7b5624c2-7713-4bb3-b8d5-ea29b5792115",
        "email": "servisas@gmail.com",
        "name": "Automobilių servisas",
        "address": "Ozo g. 10, Vilnius",
        "description": "Automobilių servisas Vilniuje",
        "approved": true
    }
}
```

### PATCH /api/v1/companies/:companyID

Updates company data by its ID.

#### Resource url
```
https://reservio.hs.vc/api/v1/companies/:companyID
```

#### Resource information
|     Response formats     |       JSON       |
|:------------------------:|:----------------:|
| Requires authentication? |       Yes        |
|      Required role       | Company or Admin |

#### Request parameters
| Name       | Type   | Required? | Description | Validations | Example                              |
|------------|--------|-----------|-------------|-------------|--------------------------------------|
| :companyID | string | yes       | Company ID  | -           | 7b5624c2-7713-4bb3-b8d5-ea29b5792115 |

| Name        | Type    | Required? | Description                                      | Validations                        | Example            |
|-------------|---------|-----------|--------------------------------------------------|------------------------------------|--------------------|
| name        | string  | no        | Company name. Can only be set by company.        | between 3 and 100 chars            | My company         |
| address     | string  | no        | Company address. Can only be set by company.     | between 5 and 200 chars            | Vilnius, Lithuania |
| description | string  | no        | Company description. Can only be set by company. | greater than or equal to 200 chars | Very good company  |
| approved    | boolean | no        | Is company approved? Can only be set by admin.   | -                                  | true               |

#### Successful response
HTTP status code: `204`

#### Possible errors:

| name                            | message           | HTTP status code |
|---------------------------------|-------------------|------------------|
| [unauthorized](#unauthorized)   |                   |                  |
| [faulty_token](#faulty-token)   |                   |                  |
| [invalid_token](#invalid-token) |                   |                  |
| [forbidden](#forbidden)         |                   |                  |
| [invalid_input](#invalid-input) |                   |                  |
| company_not_found               | Company not found | 404              |
| nothing_to_update               | Nothing to update | 400              |
| [unknown](#unknown)             |                   |                  |

#### Example request
```
PATCH /api/v1/companies/7b5624c2-7713-4bb3-b8d5-ea29b5792115 HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json

{
    "name": "My company",
    "address": "Vilnius, Lithuania",
}
```

#### Example response
```
HTTP/1.1 204 No Content
Content-Type: application/json
```

### DELETE /api/v1/companies/:companyID

Deletes company by its ID.

#### Resource url
```
https://reservio.hs.vc/api/v1/companies/:companyID
```

#### Resource information
|     Response formats     | JSON  |
|:------------------------:|:-----:|
| Requires authentication? |  Yes  |
|      Required role       | Admin |

#### Request parameters
| Name       | Type   | Required? | Description | Validations | Example                              |
|------------|--------|-----------|-------------|-------------|--------------------------------------|
| :companyID | string | yes       | Company ID  | -           | 7b5624c2-7713-4bb3-b8d5-ea29b5792115 |

#### Successful response
HTTP status code: `204`

#### Possible errors:

| name                            | message           | HTTP status code |
|---------------------------------|-------------------|------------------|
| [unauthorized](#unauthorized)   |                   |                  |
| [faulty_token](#faulty-token)   |                   |                  |
| [invalid_token](#invalid-token) |                   |                  |
| [forbidden](#forbidden)         |                   |                  |
| [invalid_input](#invalid-input) |                   |                  |
| company_not_found               | Company not found | 404              |
| [unknown](#unknown)             |                   |                  |

#### Example request
```
DELETE /api/v1/companies/7b5624c2-7713-4bb3-b8d5-ea29b5792115 HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json
```

#### Example response
```
HTTP/1.1 204 No Content
Content-Type: application/json
```

### POST /api/v1/companies/:companyID/services

Create service endpoint. Creates new company service and returns its data.

#### Resource url
```
https://reservio.hs.vc/api/v1/companies/:companyID/services
```

#### Resource information
|     Response formats     |  JSON   |
|:------------------------:|:-------:|
| Requires authentication? |   Yes   |
|      Required role       | Company |

#### Request parameters
| Name       | Type   | Required? | Description | Validations | Example                              |
|------------|--------|-----------|-------------|-------------|--------------------------------------|
| :companyID | string | yes       | Company ID  | -           | 7b5624c2-7713-4bb3-b8d5-ea29b5792115 |

| Name            | Type                      | Required? | Description              | Validations                      | Example                                          |
|-----------------|---------------------------|-----------|--------------------------|----------------------------------|--------------------------------------------------|
| title           | string                    | yes       | Service title            | between 3 and 100 chars          | Men's haircut                                    |
| description     | string                    | yes       | Service description      | greater than or equal to 5 chars | Cheap men's haircut                              |
| specialistName  | string                    | no        | Service specialist name  | between 5 and 100 chars          | John                                             |
| specialistPhone | string                    | no        | Service specialist phone | e164 standart                    | +37064343333                                     |
| visitDuration   | number                    | yes       | Service visit duration   | greater than 0                   | 30                                               |
| workSchedule    | Map<Weekday, DaySchedule> | yes       | Service work schedule    | size greater than 0              | { "monday": { "from": "08:00", "to": "17:00" } } |

Weekday type:

| Name | Type   | Required? | Description            | Validations                                                            | Example |
|------|--------|-----------|------------------------|------------------------------------------------------------------------|---------|
|      | string | yes       | Work schedule week day | One of: monday, tuesday, wednesday, thursday, friday, saturday, sunday | friday  |

DaySchedule type:

| Name | Type   | Required? | Description    | Validations       | Example |
|------|--------|-----------|----------------|-------------------|---------|
| from | string | yes       | Day start time | time format HH:mm | 10:30   |
| to   | string | yes       | Day end time   | time format HH:mm | 18:00   |

#### Successful response
HTTP status code: `200`

Fields:

| Name            | Type                      | Can be null? | Description              |
|-----------------|---------------------------|--------------|--------------------------|
| id              | string                    | no           | Service ID               |
| companyID       | string                    | no           | Company ID               |
| title           | string                    | no           | Service title            |
| description     | string                    | no           | Service description      |
| specialistName  | string                    | yes          | Service specialist name  |
| specialistPhone | string                    | yes          | Service specialist phone |
| visitDuration   | number                    | no           | Service visit duration   |
| workSchedule    | Map<Weekday, DaySchedule> | no           | Service work schedule    |

#### Possible errors:

| name                            | message              | HTTP status code |
|---------------------------------|----------------------|------------------|
| [unauthorized](#unauthorized)   |                      |                  |
| [faulty_token](#faulty-token)   |                      |                  |
| [invalid_token](#invalid-token) |                      |                  |
| [forbidden](#forbidden)         |                      |                  |
| [invalid_input](#invalid-input) |                      |                  |
| company_not_found               | Company not found    | 404              |
| [unknown](#unknown)             |                      |                  |

#### Example request
```
POST /api/v1/companies/7b5624c2-7713-4bb3-b8d5-ea29b5792115/services HTTP/1.1
Host: reservio.hs.vc
Content-Type: application/json

{
    "title": "Men's haircut",
    "description": "Cheap men's haircut",
    "specialistName": "John",
    "specialistPhone": "+37064343333",
    "visitDuration": 30,
    "workSchedule": {
        "monday": {
            "from": "08:00",
            "to": "17:00"
        },
        "tuesday": {
            "from": "11:45",
            "to": "17:00"
        },
        "wednesday": {
            "from": "08:00",
            "to": "17:00"
        }
    }
}
```

#### Example response
```
HTTP/1.1 200 OK
Content-Type: application/json

{
    "error": null,
    "data": {
        "id": "fda84995-0755-40e3-8ed4-129fc774125b",
        "companyID": "e91f4c92-1371-48a1-a745-7d66d2178e15",
        "title": "Men's haircut",
        "description": "Cheap men's haircut",
        "specialistName": "John",
        "specialistPhone": "+37064343333",
        "visitDuration": 30,
        "workSchedule": {
            "monday": {
                "from": "08:00",
                "to": "17:00"
            },
            "thursday": {
                "from": "10:15",
                "to": "14:00"
            },
            "tuesday": {
                "from": "11:45",
                "to": "17:00"
            }
        }
    }
}
```

## 5. Išvados

Šio modulio metu pavyko sėkmingai įgyvendinti užsibrėžtus projekto tikslus - sukurti *backend API*, ją tinkamai apsaugoti panaudojant *JWT token* authentifikacijai ir autorizacijai, be to, sukurti naudotojo sąsajos dalį, viską patalpinti debesyje, jog būtų galima išorinė prieiga prie sistemos, bei galiausiai, parengti galutinę ataskaitą. Tiesa, projektas nėra visiškai užbaigtas ir jį dar reikėtų nemažai patobulinti norint paleisti į rinką.
