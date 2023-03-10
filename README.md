## Uma Corrida de Revezamento
Este é um projeto em Go que simulará uma corrida de revezamento.

### Problema
A corrida de revezamento surgiu durante os jogos olímpicos da Grécia e, de maneira geral, é uma competição em que os membros de uma equipe revezam-se na conclusão de partes do local no qual a competição é feita ou realizando todos uma determinada ação. Na modalidade do atletismo, a corrida de revezamento consiste de equipes de corredores, cada uma em uma raia, e esses corredores revezam-se em distâncias definidas carregando um bastão.
Ao percorrer a distância estabelecida, o bastão deve ser passado de um atleta para outro e assim sucessivamente, de forma síncrona, até chegar ao término do percurso total.
Nos tempos modernos há dois tipos básicos de corrida de revezamento, 4 x 100 e 4 x 400: a corrida é dividida em quatro etapas, cada uma contendo respectivamente 100 e 400 metros a serem percorridos. No fim de cada etapa, o atleta tem um espaço de 20 metros para entregar o bastão. Esse momento é delicado, pois uma entrega incorreta
pode gerar a desclassificação, o que também pode ocorrer caso o bastão não seja entregue nesse espaço de 20 metros.
Normalmente o corredor que vai passar o bastão para o próximo dá um sinal sonoro (um grito, por exemplo) para chamar sua atenção para que estenda a mão e esteja pronto para imediatamente pegar o bastão. Dessa forma, a transmissão do bastão requer uma sincronização perfeita para que a etapa não seja perdida, sendo necessário que os dois corredores que estejam envolvidos na transmissão do bastão estejam prontos exatamente ao mesmo tempo.
Escreva um programa em Go que simule uma corrida de revezamento em que uma equipe possui quatro corredores.
O segundo, terceiro e quarto corredores não podem começar a correr até que recebam o bastão entregue pelo corredor
que o antecedeu. Para simular a corrida do corredor na etapa em questão, pode-se utilizar o método Sleep provido pelo pacote time por um intervalo de tempo predefinido. O bastão deve ser passado para o próximo corredor até que o último corredor conclua o trecho a ser percorrido, momento em que o programa deverá ser encerrado. Por questões de simplicidade, desconsidere as regras válidas no mundo real em que o bastão necessita ser entregue nos 20 metros finais de cada etapa a ser percorrida.
Extra: Modifique o programa implementado para que haja mais de uma equipe de atletas participando da corrida de revezamento, cada uma em uma raia, de forma similar ao que ocorre no mundo real. Ao término de sua execução,
o programa deverá informar qual das equipes venceu a prova.

### Pré-requisitos
É necessário ter o Go instalado na máquina.

### Como executar
Para rodar o projeto, basta abrir o terminal, navegar até a pasta do projeto e digitar o comando abaixo:

```console
$ go run main.go
```

Isso irá compilar e executar o arquivo main.go.
