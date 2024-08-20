# go-expert-lab-google-cloud-run
Pós Go Expert Lab Cloud Run

## Configuração
```
- Renomeie o arquivo .env.example para .env
- Substitua WEATHER_API_KEY pela sua chave da WeatherAPI 
```

### Rodar os testes
```bash
    make test
```

### Buildar a imagem docker e inicar a aplicação
```bash
    make start
```

### Parar a aplicação
```bash
    make stop
```

### Remover o container
```bash
    make clean
```

### Google Cloud Run

Acesse https://go-expert-lab-google-cloud-run-zdsocrcpfa-uc.a.run.app/weather/{cep}, informando o CEP desejado.

## <a name="license"></a> License

Copyright (c) 2024 [Hugo Castro Costa]

[Hugo Castro Costa]: https://github.com/hgtpcastro