# Desafio Técnico V3

## Sobre a implemetação

A descrição original do desafio pode ser encontrada aqui:
https://github.com/v3-tecnologia/desafio/blob/main/README.md


Para esse projeto, foi criado um aplicativo escrito em Kotlin, com a biblioteca Jetpack Compose e arquitetura MVVM.

O projeto está apontando para o Java 18, então é necessário ter essa versão instalada e configurada, tanto nas variáveis de ambiente, como no Android Studio, ou fazer o projeto apontar para outra versão do compilador.


# Foram utilizados:

- Dagger Hilt para fazer a injeção de dependências.
- Retrofit para chamadas de rede.
- Moshi para serialização de objetos.
- SharedPreferences para salvar as informações localmente no cache do app.
- Material para ícones.
- FusedLocation para pegar o posicionamento GPS.
- Accompanist para gerenciamento de permissões.
- ComposeSensors para os dados do Gyroscópio.
- Camerax para fotos
- MLKit para detecção de rostos.

Também usei alguns arquivos isolados desse repositório: https://github.com/Hechio/CameraX-FaceDetection-MlKit/tree/master, que possui uma ótima implementação do Camerax+MLKit, mas foi necessário haver todo um trabalho de adaptação e limpeza dos métodos para ficar coerente e funcional com o projeto presente.

# Alguns design patterns utilizados:
- Singleton para repositórios e serviços.
- Wrapper para encapsular as respostas dos services.
- Facade para os métodos do SharedPreferences.

# Ideia e UI
Foi especificado que o aplicativo não teria tela, que os dados deveriam ser tratados apenas no background, por isso tive a ideia de fazer uma tela que poderia ser "ligada ou desligada" com um botão power para atender à especificação, mas também tendo um modo de visualizar alguns status sem ter q depender apenas olhar os logs da aplicação ou o debugger.

A tela principal tem o botão mencionado. No modo Desligado a tela fica escura, mas o processamento de dados ainda está rolando no background. 
No modo Ligado há um feed de logs reportando os status do background.
Também há mais dois botões embaixo, um para trocar a câmera q está sendo usada, entre a Traseira e a Frontal.
E um botão para limpar o feed. Importante: o botão não vai apagar os dados do Cache do aplicativo, apenas a lista q está sendo exibida! A lista é rolável e otimizada, então não tem problema se crescer demais.

O app fica constantemente coletando dados no background, da câmera, gps e giroscópio, mas não salva nada. Apenas a cada 10 segundos os dados daquele momento são salvos no cache local e enviados para os endpoints fake. A photo só é salva se um rosto for detectado.
Estou fazendo requisições reais para a URL https://jsonplaceholder.typicode.com/, mas o response das chamadas está sendo ignorado, já q sempre vai retornar erro.

# Pendências 
De pendências do projeto devido ao tempo limitado, ficaram: o crop da imagem e os testes, apenas. Todo o resto foi implementado.
Há alguns comentários pelo código com //TODO em locais que deveriam receber atenção, caso esse fosse um projeto real que fosse ser continuado.
