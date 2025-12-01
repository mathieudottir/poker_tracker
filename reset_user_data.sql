-- Script pour vider les données utilisateur et vérifier les settings
-- Connexion: psql -U postgres -d poker_tracker -f reset_user_data.sql

-- 1. Afficher les settings utilisateur actuels
\echo '=== Settings utilisateur actuel ==='
SELECT id, username, player_name, wina_status, hh_directory, dev_mode, created_at
FROM users
WHERE username = 'mathieu';

-- 2. Afficher les statuts rakeback disponibles
\echo ''
\echo '=== Statuts rakeback Winamax disponibles ==='
SELECT status_name, rakeback_pct, miles_required, annual_bonus
FROM rakeback_status
ORDER BY rakeback_pct;

-- 3. Vider les données de tournois et mains pour l'utilisateur mathieu
\echo ''
\echo '=== Suppression des données de tournois et mains ==='
DELETE FROM hands WHERE tournament_id IN (
    SELECT id FROM tournaments WHERE user_id = (
        SELECT id FROM users WHERE username = 'mathieu'
    )
);
DELETE FROM tournaments WHERE user_id = (
    SELECT id FROM users WHERE username = 'mathieu'
);

-- 4. Vérifier que les données ont été supprimées
\echo ''
\echo '=== Vérification ==='
SELECT
    (SELECT COUNT(*) FROM tournaments WHERE user_id = (SELECT id FROM users WHERE username = 'mathieu')) as tournaments_count,
    (SELECT COUNT(*) FROM hands WHERE tournament_id IN (SELECT id FROM tournaments WHERE user_id = (SELECT id FROM users WHERE username = 'mathieu'))) as hands_count;

\echo ''
\echo '=== Données supprimées avec succès ! ==='
\echo ''
\echo 'Pour mettre à jour le wina_status, exécutez:'
\echo 'UPDATE users SET wina_status = ''Diamond 2'' WHERE username = ''mathieu'';'
\echo ''
\echo 'Statuts disponibles: Aluminium, Bronze, Argent, Or, Platine, Diamond 1, Diamond 2, Diamond 3, Diamond 4'
