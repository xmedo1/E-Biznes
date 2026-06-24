**Zadanie 08**

Należy skonfigurować klienta Oauth2 (4.0). Dane o użytkowniku wraz z
tokenem powinny być przechowywane po stronie bazy serwera, a nowy
token (inny niż ten od dostawcy) powinien zostać wysłany do klienta
(React). Można zastosować mechanizm sesji lub inny dowolny (5.0).
Zabronione jest tworzenie klientów bezpośrednio po stronie React'a
wyłączając z komunikacji aplikację serwerową.

Prawidłowa komunikacja: react-sewer-dostawca-serwer(via return uri)-react.

:white_check_mark: 3.0 logowanie przez aplikację serwerową (bez Oauth2) [[commit](https://github.com/xmedo1/E-Biznes/commit/dc1bb1eaa55989f77d2f77ac77765a5d51460d45)] \
:white_check_mark: 3.5 rejestracja przez aplikację serwerową (bez Oauth2) [[commit](https://github.com/xmedo1/E-Biznes/commit/04dbaf37510c20b13bf4619cebd1426ced150c20)] \
:white_check_mark: 4.0 logowanie via Google OAuth2 [[commit](https://github.com/xmedo1/E-Biznes/commit/894c3f525f9e786157ec9dd28df9cbf1aee2ed12)] \
:white_check_mark: 4.5 logowanie via Facebook lub Github OAuth2 [[commit](https://github.com/xmedo1/E-Biznes/commit/d488e7acf1c5190c3dd7d8461f1b9c7f4455768d)] \
:white_check_mark: 5.0 zapisywanie danych logowania OAuth2 po stronie serwera [[commit](https://github.com/xmedo1/E-Biznes/commit/e1a41342d12761291683b99b4ec5eed6fadb48fb)] 

Klucz należy uzyskać na:
- https://console.cloud.google.com/apis/dashboard,
- https://developers.facebook.com/

[Link](https://www.youtube.com/watch?v=ZiUKDB4M_rY) do filmu. \
[Link](https://www.youtube.com/watch?v=5XFwd_CYRDk) do filmu z demo OAuth2 dla Google i Github.
