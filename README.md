## Go RestApi

Objetivo deste projeto foi criar uma Api Rest usando Golang e Frameworks

Ferramentas utilizadas:

    - Docker
    - PostgreSQL
    - React

Como subir o servidor:

    1 - Primeiro tenha o docker instalado em sua máquina
    
    2 - Após isso rode `go mod tidy`
    
    3 - Logo após fazer a instalação dos pacotes go, rode:
    `docker-compose up`
    
    4 - Iniciará o banco de dados com o valor:
        USER - root
        PASS - root
        HOST - root
        
    5 - Instale os módulos Node abrindo a pasta de Frontend e rodando:
        - npm version
        - npm install
        - npm update
        - npm start 

Considerações finais:

    - Frontend :3000 / Backend :8000 / BD :5432
    - CRUD (Create/Read/Update/Delete)
