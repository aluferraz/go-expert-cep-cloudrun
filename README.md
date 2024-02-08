# Busca CEP  - Cloudrun


Variáveis de ambiente:

| Variável        | Descrição                                                         |
|-----------------|-------------------------------------------------------------------|
| WEBSERVER_PORT  | A porta em que o servidor web estará disponível.                  |
| WEATHER_API_KEY | Chave de API para acessar a API de clima.                         |
| WEATHER_API_URL | URL da API de clima para obter dados de temperatura.              |
| CEP_API_URL     | URL da API de CEP para obter dados de localidade a partir do CEP. |


O projeto usa Viper para gerenciar as váriaveis de ambiente, que podem ser configuradas no OS, Google Cloud Console ou em um arquivo .env

O projeto possui testes integrados ao github actions e também pode ser testado com ``go test -v ./...``

Arquivo principal do projeto: https://github.com/aluferraz/go-expert-cep-cloudrun/blob/9af4e910f335af93eb9e50952489c6fd6a74995a/internal/usecase/get_temperature/get_temperature.go#L1

URL CloudRun: [https://go-expert-cep-cloudrun-tg3vmarlxq-uc.a.run.app](https://go-expert-cep-cloudrun-tg3vmarlxq-uc.a.run.app/?zipcode=30140091)https://go-expert-cep-cloudrun-tg3vmarlxq-uc.a.run.app/?zipcode=30140091 (Coloque o CEP desejado como parametro de url)
