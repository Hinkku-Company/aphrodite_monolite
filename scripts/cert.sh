# Genera ACCESS_TOKEN_PRIVATE_KEY.pem y ACCESS_TOKEN_PUBLIC_KEY.pem
openssl ecparam -name prime256v1 -genkey -noout -out ACCESS_TOKEN_PRIVATE_KEY.pem
openssl ec -in ACCESS_TOKEN_PRIVATE_KEY.pem -pubout -out ACCESS_TOKEN_PUBLIC_KEY.pem

# Genera REFRESH_TOKEN_PRIVATE_KEY.pem y REFRESH_TOKEN_PUBLIC_KEY.pem
openssl ecparam -name prime256v1 -genkey -noout -out REFRESH_TOKEN_PRIVATE_KEY.pem
openssl ec -in REFRESH_TOKEN_PRIVATE_KEY.pem -pubout -out REFRESH_TOKEN_PUBLIC_KEY.pem