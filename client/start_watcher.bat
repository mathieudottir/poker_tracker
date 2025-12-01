@echo off
REM Watcher Client pour Windows

echo ================================
echo Winamax Expresso Tracker Watcher
echo ================================
echo.

REM Vérifier Python
python --version >nul 2>&1
if errorlevel 1 (
    echo ERREUR: Python n'est pas installe ou pas dans le PATH
    echo Telecharger Python sur: https://www.python.org/downloads/
    pause
    exit /b 1
)

REM Installer les dépendances si nécessaire
echo Installation des dependances...
pip install -q -r requirements.txt

REM Lancer le watcher
echo.
python watcher_client.py

pause
