# Documentação do Projeto

## Funcionalidades

- **Serviço em Segundo Plano para Coleta de Dados**  
  O aplicativo inclui um serviço que opera em segundo plano, projetado para coletar periodicamente dados de localização e giroscópio do dispositivo. Este serviço é configurado para realizar a coleta a cada 10 segundos, garantindo que informações detalhadas sobre a localização e movimento do dispositivo sejam capturadas continuamente. Isso é crucial para aplicações que exigem monitoramento constante, mesmo quando o aplicativo não está visível ou em uso ativo.

- **Persistência de Dados com Room**  
  A persistência dos dados coletados é gerida pela biblioteca Room, que fornece uma camada de abstração sobre o banco de dados SQLite. O Room facilita o armazenamento local dos dados de localização e giroscópio, permitindo que eles sejam salvos e recuperados de forma eficiente. Esta abordagem garante que os dados sejam preservados mesmo quando o dispositivo está offline, e que possam ser acessados quando o aplicativo é reaberto ou sincronizado com o backend.

- **Envio de Dados para Backend via Retrofit**  
  O aplicativo utiliza o Retrofit para enviar os dados coletados para um servidor remoto. O Retrofit é uma biblioteca que simplifica a comunicação com serviços web, permitindo o envio de dados de forma segura e eficiente. Os dados enviados incluem informações de localização e giroscópio, além de um ID único do dispositivo. Esse ID é usado para identificar o dispositivo de forma exclusiva e só é alterado em casos específicos, como restauração do sistema ou outras mudanças significativas.

- **Captura de Fotos e Detecção de Rostos**  
  O aplicativo possui uma funcionalidade para capturar fotos usando a câmera do dispositivo. Após a captura, o aplicativo utiliza tecnologias de reconhecimento facial para verificar se as fotos contêm rostos. Esta funcionalidade pode ser aplicada para validar identidades, registrar presenças ou outras aplicações que exigem verificação visual. A detecção de rostos é realizada automaticamente após cada captura de foto.

- **Notificações ao Usuário**  
  Durante a execução do serviço de coleta de dados, o aplicativo mantém o usuário informado através de notificações persistentes na barra de status do Android. Essas notificações indicam que o serviço está ativo e operando em segundo plano, fornecendo transparência sobre as operações em andamento e garantindo que o usuário esteja ciente da coleta de dados contínua.

## Requisitos

- **Android Studio**  
  O ambiente de desenvolvimento integrado (IDE) utilizado para o desenvolvimento e teste do aplicativo. Android Studio é necessário para criar, depurar e compilar o projeto, além de oferecer ferramentas para simulação e controle de versão.

- **Gradle**  
  Sistema de automação de build utilizado para gerenciar as dependências e compilar o projeto. Gradle facilita a integração de bibliotecas externas e a configuração do processo de build, garantindo que todas as dependências sejam resolvidas e que o aplicativo seja construído corretamente.

- **Permissões**  
  O aplicativo necessita de permissões específicas para operar corretamente. As permissões incluem acesso a dados de localização para a coleta contínua, capacidade para operar em segundo plano e acessar a câmera e arquivos da galeria para captura e armazenamento de fotos. É essencial que todas as permissões sejam configuradas corretamente no arquivo de manifesto do aplicativo para garantir que todas as funcionalidades possam ser executadas sem problemas.

## Configuração do Projeto

### Clonagem do Repositório

Para iniciar o desenvolvimento com o projeto, você deve clonar o repositório para o seu ambiente local. Isso garante que você tenha acesso ao código-fonte mais recente e possa começar a trabalhar nas funcionalidades e personalizações do aplicativo.

### Configuração das Permissões

É necessário declarar as permissões apropriadas no arquivo de manifesto do projeto para garantir que o aplicativo tenha acesso aos recursos necessários, como localização, câmera e armazenamento. Essas permissões devem ser configuradas conforme os requisitos específicos do aplicativo e as diretrizes da plataforma Android.

### Configuração dos Serviços

O projeto inclui dois serviços principais: um para coleta de dados de localização e giroscópio, e outro para captura de fotos. Cada serviço deve ser configurado para iniciar e operar conforme necessário, garantindo que todas as funcionalidades sejam executadas corretamente. O serviço de coleta de dados deve ser iniciado a partir da `MainActivity` e configurado para funcionar em segundo plano. O serviço de captura de fotos deve operar separadamente e verificar a presença de rostos nas imagens capturadas.

### Configuração do Room

O Room deve ser configurado para armazenar localmente os dados de localização e giroscópio. Isso envolve a criação das entidades, DAOs (Data Access Objects) e a configuração do banco de dados para garantir que os dados sejam persistidos corretamente.

### Configuração do Retrofit

O Retrofit deve ser configurado para gerenciar as requisições HTTP para o backend. Isso inclui a definição das interfaces de API, configuração dos conversores de dados e gerenciamento de endpoints para garantir que os dados coletados sejam enviados corretamente para o servidor remoto.

## Notas Adicionais

- **Permissões**  
  O aplicativo deve verificar e solicitar permissões adequadas com base na versão do Android e nas políticas de privacidade do usuário.

- **Serviços**  
  A operação contínua dos serviços é essencial para o funcionamento correto do aplicativo. Certifique-se de que os serviços estão configurados para iniciar e operar conforme esperado, e que o usuário esteja sempre informado sobre as atividades em segundo plano.
