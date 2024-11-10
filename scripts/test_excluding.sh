#!/bin/bash
export GOFLAGS="-buildvcs=false"

TEST_DIRS=$(go list ./... | grep -v '/mocks' \
                        | grep -v '/handlers' \
                        | grep -v '/repositories' \
                        | grep -v '/cmd' \
                        | grep -v '/database' \
                        | grep -v '/server' \
                        | grep -v '/routes' \
                        | grep -v '/dtos' \
                        | grep -v '/docs')


chmod +x ./scripts/test_excluding.sh

go test --cover $TEST_DIRS
if [[ $? -eq 0 ]]; then
    echo "Testes executados com sucesso!"
else
    echo "Falha na execução dos testes."
    exit 1
fi
