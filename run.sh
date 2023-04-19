#!/bin/bash

# Construa a imagem do contêiner
docker build -t mscadendbr:latest .

# Execute o contêiner e publique a porta
docker run -p 5001:5001 mscadendbr
