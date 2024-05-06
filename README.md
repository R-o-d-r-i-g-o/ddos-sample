# Aplicativo de Teste de Estresse em Go (ddos-sample)

Este é um aplicativo de teste de estresse em Go projetado para simular cargas pesadas em sistemas para verificar sua capacidade de resposta e desempenho com base em goroutines (processamento concorrente).

## Pré-requisitos

Certifique-se de ter o Go instalado na sua máquina. Você pode fazer o download e instalar o Go em: https://golang.org/dl/. Configure as variáveis de ambiente nativas do Go conforme necessário para um bom desempenho e uso dos cores de CPU disponíveis.

## Configurações

Configure as seguintes variáveis de ambiente para utilizar o aplicativo com base no número de requisições por segundo e no tempo (em segundos) de duração do teste. A URL do site em questão também deve ser inserida para a simulação dos testes.

```bash
# stress-test data
CONNECTIONS_PER_SECOND=
TIME_IN_SECONDS=

# website info
PATH_URL=
```

Lembre-se de substituir os valores CONNECTIONS_PER_SECOND, TIME_IN_SECONDS e PATH_URL pelas configurações desejadas para seu teste.
