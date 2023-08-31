# Projeto Laboratório - Micro-serviço msCadEndBR (Cadastro de Endereços Brasil)
# Visão geral
O projeto CadEndBR é um laboratório desenvolvido na linguagem Go. Foi concebido com o intuíto de apoiar desenvolvedores em um curso que conclui para trabalharem com orquestrações, parsers, enriquecimento de dados e integrações com outros micros-services públicos como o viaCEP e BrasilAberto através do API Gateway Mulesoft. Este micro-serviço foi projetado para ser um LAB com uma solução eficiente e escalável, sendo hospedado em um ambiente cloud e pronto para ser consumido em integrações com o Mulesoft, objetivo do estudo.

O principal objetivo do CadEndBR é fornecer uma base sólida para a orquestração e enriquecimento de dados relacionados a endereços no Brasil, utilizando a linguagem DataWeave para manipulação e transformação dos dados.

# Contexto de uso
O micro-serviço CadEndBR foi desenvolvido para ser usado em conjunto com o API Gateway Mulesoft, onde será integrado a outros micro-serviços para realizar chamadas e transformações proporcionando experiencia de aprendizado no Mulesoft. A integração é feita por meio da orquestração e enriquecimento de dados (parser) usando a linguagem DataWeave e Expression Mode, utilização de Patterns de integração como o próprio Adapter facilitando a manipulação e a transformação dos dados recebidos de diferentes fontes de APIs e micro-services que no final irá compor um novo body com uma estrutura em json para realizar um post em uma base de dados no qual é esperado por este micro-serviço (CadEndBR) em cloud.

# Principais funcionalidades
- Exposição de endpoints RESTful para a realização de post e consultas. O micro-service conta com operações públicas para a geração de token e a validação deste token antes mesmo de consumir as operações privadas do micro-serviço.
- Basicamente você passará para a API exposta no Mulesoft o nome, o cep e o número da casa através do verbo POST, pois a intenção com base neste cep, recuperar o endereço em um destes micros-services publicos, realizar as devidas orquestrações e enriquecimento de dados usando o patttern adapter e por fim, o post já com os dados tratados conforme a necessidade do micro-service CadEndBR.
- Veja, você enviará um post, mas no MuleApp, será feito uma orquestração e antes dele fazer o post, irá reencaminhar parte da entrada do seus dados (cep) para outros micro-services (viaCEP ou brasilAberto) ambos buscadores de logradouro com base em cep e que tem seu payload de retorno distintos... Por isso precisamos de um adapter. A integração passará parte destes destes dados (cep), que devolverá o endereço deste CEP.
- Então, a orquestração da APP no Mulesoft, avaliará quais micro-servicos ela irá usar e então, deverá, receber esta informação, montar um outro payload com a junção do nome + o endereço retornado + o número residencial e então seu fluxo irá executará um POST para gerar e armazenar o token (publico) obtido no micro-serviço CadEndBR e na sequencia irá chamar outro POST passando o body montado no request da operação enderecos que por sua vez irá persistir estes dados em em núvem.
- Integração com o API Gateway Mulesoft para orquestração e enriquecimento de dados.
- Uso da linguagem DataWeave para transformação e manipulação dos dados.
- Hospedagem em ambiente cloud para escalabilidade e disponibilidade.

# Início rápido
Para começar a utilizar o micro-serviço CadEndBR, siga os passos abaixo:

- Clone o repositório Git:

bash
Copy code
 - git clone https://github.com/seuusuario/msCadEndBR.git
 - Navegue até a pasta do projeto e instale as dependências:

bash
Copy code
 - cd msCadEndBR
 - go get -d ./...
 - Compile e execute o projeto:

bash
Copy code
 - go build
 - ./msCadEndBR
 - Agora o micro-serviço está em execução e pronto para ser utilizado na integração com o Mulesoft.

# Suporte e contribuição
Se você tiver alguma dúvida, sugestão ou deseja contribuir para o projeto CadEndBR, sinta-se à vontade para abrir uma issue ou enviar um pull request no repositório Git.

Agradecemos seu interesse e apoio no desenvolvimento deste projeto!
