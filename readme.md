# Repositório para implementações, exemplos e exercícios de Algoritmo e Estruturas de Dados I

## ArrayList — Tempo de Cache

A comparação entre ordens de acesso a elementos pode ser feita através do script compacarao.bat (somente Windows), que pode ser checado via [aqui](linkdogit).

### Por que haveria diferença?

A diferença de tempo entre os resultados dá-se pelo melhor aproveitamento do [cache](pt.wikipedia.org/wiki/Cache_do_processador) quando se percorre a lista diretamente. Este ganho não é garantido para qualquer tipo de estrutura e poderia ser afetado pela estratégia de cache do processador, sendo sequer garantido este ganho para listas como um todo.

Como a ArrayList guarda seus valores numa espécie de _array_, os dados são dispostos continuamente na memória. Uma estratégia comum de cache é, ao trazer os dados de uma posição de memória para o processador, aproveitar a 'viagem' para trazer também dados na proximos ao dado lido, devido às chances desses dados serem requiridos futuramente.

Não se pode esperar uma diferença significativa de uma lista duplamente encadeada ainda que seja também uma lista, entretanto. Como a alocação de elementos nela se faz isoladamente, eles não necessariamente estarão próximos e possuem menos chances de serem trazidos 'de uma vez só' ao processador.

### Como são feitos os testes?

Primeiramente é criado um ArrayList de $10^N$ elementos, com $N \in [1, 8]$, e são adicionados $N-1$ elementos indo de $1$ a $N-1$. Através da _package_ [_Time_](https://pkg.go.dev/time@go1.21.0) da biblioteca padrão do Go é registrado o momento antes e depois da chamada de AddPos para adicionar 0 à posição 0. Cada teste de N elementos é feito 20 vezes para ambas as ordens, sendo mantido o menor tempo obtido nestas 20 vezes para comparação.

A forma que Go utiliza para contabilizar o tempo não é diferenciada entre um tempo de parede ou um tempo de ciclos de clock na descrição de _Time_.

Só se é mantido o tempo mínimo porque quaisquer atrasos em relação a estes são definitivamente não-relacionados à execução do programa.