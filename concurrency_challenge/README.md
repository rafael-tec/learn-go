## Desafio de Concorrência
Desafio de concorrência usando Go, do módulo do curso Go Expert da Full Cycle.

### Objetivo
Fazer integração http com duas APIS para buscar o endereço pelo o resultado mais, considerar a resposta mais rápida.

As duas requisições devem ser feitas simultaneamente para as seguintes APIs:

- https://brasilapi.com.br/api/cep/v1/{cep}
- http://viacep.com.br/ws/{cep}/json/

### Requisitos
- Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

- O resultado da requisição deverá ser exibido no command line com os dados do endereço, bem como qual API a enviou.

- Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
