# Busca CEP


Variáveis de ambiente:

| Variável        | Descrição                                                         |
|-----------------|-------------------------------------------------------------------|
| WEBSERVER_PORT  | A porta em que o servidor web estará disponível.                  |
| WEATHER_API_KEY | Chave de API para acessar a API de clima.                         |
| WEATHER_API_URL | URL da API de clima para obter dados de temperatura.              |
| CEP_API_URL     | URL da API de CEP para obter dados de localidade a partir do CEP. |


O projeto usa Viper para gerenciar as váriaveis de ambiente, que podem ser configuradas no OS, Google Cloud Console ou em um arquivo .env

O projeto possui testes integrados ao github actions e também pode ser testado com ``go test -v ./...``

URL CloudRun:
