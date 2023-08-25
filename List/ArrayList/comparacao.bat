@echo OFF
IF NOT EXIST "testeCache/testeDireto.exe" (
    echo Buildando o testeDireto.exe:
    cd testeCache
    @echo ON
    go build -tags direto -o testeDireto.exe
    @echo OFF
    echo.
    cd ../
)

IF NOT EXIST "testeCache/testeReverso.exe" (
    echo Buildando o testeDireto.exe:
    cd testeCache
    @echo ON
    go build -o testeReverso.exe
    @echo OFF
    echo.
    cd ../
)

echo Iniciando testes para o AddPos direto:
@echo ON
testeCache\testeDireto.exe
@echo OFF
echo. & echo Iniciando testes para o AddPos reverso:
@echo ON
testeCache\testeReverso.exe
@echo OFF