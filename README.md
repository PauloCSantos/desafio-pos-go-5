# Desafio 5 Pós Go Lang - Full Cycle

## Objetivo

. Desenvolver um sistema em Go que receba um CEP, identifica a cidade e retorna o clima atual (temperatura em graus celsius, fahrenheit e kelvin) juntamente com a cidade. Esse sistema deverá implementar OTEL(Open Telemetry) e Zipkin.

## Funcionalidades

- Endpoint capaz de responder através do schema: { "cep": "29902555" }
- Retorna erro quando não recebeu CEP válido
- Quando não encontra a cidade
- Quando não encontra a temperatura da cidade

## Requisitos

**⚠️ Importante:**  
 Para utilizar esta aplicação, configure a variável de ambiente `API_KEY` com uma chave válida.

**APIKEY do weatherapi**: https://www.weatherapi.com/

## Configuração

1. No dir serviceB crie o arquivo .env usando o .env.example

   - Preencher com sua key do weatherapi

2. Na raiz do projeto

   - docker compose up

3. Use o endpoint `http://localhost:8000/getTemperature`:

   - É necessario ser POST
   - Enviar o cep nesse formato { "cep": "29902555" }
   - Só é valido CEP com numeros

4. Para acessar o ZIPKIN

   - `http://localhost:9411/`

5. OBS

   - Existe o arquivo .http que faz a chamada só é necessario a extensão REST Client
   - Para executar usando o devcontainer é necessario alterar o client do serviceA
     Use BaseURL: "http://localhost:8080/temperatureByCEP"
