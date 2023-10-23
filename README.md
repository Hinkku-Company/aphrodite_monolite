# aphrodite_monolite
Monolito para el API de aphrodite

# Pending

- [ ] Info User
- [ ] Lógica redis 
    - Se valida el token del header en cada petición para verificar contra redis si existe, si no existe en redis, se deniega el acceso y si existe en redis se valida con auth.ValidToken (el mismo escenario para token refresh)